package user

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/Hareshutit/ShopEase/config/user"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/rs/zerolog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"google.golang.org/grpc"

	serverGrpc "github.com/Hareshutit/ShopEase/internal/user/delivery/grpc"
	v2 "github.com/Hareshutit/ShopEase/internal/user/delivery/http"
	"github.com/Hareshutit/ShopEase/internal/user/repository"
	"github.com/Hareshutit/ShopEase/internal/user/usecase"
	authmiddlevare "github.com/Hareshutit/ShopEase/pkg/middleware"
)

func AsyncRunHTTP(server *echo.Echo, log zerolog.Logger, cfg config.Config) error {
	go func() {
		err := server.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port))
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Shutting down the server")
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

	return server.Shutdown(ctx)
}

func AsyncRunGrpc(server *grpc.Server, log zerolog.Logger, cfg config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grcp.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("Error load grpc-server")
	}

	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatal().Err(err).Msg("Error run grpc-server")
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	server.GracefulStop()

	return nil
}

func Run(log zerolog.Logger, cfg config.Config) {

	ctx := context.Background()

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Db.User, cfg.Db.DataBaseName, cfg.Db.Password, cfg.Db.Host,
		cfg.Db.Port, cfg.Db.Sslmode)

	Repository := repository.CreatePostgressRepository(dsn, log)

	command, query := usecase.NewUsecase(ctx, Repository, &Repository, log)

	serverHandler := v2.CreateHttpServer(command, query, log)

	grpcHandler := serverGrpc.CreateGrpcServer(command, query, log)

	serverG := grpc.NewServer()

	serverH := echo.New()

	serverH.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	log = log.With().Str("Layer", "Starting").Logger()

	swagger, err := v2.GetSwagger()
	if err != nil {
		log.Fatal().Err(err).Msg("Error load swagger specification")
	}
	swagger.Servers = nil

	instAuth, err := authmiddlevare.NewInstanceAuthenticator(cfg.Authorization.Verify, jwa.ES256, "shopease.com", "auth.shopease.com")
	if err != nil {
		log.Fatal().Err(err).Msg("Error load authenticator")
	}

	mw, err := authmiddlevare.CreateMiddlewareAccess(instAuth, swagger)
	if err != nil {
		log.Fatal().Err(err).Msg("Error load middleware")
	}

	serverH.Use(mw...)

	v2.RegisterHandlers(serverH, &serverHandler)

	serverGrpc.RegisterUserServer(serverG, &grpcHandler)

	errs := make(chan error, 2)
	defer close(errs)

	go func() {
		errs <- AsyncRunHTTP(serverH, log, cfg)
	}()

	go func() {
		errs <- AsyncRunGrpc(serverG, log, cfg)
	}()
	err = <-errs

	log.Warn().Err(err).Msg("Terminating aplication")
}
