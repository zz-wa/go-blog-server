package router

import (
	"blog_r/internal/api"
	"blog_r/internal/middleware"

	"github.com/labstack/echo/v5"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerMiddleware)
	e.Static("/uploads", "uploads")

	e.GET("/health", api.Health)
	v1 := e.Group("/api/v1")
	public := v1.Group("/public")

	public.POST("/register", api.Register)
	public.POST("/login", api.Login, middleware.RateLimitMiddleware)
	public.GET("/articles", api.ArticleList)
	public.GET("/articles/:id", api.ArticleDetail)

	userGroup := v1.Group("/user", middleware.AuthMiddleware, middleware.PermissionMiddleware)
	userGroup.GET("/profile", api.Profile)
	userGroup.POST("/upload", api.Upload)

	adminGroup := v1.Group("/admin", middleware.AuthMiddleware, middleware.AdminMiddleware, middleware.PermissionMiddleware, middleware.OperationLogMiddleware)
	adminGroup.GET("/articles", api.AdminArticleList)
	adminGroup.GET("/articles/:id", api.ArticleDetail)
	adminGroup.POST("/articles", api.CreateArticle)
	adminGroup.PUT("/articles/:id", api.UpdateArticle)
	adminGroup.DELETE("/articles/:id", api.DeleteArticle)
	adminGroup.POST("/categories", api.CreateCategory)
	adminGroup.GET("/categories", api.CategoryList)
	adminGroup.GET("/categories/:id", api.CategoryDetail)
	adminGroup.PUT("/categories/:id", api.UpdateCategory)
	adminGroup.DELETE("/categories/:id", api.DeleteCategory)
	adminGroup.POST("/tags", api.CreateTag)
	adminGroup.GET("/tags", api.TagList)
	adminGroup.GET("/tags/:id", api.TagDetail)
	adminGroup.PUT("/tags/:id", api.UpdateTag)
	adminGroup.DELETE("/tags/:id", api.DeleteTag)
	adminGroup.POST("/roles", api.CreateRole)
	adminGroup.GET("/roles", api.RoleList)
	adminGroup.GET("/roles/:id", api.RoleDetail)
	adminGroup.PUT("/roles/:id", api.UpdateRole)
	adminGroup.DELETE("/roles/:id", api.DeleteRole)
	adminGroup.POST("/menus", api.CreateMenu)
	adminGroup.GET("/menus", api.MenuList)
	adminGroup.GET("/menus/:id", api.MenuDetail)
	adminGroup.PUT("/menus/:id", api.UpdateMenu)
	adminGroup.DELETE("/menus/:id", api.DeleteMenu)
	adminGroup.GET("/login-logs", api.LoginLogList)
	adminGroup.GET("/operation-logs", api.OperationLogList)
	adminGroup.GET("/userlist", api.UserList)
	adminGroup.PUT("/users/:id", api.UpdateUser)
	adminGroup.PUT("/users/:id/password", api.ResetPassword)
	adminGroup.PUT("/users/:id/status", api.ChangeUserStatus)

	return e
}
