package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwt"
)

type JWSValidator interface {
	ValidateJWS(jws string) (jwt.Token, error)
}

var (
	ErrNoAuthHeader      = errors.New("Authorization header is missing")
	ErrInvalidCookie     = errors.New("Authorization cookie is invalid")
	ErrInvalidAuthHeader = errors.New("Authorization header is malformed")
	ErrClaimsInvalid     = errors.New("Provided claims do not match expected scopes")
)

func NewAuthenticatorAccess(v JWSValidator) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return AuthenticateAccess(v, ctx, input)
	}
}

func NewAuthenticatorRefresh(v JWSValidator) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return AuthenticateRefresh(v, ctx, input)
	}
}

func AuthenticateAccess(v JWSValidator, ctx context.Context,
	input *openapi3filter.AuthenticationInput) error {

	//fmt.Println(input.SecuritySchemeName)cookieAuth
	if input.SecuritySchemeName != "bearerAuth" {
		return fmt.Errorf("security scheme %s != 'bearerAuth'", input.SecuritySchemeName)
	}

	jws, err := GetJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}

	_, err = v.ValidateJWS(jws)
	if err != nil {
		return fmt.Errorf("validating JWS: %w", err)
	}

	return nil
}

func AuthenticateRefresh(v JWSValidator, ctx context.Context,
	input *openapi3filter.AuthenticationInput) error {

	//fmt.Println(input.SecuritySchemeName)cookieAuth
	if input.SecuritySchemeName != "cookieAuth" {
		return fmt.Errorf("security scheme %s != 'cookieAuth'", input.SecuritySchemeName)
	}

	jws, err := GetJWSFromCookie(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}

	_, err = v.ValidateJWS(jws)
	if err != nil {
		return fmt.Errorf("validating JWS: %w", err)
	}

	return nil
}

func GetJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")

	if authHdr == "" {
		return "", ErrNoAuthHeader
	}

	prefix := "bearerAuth "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}

func GetJWSFromCookie(req *http.Request) (string, error) {
	authCook, err := req.Cookie("Refresh")

	if err != nil || authCook.Value == "" {
		return "", ErrInvalidCookie
	}

	return authCook.Value, nil
}

func CreateMiddlewareAccess(v JWSValidator, swagger *openapi3.T) ([]echo.MiddlewareFunc, error) {

	validator := oapimiddleware.OapiRequestValidatorWithOptions(swagger,
		&oapimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticatorAccess(v),
			},
		})

	return []echo.MiddlewareFunc{validator}, nil
}

func CreateMiddlewareRefresh(v JWSValidator, swagger *openapi3.T) ([]echo.MiddlewareFunc, error) {

	validator := oapimiddleware.OapiRequestValidatorWithOptions(swagger,
		&oapimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticatorRefresh(v),
			},
		})

	return []echo.MiddlewareFunc{validator}, nil
}
