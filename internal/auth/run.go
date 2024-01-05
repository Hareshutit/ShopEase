package auth

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/Hareshutit/ShopEase/config/auth"
	serverGrpc "github.com/Hareshutit/ShopEase/internal/auth/delivery/grpc"
	"github.com/Hareshutit/ShopEase/internal/auth/repository"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/rs/zerolog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"

	"github.com/Hareshutit/ShopEase/internal/auth/usecase"

	v2 "github.com/Hareshutit/ShopEase/internal/auth/delivery/http"
	authmiddlevare "github.com/Hareshutit/ShopEase/pkg/middleware"
)

func AsyncRunHTTP(serverH *echo.Echo) error {
	go func() {
		err := serverH.Start(fmt.Sprintf("0.0.0.0:%d", 8082))
		if err != nil && err != http.ErrServerClosed {
			serverH.Logger.Fatal("shutting down the server")
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	timeout := time.Duration(10)
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return serverH.Shutdown(ctx)
}

func AsyncRunGrpc(serverG *grpc.Server, lis net.Listener) error {
	go func() {
		err := serverG.Serve(lis)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
			os.Exit(1)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	serverG.GracefulStop()

	return nil
}

func Run(cfg config.Config) {
	ctx := context.Background()

	swagger, err := v2.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки спецификации swagger\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	redisRepo := repository.CreateRedisRepository(cfg)
	command, query := usecase.NewUsecase(ctx, &redisRepo, cfg)

	serverHandler := v2.CreateHttpServer(command, query)

	grpcHandler := serverGrpc.CreateGrpcServer(command, query)

	serverG := grpc.NewServer()

	serverH := echo.New()

	instAuth, err := authmiddlevare.NewInstanceAuthenticator(&cfg.KeyValue.Refresh.PublicKey, jwa.ES256, "shopease.com", "auth.shopease.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	mw, err := authmiddlevare.CreateMiddlewareRefresh(instAuth, swagger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	logger := zerolog.New(os.Stdout)
	serverH.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	serverH.Use(mw...)

	v2.RegisterHandlers(serverH, &serverHandler)

	serverGrpc.RegisterAuthServer(serverG, &grpcHandler)

	errs := make(chan error, 2)
	defer close(errs)

	go func() {
		errs <- AsyncRunHTTP(serverH)
	}()

	go func() {
		errs <- AsyncRunGrpc(serverG, lis)
	}()
	err = <-errs

	log.Warn("Terminating aplication:", err)
}
