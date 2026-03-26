package api

import (
	"blog_r/internal/request"
	"blog_r/internal/response"
	operation_log "blog_r/internal/service/operation_log"
	"net/http"

	"github.com/labstack/echo/v5"
)

func OperationLogList(c *echo.Context) error {
	req := new(request.LogListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, err.Error()))
	}

	svc := operation_log.NewOperationLogService()
	list, total, err := svc.OperationLogList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error for get details"))
	}

	return c.JSON(http.StatusOK, response.OK(response.OperationLogListRes{
		List:  list,
		Total: total,
	}))
}
