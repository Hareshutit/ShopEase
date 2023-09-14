package v2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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

	accessToken, code, err := d.command.CreateAccessToken.Create(wd.GetValue())
	if err != nil {
		return sendUserError(ctx, code, fmt.Sprintf("%v", err))
	}

	ctxn := context.TODO()

	refreshToken, code, err := d.command.CreateRefreshToken.Create(ctxn, wd.GetValue())
	if err != nil {
		return sendUserError(ctx, code, fmt.Sprintf("%v", err))
	}

	cookie := new(http.Cookie)
	cookie.Name = "Refresh"
	cookie.Value = string(refreshToken)
	cookie.Expires = time.Now().Add(30 * 24 * time.Hour)
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Secure = true
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, string(accessToken))
}

func (d *HttpServer) Refresh(ctx echo.Context) error {
	cookie, err := ctx.Cookie("Refresh")
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	refreshToken := cookie.Value
	ctxn := context.TODO()

	newRefreshToken, code, err := d.command.UpdateRefreshToken.Update(ctxn, refreshToken)
	if err != nil {
		return sendUserError(ctx, code, fmt.Sprintf("%v", err))
	}
	cookie.Value = string(newRefreshToken)

	accessToken, code, err := d.command.CreateAccessToken.Create(wd.GetValue())
	if err != nil {
		return sendUserError(ctx, code, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusOK, accessToken)
}

func (d *HttpServer) Logout(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, nil)
}
