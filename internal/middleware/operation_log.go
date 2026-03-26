package middleware

import (
	"blog_r/internal/model"
	"blog_r/internal/repository/operation_log"
	"io"
	"strings"

	"github.com/labstack/echo/v5"
)

//“谁在什么时候对哪个接口提交了什么，并且结果如何

func OperationLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		method := c.Request().Method
		if method == "GET" {
			return next(c)
		}
		userID := 0
		if raw := c.Get("user_id"); raw != nil {
			switch v := raw.(type) {
			case int:
				userID = v
			case int64:
				userID = int(v)
			case float64:
				userID = int(v)
			}
		}

		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			b = []byte("")
		}
		body := string(b)
		c.Request().Body = io.NopCloser(strings.NewReader(body))
		nextErr := next(c)
		_, status := echo.ResolveResponseStatus(c.Response(), nextErr)
		log := &model.OperationLog{
			UserID: userID,
			Method: method,
			Body:   body,
			Path:   c.Path(),
			Status: status,
		}
		_ = operation_log.InsertOperationLog(log)
		return nextErr
	}
}
