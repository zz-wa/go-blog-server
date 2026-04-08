package middleware

import (
	"blog_r/internal/global"
	"blog_r/internal/repository/user"
	"blog_r/internal/response"
	"net/http"

	"github.com/labstack/echo/v5"
)

/*
写一个 PermissionMiddleware
中间件流程只做这 4 件事：
从 context 取 user_id
根据 user_id 查用户，拿到角色（你现在是 Role int）
把角色映射成 Casbin 的 subject（比如 role=1 -> "admin"，否则 "reader"）
调 global.Enforcer.Enforce(subject, 请求路径, 请求方法)，false 就返回 403
*/
func PermissionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		if global.Enforcer == nil {
			return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "500"))
		}
		var sub string
		userID := c.Get("user_id")
		if _, ok := userID.(int); !ok {
			return c.JSON(403, response.Fail(response.CodeForbidden, "forbidden"))
		}
		users, err := user.NewRepo().GetByID(userID.(int))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get user data"))
		}
		role := users.Role
		if role == 1 {
			sub = "admin"
		} else {
			sub = "reader"

		}
		obj := c.Request().URL.Path
		act := c.Request().Method
		ok, err := global.Enforcer.Enforce(sub, obj, act)
		if err != nil {
			return c.JSON(500, response.Fail(response.CodeServerError, "serverError"))
		}
		if !ok {
			return c.JSON(http.StatusForbidden, response.Fail(response.CodeForbidden, "forbidden"))
		}
		return next(c)
	}
}
