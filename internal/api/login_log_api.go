package api

import (
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/login_log"
	"net/http"

	"github.com/labstack/echo/v5"
)

/*
写函数 LoginLogList(c *echo.Context) error
流程：Bind query -> Validate -> 调 NewLoginLogService().LoginLogList(req) -> 返回 response.OK(list+total)
错误处理沿用你现有风格：参数错 400，服务错 500
*/

func LoginLogList(c *echo.Context) error {
	req := new(request.LogListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, err.Error()))
	}
	loginLogService := login_log.NewLoginLogService()
	list, total, err := loginLogService.LoginLogList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error for get details"))
	}
	return c.JSON(http.StatusOK, response.OK(response.LoginLogListRes{
		List:  list,
		Total: total,
	}))
}
