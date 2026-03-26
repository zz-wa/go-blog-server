package middleware

import (
	"blog_r/internal/pkg/jwt"
	"blog_r/internal/response"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return c.JSON(http.StatusUnauthorized, response.Fail(response.CodeUnauthorized, "未登录"))
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, response.Fail(response.CodeUnauthorized, "token format error"))
		}
		tokenString := parts[1]

		claims, err := jwt.VerifyToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.Fail(response.CodeUnauthorized, "invalid token"))
		}
		c.Set("user_id", claims.UserID)
		return next(c)
	}
}
