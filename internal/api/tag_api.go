package api

import (
	"blog_r/internal/model"
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/tag"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func CreateTag(c *echo.Context) error {
	req := new(request.CreateTagReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	tagService := tag.NewTagService()
	if err := tagService.CreateTag(req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "服务器错误"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func TagList(c *echo.Context) error {
	req := new(request.TagListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}

	tagService := tag.NewTagService()
	list, total, err := tagService.GetTagList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error for get details"))
	}
	return c.JSON(http.StatusOK, response.OK(response.TagListRes{
		List:  list,
		Total: total,
	}))
}

func TagDetail(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request for id"))
	}
	tagService := tag.NewTagService()
	detail, err := tagService.GetTagByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get tag detail"))
	}
	return c.JSON(http.StatusOK, response.OK(detail))
}

func UpdateTag(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request for id"))
	}
	req := new(request.CreateTagReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	updateT := &model.Tag{
		Name:  req.Name,
		Color: req.Color,
	}
	tagService := tag.NewTagService()
	err = tagService.UpdateTag(id, updateT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to update tag"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func DeleteTag(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	tagService := tag.NewTagService()
	err = tagService.DeleteTag(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to delete tag"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}
