package api

import (
	"blog_r/internal/response"
	"blog_r/internal/service/upload"
	"net/http"

	"github.com/labstack/echo/v5"
)

func Upload(c *echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "bad request"))
	}
	local := &upload.Local{}
	_, filePath, err := local.Uploads(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, response.OK(filePath))

}
