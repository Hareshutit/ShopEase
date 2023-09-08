// Package v2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package v2

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Вернуть объявления из корзины.
	// (GET /api/v1/cart)
	GetCart(ctx echo.Context) error
	// Удалить из корзины товар.
	// (DELETE /api/v1/cart/{id})
	RemoveCart(ctx echo.Context, id string) error
	// Добавить в корзину.
	// (POST /api/v1/cart/{id})
	AddCart(ctx echo.Context, id string) error
	// Вернуть избранное.
	// (GET /api/v1/favorite)
	GetFavorite(ctx echo.Context) error
	// Удалить из избранных товар.
	// (DELETE /api/v1/favorite/{id})
	RemoveFavorite(ctx echo.Context, id string) error
	// Добавить в избранное.
	// (POST /api/v1/favorite/{id})
	AddFavorite(ctx echo.Context, id string) error
	// Возврат массива постов.
	// (GET /api/v1/post)
	GetMiniPost(ctx echo.Context, params GetMiniPostParams) error
	// Создать новое объявление.
	// (POST /api/v1/post)
	CreatePost(ctx echo.Context) error
	// Удалить объявление.
	// (DELETE /api/v1/post/{id})
	DeletePost(ctx echo.Context, id string) error
	// Вернуть объявление по id.
	// (GET /api/v1/post/{id})
	GetIdPost(ctx echo.Context, id string) error
	// Обновить объявление.
	// (PATCH /api/v1/post/{id})
	UpdatePost(ctx echo.Context, id string) error
	// Поиск.
	// (GET /api/v1/search)
	Search(ctx echo.Context, params SearchParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCart converts echo context to params.
func (w *ServerInterfaceWrapper) GetCart(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCart(ctx)
	return err
}

// RemoveCart converts echo context to params.
func (w *ServerInterfaceWrapper) RemoveCart(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RemoveCart(ctx, id)
	return err
}

// AddCart converts echo context to params.
func (w *ServerInterfaceWrapper) AddCart(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddCart(ctx, id)
	return err
}

// GetFavorite converts echo context to params.
func (w *ServerInterfaceWrapper) GetFavorite(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFavorite(ctx)
	return err
}

// RemoveFavorite converts echo context to params.
func (w *ServerInterfaceWrapper) RemoveFavorite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RemoveFavorite(ctx, id)
	return err
}

// AddFavorite converts echo context to params.
func (w *ServerInterfaceWrapper) AddFavorite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddFavorite(ctx, id)
	return err
}

// GetMiniPost converts echo context to params.
func (w *ServerInterfaceWrapper) GetMiniPost(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMiniPostParams
	// ------------- Required query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, true, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Required query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, true, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", ctx.QueryParams(), &params.Sort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort: %s", err))
	}

	// ------------- Optional query parameter "user" -------------

	err = runtime.BindQueryParameter("form", true, false, "user", ctx.QueryParams(), &params.User)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user: %s", err))
	}

	// ------------- Optional query parameter "tag" -------------

	err = runtime.BindQueryParameter("form", true, false, "tag", ctx.QueryParams(), &params.Tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tag: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMiniPost(ctx, params)
	return err
}

// CreatePost converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePost(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePost(ctx)
	return err
}

// DeletePost converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePost(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePost(ctx, id)
	return err
}

// GetIdPost converts echo context to params.
func (w *ServerInterfaceWrapper) GetIdPost(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetIdPost(ctx, id)
	return err
}

// UpdatePost converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePost(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdatePost(ctx, id)
	return err
}

// Search converts echo context to params.
func (w *ServerInterfaceWrapper) Search(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchParams
	// ------------- Required query parameter "query" -------------

	err = runtime.BindQueryParameter("form", true, true, "query", ctx.QueryParams(), &params.Query)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter query: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Search(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/v1/cart", wrapper.GetCart)
	router.DELETE(baseURL+"/api/v1/cart/:id", wrapper.RemoveCart)
	router.POST(baseURL+"/api/v1/cart/:id", wrapper.AddCart)
	router.GET(baseURL+"/api/v1/favorite", wrapper.GetFavorite)
	router.DELETE(baseURL+"/api/v1/favorite/:id", wrapper.RemoveFavorite)
	router.POST(baseURL+"/api/v1/favorite/:id", wrapper.AddFavorite)
	router.GET(baseURL+"/api/v1/post", wrapper.GetMiniPost)
	router.POST(baseURL+"/api/v1/post", wrapper.CreatePost)
	router.DELETE(baseURL+"/api/v1/post/:id", wrapper.DeletePost)
	router.GET(baseURL+"/api/v1/post/:id", wrapper.GetIdPost)
	router.PATCH(baseURL+"/api/v1/post/:id", wrapper.UpdatePost)
	router.GET(baseURL+"/api/v1/search", wrapper.Search)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
