package api

import (
	"blog_r/internal/model"
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/article"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func CreateArticle(c *echo.Context) error {
	req := new(request.CreateArticleReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, err.Error()))
	}
	articleService := article.NewArticleService()
	if err := articleService.CreateArticle(req); err != nil {
		return c.JSON(500, response.Fail(response.CodeServerError, "服务器内部错误"))
	}
	return c.JSON(200, response.OK(nil))
}

// ArticleList 公开接口：默认只返回已发布(status=1)的文章
func ArticleList(c *echo.Context) error {
	req := new(request.ArticleListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	if req.Status == nil {
		one := 1
		req.Status = &one
	}
	articleService := article.NewArticleService()
	list, total, err := articleService.GetArticleList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error for get details"))
	}
	return c.JSON(http.StatusOK, response.OK(response.ArticleListRes{
		List:  list,
		Total: total,
	}))
}

// AdminArticleList 管理员接口：不传 status 则返回全部文章
func AdminArticleList(c *echo.Context) error {
	req := new(request.ArticleListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	articleService := article.NewArticleService()
	list, total, err := articleService.GetArticleList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error for get details"))
	}
	return c.JSON(http.StatusOK, response.OK(response.ArticleListRes{
		List:  list,
		Total: total,
	}))
}

func ArticleDetail(c *echo.Context) error {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid id number"))
	}
	articleService := article.NewArticleService()
	articleDetail, err := articleService.GetArticleByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error"))
	}
	return c.JSON(http.StatusOK, response.OK(response.ArticleDetailRes{
		Article: articleDetail,
	}))
}

func UpdateArticle(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "id error"))
	}
	req := new(request.CreateArticleReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, err.Error()))
	}

	updateA := &model.Article{
		Title:      req.Title,
		Content:    req.Content,
		Summary:    req.Summary,
		Cover:      req.Cover,
		CategoryID: req.CategoryID,
		Status:     req.Status,
	}

	articleService := article.NewArticleService()
	err = articleService.UpdateArticle(id, updateA, req.Tags)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to update article"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func DeleteArticle(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	articleService := article.NewArticleService()
	err = articleService.DeleteArticle(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to delete article"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}
