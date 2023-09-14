package v2

import (
	"context"
	"errors"
	"net/http"

	"github.com/Hareshutit/ShopEase/internal/post/domain"

	app "github.com/Hareshutit/ShopEase/internal/post/usecase"

	"github.com/Hareshutit/ShopEase/pkg/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/post/models.cfg.yml ../../../../api/openapi/post/post.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/post/server.cfg.yml ../../../../api/openapi/post/post.yml

func sendPostError(ctx echo.Context, code int, message error) error {
	postErr := ErrorHTTP{
		Code:    int32(code),
		Message: message.Error(),
	}
	err := ctx.JSON(code, postErr)
	return err
}

type HttpServer struct {
	command app.Commands
	query   app.Queries
}

func (a *HttpServer) CreatePost(ctx echo.Context) error {
	var post CreatePost

	err := ctx.Bind(&post)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, errors.New("Incorrect request format"))
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	userId := jwt.ClaimParse(headerAuth, "id")

	uuidUser, err := uuid.Parse(userId)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, err)
	}

	postDTO := domain.Post{
		UserID:      &uuidUser,
		Title:       &post.Title,
		Description: &post.Description,
		Price:       &post.Price,
		Category:    &post.Category,
		PathImages:  &post.PathImages,
	}

	uuidPost, code, err := a.command.Create.Handle(context.Background(), postDTO)
	if err != nil {
		return sendPostError(ctx, code, err)
	}

	return ctx.JSON(http.StatusCreated, uuidPost)
}

func (a *HttpServer) UpdatePost(ctx echo.Context, id string) error {
	var post EditPost

	err := ctx.Bind(&post)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, errors.New("Incorrect request format"))
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	idUser := jwt.ClaimParse(headerAuth, "id")

	uuidUser, err := uuid.Parse(idUser)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, err)
	}

	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, err)
	}

	postDTO := domain.Post{
		Id:          &uuidPost,
		UserID:      &uuidUser,
		Title:       post.Title,
		Description: post.Description,
		Status:      post.Status,
		Price:       post.Price,
		Category:    post.Category,
		PathImages:  post.PathImages,
	}

	code, err := a.command.Update.Handle(context.Background(), postDTO)
	if err != nil {
		return sendPostError(ctx, code, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (a *HttpServer) DeletePost(ctx echo.Context, id string) error {
	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, err)
	}

	code, err := a.command.Delete.Handle(context.Background(), uuidPost)
	if err != nil {
		return sendPostError(ctx, code, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (a HttpServer) GetIdPost(ctx echo.Context, id string) error {
	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, err)
	}

	resultDTO, code, err := a.query.GetById.Handle(context.Background(), uuidPost)
	if err != nil {
		return sendPostError(ctx, code, err)
	}

	result := FullPost{
		Category:    *resultDTO.Category,
		Status:      *resultDTO.Status,
		Description: *resultDTO.Description,
		PathImages:  *resultDTO.PathImages,
		Price:       *resultDTO.Price,
		Title:       *resultDTO.Title,
		UserId:      resultDTO.UserID.String(),
		Views:       *resultDTO.Views,
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (a HttpServer) GetMiniPost(ctx echo.Context, params GetMiniPostParams) error {

	var uuidUserPointer *uuid.UUID

	if params.User != nil {
		uuidUser, err := uuid.Parse(*params.User)
		if err != nil {
			return sendPostError(ctx, http.StatusBadRequest, err)
		}
		uuidUserPointer = &uuidUser
	}

	param := domain.Parameters{
		Offset:   &params.Offset,
		Limit:    &params.Limit,
		Status:   params.Status,
		Sort:     params.Sort,
		UserId:   uuidUserPointer,
		Category: params.Tag,
	}

	resultDTO, code, err := a.query.GetMiniObject.Handle(context.Background(), param)
	if err != nil {
		return sendPostError(ctx, code, err)
	}

	var result []MiniPost
	for _, post := range resultDTO {
		p := MiniPost{
			PostId:     post.Id.String(),
			PathImages: *post.PathImages,
			Price:      *post.Price,
			Title:      *post.Title,
			UserId:     post.UserID.String(),
			Views:      *post.Views,
		}
		result = append(result, p)
	}

	if result == nil {
		return ctx.NoContent(http.StatusNoContent)
	}

	return ctx.JSON(http.StatusOK, result)
}
