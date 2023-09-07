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

	config "github.com/Hareshutit/ShopEase/config/post"

	authmiddlevare "github.com/Hareshutit/ShopEase/pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	v2 "github.com/Hareshutit/ShopEase/internal/post/delivery/http"
)

func AsyncRunHTTP(e *echo.Echo, cfg config.Config) error {
	go func() {
		err := e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port))
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

	e := echo.New()

	fa, err := authmiddlevare.NewInstanceAuthenticator()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	mw, err := authmiddlevare.CreateMiddleware(fa, swagger, "shopeaseFrontend", "shopeaseUser")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}
	e.Use(middleware.Logger())
	e.Use(mw...)

	v2.RegisterHandlers(e, &serverHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port)))
}
