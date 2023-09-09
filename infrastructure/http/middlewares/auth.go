package middleware

import (
	"net/http"
	"strings"

	"github.com/FianGumilar/e-wallet/interfaces"
	"github.com/labstack/echo/v4"
)

func Authenticate(userService interfaces.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer", "")
			if token == "" {
				return c.String(http.StatusUnauthorized, "failed to authorize")
			}
			user, err := userService.Validate(c.Request().Context(), token)
			if err != nil {
				return c.String(http.StatusUnauthorized, "failed to authorize")
			}
			c.Set("x-user", user)
			return next(c)
		}
	}
}
