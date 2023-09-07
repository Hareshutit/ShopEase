package jwt

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func ClaimParse(jwts string, claim string) string {

	prefix := "bearerAuth "
	token, _, err := new(jwt.Parser).ParseUnverified(strings.TrimPrefix(jwts, prefix), jwt.MapClaims{})
	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return fmt.Sprint(claims[claim])
	}

	return ""
}
