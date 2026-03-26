package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		start := time.Now()
		err := next(c) //调用下一个 handler
		stop := time.Now()
		duration := stop.Sub(start).Milliseconds() //计算耗时
		method := c.Request().Method
		path := c.Path()
		status := http.StatusOK
		if err != nil {
			var sc echo.HTTPStatusCoder
			if errors.As(err, &sc) {
				status = sc.StatusCode()
			} else {
				status = http.StatusInternalServerError
			}

		} else {
			if rw, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
				status = rw.Status
			}
		}
		fmt.Printf("%s %s %d %dms\n", method, path, status, duration)

		return err
	}
}
