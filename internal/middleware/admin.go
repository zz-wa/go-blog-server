package middleware

import (
	"blog_r/internal/repository/user"
	"blog_r/internal/response"

	"github.com/labstack/echo/v5"
)

/*函数名：AdminMiddleware
逻辑：
从 context 取 user_id（AuthMiddleware 已经放进去了）
查数据库用户
如果 role != 1 → 返回 403
是管理员 → 放行
*/

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		userID := c.Get("user_id")
		if _, ok := userID.(int); !ok {
			return c.JSON(403, response.Fail(response.CodeForbidden, "forbidden"))
		}
		users, err := user.GetByID(userID.(int))
		if err != nil {
			return c.JSON(500, response.Fail(response.CodeServerError, "服务器内部错误"))
		}
		if users.Role != 1 {
			return c.JSON(403, response.Fail(response.CodeForbidden, "forbidden"))
		}

		return next(c)

	}
}
