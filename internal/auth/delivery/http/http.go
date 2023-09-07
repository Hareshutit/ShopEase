package v2

import (
	"context"
	"fmt"
	"log"
	"net/http"

	app "github.com/Hareshutit/ShopEase/internal/auth/usecase"

	servGrpc "github.com/Hareshutit/ShopEase/pkg/grpc/user"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/auth/models.cfg.yml ../../../../api/openapi/auth/auth.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/auth/server.cfg.yml ../../../../api/openapi/auth/auth.yml

func sendUserError(ctx echo.Context, code int, message string) error {
	userErr := ErrorHTTP{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, userErr)
	return err
}

type HttpServer struct {
	command app.Commands
	query   app.Queries
}

func (d *HttpServer) Login(ctx echo.Context) error {

	var data SignUp

	err := ctx.Bind(&data)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}

	grcpConn, err := grpc.Dial(
		"user_service:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	sessManager := servGrpc.NewUserClient(grcpConn)

	cr := servGrpc.UserCheck{Login: data.Login, Password: data.Password}

	wd, err := sessManager.CheckAccount(context.Background(), &cr)

	if wd.GetValue() == "" {
		return sendUserError(ctx, http.StatusBadRequest, "Ошибка логина или паролся")
	}

	claims := make(map[string]string)
	claims["id"] = wd.GetValue()

	result, err := d.command.CreateToken.CreateJWSWithClaims(claims, "appUniqFront", "auth")
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusOK, string(result))
}
