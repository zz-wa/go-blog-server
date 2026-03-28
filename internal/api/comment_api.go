package api

import (
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/comment"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func CreateComment(c *echo.Context) error {
	req := new(request.CreateCommentReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, err.Error()))
	}

	userID, _ := c.Get("user_id").(int)
	req.UserID = userID

	idStr := c.Param("article_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "param article id error"))
	}
	req.ArticleID = id
	CommentService := comment.NewCommentService()
	if err := CommentService.CreateComment(req); err != nil {
		return c.JSON(500, response.Fail(response.CodeServerError, "服务器内部错误"))
	}
	return c.JSON(200, response.OK(nil))
}

func CommentList(c *echo.Context) error {
	req := new(request.CommentListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}

	idStr := c.Param("article_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "param article id error"))
	}
	req.ArticleID = id

	CommentService := comment.NewCommentService()
	commentlist, total, err := CommentService.GetCommentList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error for get details"))
	}
	return c.JSON(http.StatusOK, response.OK(response.CommentListResp{
		List:  commentlist,
		Total: total,
	}))
}

func DeleteComment(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}

	CommentService := comment.NewCommentService()
	if err := CommentService.DeleteComment(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to delete article"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}
