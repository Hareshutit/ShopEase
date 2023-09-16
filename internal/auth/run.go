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

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"

	"github.com/Hareshutit/ShopEase/internal/auth/usecase"

	v2 "github.com/Hareshutit/ShopEase/internal/auth/delivery/http"
	authmiddlevare "github.com/Hareshutit/ShopEase/pkg/middleware"
)

func AsyncRunHTTP(e *echo.Echo) error {
	go func() {
		err := e.Start(fmt.Sprintf("0.0.0.0:%d", 8082))
		if err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
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

	return e.Shutdown(ctx)
}

func AsyncRunGrpc(server *grpc.Server, lis net.Listener) error {
	go func() {
		err := server.Serve(lis)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
			os.Exit(1)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	server.GracefulStop()

	return nil
}

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIN2dALnjdcZaIZg4QuA6Dw+kxiSW502kJfmBN3priIhPoAoGCCqGSM49
AwEHoUQDQgAE4pPyvrB9ghqkT1Llk0A42lixkugFd/TBdOp6wf69O9Nndnp4+HcR
s9SlG/8hjB2Hz42v4p3haKWv3uS1C6ahCQ==
-----END EC PRIVATE KEY-----`

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

	server := grpc.NewServer()

	e := echo.New()

	instAuth, err := authmiddlevare.NewInstanceAuthenticator(cfg.KeyValue.Refresh, jwa.ES256, "shopease.com", "auth.shopease.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	mw, err := authmiddlevare.CreateMiddlewareRefresh(instAuth, swagger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	e.Use(middleware.Logger())
	e.Use(mw...)

	v2.RegisterHandlers(e, &serverHandler)

	serverGrpc.RegisterAuthServer(server, &grpcHandler)

	errs := make(chan error, 2)
	go func() {
		errs <- AsyncRunHTTP(e)
	}()

	go func() {
		errs <- AsyncRunGrpc(server, lis)
	}()
	err = <-errs

	log.Warn("Terminating aplication:", err)
}
