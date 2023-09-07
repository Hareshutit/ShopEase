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

	serverGrpc "github.com/Hareshutit/ShopEase/internal/auth/delivery/grpc"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"

	"github.com/Hareshutit/ShopEase/internal/auth/usecase"

	v2 "github.com/Hareshutit/ShopEase/internal/auth/delivery/http"
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

func Run() {
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
	command, query := usecase.NewUsecase(ctx)

	serverHandler := v2.CreateHttpServer(command, query)

	grpcHandler := serverGrpc.CreateGrpcServer(command, query)

	server := grpc.NewServer()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(oapimiddleware.OapiRequestValidator(swagger))

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
