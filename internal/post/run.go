package post

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hareshutit/ShopEase/internal/post/usecase"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/rs/zerolog"

	config "github.com/Hareshutit/ShopEase/config/post"

	authmiddlevare "github.com/Hareshutit/ShopEase/pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	v2 "github.com/Hareshutit/ShopEase/internal/post/delivery/http"
)

func AsyncRunHTTP(serverH *echo.Echo, cfg config.Config) error {
	go func() {
		err := serverH.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port))
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

func Run(cfg config.Config) {

	ctx := context.Background()

	swagger, err := v2.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки спецификации swagger\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	command, query := usecase.NewUsecase(ctx, cfg)

	serverHandler := v2.CreateHttpServer(command, query)

	serverH := echo.New()

	instAuth, err := authmiddlevare.NewInstanceAuthenticator(cfg.Authorization.Verify, jwa.ES256, "shopease.com", "auth.shopease.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	mw, err := authmiddlevare.CreateMiddlewareAccess(instAuth, swagger)
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
	serverH.Logger.Fatal(serverH.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port)))
}
