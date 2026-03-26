package api

import (
	"blog_r/internal/response"
	"net/http"

	"github.com/labstack/echo/v5"
)

func Health(c *echo.Context) error {
	return c.JSON(http.StatusOK, response.OK(map[string]string{
		"status": "ok",
	}))
}
