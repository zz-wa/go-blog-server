package api

import (
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/config"
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetConfigList(c *echo.Context) error {
	ConfigService := config.NewConfigService()
	configList, err := ConfigService.GetConfigList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get configList"))
	}

	return c.JSON(http.StatusOK, response.OK(configList))
}

func GetConfigAbout(c *echo.Context) error {
	ConfigService := config.NewConfigService()
	about, err := ConfigService.GetConfigAbout("about")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get about"))
	}
	return c.JSON(http.StatusOK, response.OK(about))
}

func UpdateAbout(c *echo.Context) error {
	req := request.UpdateConfigReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	req.Key = "about"
	ConfigService := config.NewConfigService()
	if err := ConfigService.UpdateConfig(req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to update about"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}
func UpdateConfig(c *echo.Context) error {
	req := request.UpdateConfigReq{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get configList"))
	}
	req.Key = c.Param("key")
	ConfigService := config.NewConfigService()
	if err := ConfigService.UpdateConfig(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "fail to update configList"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}
