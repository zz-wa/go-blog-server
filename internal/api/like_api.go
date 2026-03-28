package api

import (
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/like"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func Like(c *echo.Context) error {
	req := new(request.SetORUndoLikeReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}

	req.UserID = c.Get("user_id").(int)
	if err := req.Validate(); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, err.Error()))
	}

	LikeTypeStr := c.Param("like_type")
	LikeType, err := strconv.Atoi(LikeTypeStr)
	if err != nil || LikeType < 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "param article id error"))
	}

	idStr := c.Param("target_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "param article id error"))
	}
	req.LikeType = LikeType
	req.TargetID = id
	LikeService := like.NewLikeService()
	if err := LikeService.ToggleLike(req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "can't like or undo like"))
	}

	return c.JSON(http.StatusOK, response.OK(nil))

}
