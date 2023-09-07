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
	ValidateJWS(jws string, audience string, issuer string) (jwt.Token, error)
}

var (
	ErrNoAuthHeader      = errors.New("Authorization header is missing")
	ErrInvalidAuthHeader = errors.New("Authorization header is malformed")
	ErrClaimsInvalid     = errors.New("Provided claims do not match expected scopes")
)

func NewAuthenticator(v JWSValidator, audience string, issuer string) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(v, ctx, input, audience, issuer)
	}
}

func Authenticate(v JWSValidator, ctx context.Context,
	input *openapi3filter.AuthenticationInput, audience string, issuer string) error {

	//fmt.Println(input.SecuritySchemeName)
	if input.SecuritySchemeName != "bearerAuth" {
		return fmt.Errorf("security scheme %s != 'bearerAuth'", input.SecuritySchemeName)
	}

	jws, err := GetJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}

	_, err = v.ValidateJWS(jws, audience, issuer)
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

func CreateMiddleware(v JWSValidator, swagger *openapi3.T, audience string, issuer string) ([]echo.MiddlewareFunc, error) {

	validator := oapimiddleware.OapiRequestValidatorWithOptions(swagger,
		&oapimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticator(v, audience, issuer),
			},
		})

	return []echo.MiddlewareFunc{validator}, nil
}
