package v2

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func sendUserError(ctx echo.Context, code int, message error) error {
	userErr := ErrorHTTP{
		Code:    int32(code),
		Message: message.Error(),
	}
	err := ctx.JSON(code, userErr)
	return err
}

func cookieRefresh(cookie *http.Cookie, token []byte) {
	cookie.Name = "Refresh"
	cookie.Value = string(token)
	cookie.Expires = time.Now().Add(30 * 24 * time.Hour)
	cookie.MaxAge = 30 * 24 * 60 * 60 // Время жизни в секундах
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Secure = true
	cookie.HttpOnly = true
}

func cookieClear(cookie *http.Cookie) {
	cookie.Name = "Refresh"
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	cookie.MaxAge = -1
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Secure = true
	cookie.HttpOnly = true
}
