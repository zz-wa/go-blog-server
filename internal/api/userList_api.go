package api

import (
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/user"
	"net/http"

	"github.com/labstack/echo/v5"
)

func UserList(c *echo.Context) error {
	req := new(request.UserListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}

	UserListService := user.NewUserListService()
	list, total, err := UserListService.GetUserList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error for get userList"))
	}
	return c.JSON(http.StatusOK, response.OK(response.UserListRes{
		list,
		total,
	}))

}
