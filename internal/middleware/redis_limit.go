package middleware

import (
	"blog_r/internal/global"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

func RateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := context.Background()
		ip := c.RealIP()
		key := "rate:login:" + ip
		count, err := global.Redis.Incr(ctx, key).Result()
		if err != nil {
			return err
		}
		if count == 1 {
			if err := global.Redis.Expire(ctx, key, time.Minute).Err(); err != nil {
				return err
			}
		}
		if count > 5 {
			return c.JSON(http.StatusTooManyRequests, map[string]string{"message": "too many requests"})
		}
		return next(c)
	}
}
