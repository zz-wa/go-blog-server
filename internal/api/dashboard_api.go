package api

import (
	"blog_r/internal/response"
	dashboard2 "blog_r/internal/service/dashboard"
	"net/http"

	"github.com/labstack/echo/v5"
)

func Home(c *echo.Context) error {
	svc := dashboard2.NewDashboardService()
	data, err := svc.Stats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "can't get the data"))
	}
	return c.JSON(http.StatusOK, response.OK(map[string]int64{
		"article_count":  data.ArticleCount,
		"category_count": data.CategoryCount,
		"tag_count":      data.TagCount,
		"total_views":    data.TotalViews,
	}))
}

func Dashboard(c *echo.Context) error {
	dashboard := dashboard2.NewDashboardService()
	data, err := dashboard.Stats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "can't get  the data"))
	}
	return c.JSON(http.StatusOK, response.OK(data))
}
