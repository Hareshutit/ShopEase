package v2

import (
	"context"
	"errors"
	"log"
	"net/http"

	app "github.com/Hareshutit/ShopEase/internal/auth/usecase"

	servGrpc "github.com/Hareshutit/ShopEase/pkg/grpc/user"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/auth/models.cfg.yml ../../../../api/openapi/auth/auth.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/auth/server.cfg.yml ../../../../api/openapi/auth/auth.yml

type HttpServer struct {
	command app.Commands
	query   app.Queries
}

func (d *HttpServer) Login(ctx echo.Context) error {

	var data SignUp

	err := ctx.Bind(&data)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, errors.New("Bad request"))
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
		return sendUserError(ctx, http.StatusBadRequest, errors.New("Wrong login or password"))
	}

	accessToken, code, err := d.command.CreateAccessToken.Create(wd.GetValue())
	if err != nil {
		return sendUserError(ctx, code, err)
	}

	ctxn := context.TODO()

	refreshToken, code, err := d.command.CreateRefreshToken.Create(ctxn, wd.GetValue())
	if err != nil {
		return sendUserError(ctx, code, err)
	}

	cookie := new(http.Cookie)
	cookieRefresh(cookie, refreshToken)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, string(accessToken))
}

func (d *HttpServer) Refresh(ctx echo.Context) error {
	cookie, err := ctx.Cookie("Refresh")
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}
	refreshToken := cookie.Value
	ctxn := context.TODO()

	newRefreshToken, idUser, code, err := d.command.UpdateRefreshToken.Update(ctxn, refreshToken)
	if err != nil {
		return sendUserError(ctx, code, err)
	}

	accessToken, code, err := d.command.CreateAccessToken.Create(*idUser)
	if err != nil {
		return sendUserError(ctx, code, err)
	}

	cookieRefresh(cookie, newRefreshToken)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, accessToken)
}

func (d *HttpServer) Logout(ctx echo.Context) error {
	cookie, err := ctx.Cookie("Refresh")
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}
	refreshToken := cookie.Value
	ctxn := context.TODO()

	code, err := d.command.DeleteRefreshToken.Delete(ctxn, refreshToken)
	if err != nil {
		return sendUserError(ctx, code, err)
	}

	cookieClear(cookie)
	ctx.SetCookie(cookie)

	return ctx.NoContent(http.StatusOK)
}
