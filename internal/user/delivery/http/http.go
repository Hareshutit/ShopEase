package v2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Hareshutit/ShopEase/internal/user/domain"
	app "github.com/Hareshutit/ShopEase/internal/user/usecase"

	"github.com/Hareshutit/ShopEase/pkg/jwt"

	servGrpc "github.com/Hareshutit/ShopEase/pkg/grpc/auth"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/user/models.cfg.yml ../../../../api/openapi/user/user.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/user/server.cfg.yml ../../../../api/openapi/user/user.yml

func sendUserError(ctx echo.Context, code int, message error) error {
	userErr := ErrorHTTP{
		Code:    int32(code),
		Message: message.Error(),
	}
	err := ctx.JSON(code, userErr)
	return err
}

type HttpServer struct {
	command app.Commands
	query   app.Queries
}

// Handler отвечает за создание пользователя
func (a *HttpServer) CreateUser(ctx echo.Context) error {
	var newUser CreateUser

	err := ctx.Bind(&newUser)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, errors.New("Bad request"))
	}

	user := domain.User{
		Email:       newUser.Email,
		PhoneNumber: newUser.PhoneNumber,

		Login:    newUser.Login,
		Password: newUser.Password,

		Name: newUser.Name,

		PathToAvatar: newUser.Avatar,
	}

	uuids, err := a.command.CreateUser.Handle(context.Background(), newUser.PasswordCheck, user)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}

	grcpConn, err := grpc.Dial(
		"auth_service:8085",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	sessManager := servGrpc.NewAuthClient(grcpConn)

	cr := servGrpc.Id{Id: uuids.String()}

	fmt.Println(cr.GetId())

	wd, err := sessManager.GenerateToken(context.Background(), &cr)

	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, errors.New("Ошибка генерации токена"))
	}

	cookie := new(http.Cookie)
	cookie.Name = "Refresh"
	cookie.Value = string(wd.Refresh)
	cookie.Expires = time.Now().Add(30 * 24 * time.Hour)
	cookie.MaxAge = 30 * 24 * 60 * 60 // Время жизни в секундах
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Secure = true
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusCreated, wd.GetAccess())
}

// Handler отвечает за обновление данных пользователя
func (a *HttpServer) UpdateUser(ctx echo.Context) error {
	var updateUser CreateUser

	err := ctx.Bind(&updateUser)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, errors.New("Bad request"))
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	id := jwt.ClaimParse(headerAuth, "sub")

	uuids, err := uuid.Parse(id)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}

	user := domain.User{
		Id: uuids,

		Email:       updateUser.Email,
		PhoneNumber: updateUser.PhoneNumber,

		Login:    updateUser.Login,
		Password: updateUser.Password,

		Name: updateUser.Name,

		PathToAvatar: updateUser.Avatar,
	}

	err = a.command.UpdateUser.Handle(context.Background(), user)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}

	return nil
}

// Handler отвечает за удаление пользователя
func (a *HttpServer) DeleteUser(ctx echo.Context) error {
	headerAuth := ctx.Request().Header.Get("Authorization")
	id := jwt.ClaimParse(headerAuth, "sub")

	uuids, err := uuid.Parse(id)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}
	err = a.command.DeleteUser.Handle(context.Background(), uuids) // Значения из авторизации
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}
	return nil
}

// Handler отвечает за возврат данных авторизованного пользователя
func (a HttpServer) GetUser(ctx echo.Context) error {
	var user domain.User
	headerAuth := ctx.Request().Header.Get("Authorization")
	id := jwt.ClaimParse(headerAuth, "sub")

	uuids, err := uuid.Parse(id)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}
	fmt.Println(uuids)

	user, err = a.query.GetUser.Handle(context.Background(), uuids)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}
	result := CreateUser{
		Name:          user.Name,
		Email:         user.Email,
		Login:         user.Login,
		Password:      user.Password,
		PasswordCheck: user.CheckPassword,
		PhoneNumber:   user.PhoneNumber,
		Avatar:        user.PathToAvatar,
	}
	return ctx.JSON(http.StatusOK, result)
}

// Handler отвечает за возврат данных неавторизованных пользователей
func (a *HttpServer) FindUserByID(ctx echo.Context, id string) error {
	var user domain.User
	uuids, err := uuid.Parse(id)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}
	user, err = a.query.FindByIdUser.Handle(context.Background(), uuids)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, err)
	}
	result := GetUser{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Avatar:      user.PathToAvatar,
	}
	return ctx.JSON(http.StatusOK, result)
}
