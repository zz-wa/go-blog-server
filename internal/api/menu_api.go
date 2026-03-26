package api

import (
	"blog_r/internal/model"
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/menu"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func CreateMenu(c *echo.Context) error {
	req := new(request.CreateMenuReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	menuService := menu.NewMenuService()
	if err := menuService.CreateMenu(req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, err.Error()))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func MenuList(c *echo.Context) error {
	req := new(request.MenuListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	menuService := menu.NewMenuService()
	list, total, err := menuService.GetMenuList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get menu list"))
	}
	return c.JSON(http.StatusOK, response.OK(map[string]interface{}{
		"list":  list,
		"total": total,
	}))
}

func MenuDetail(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request for id"))
	}
	menuService := menu.NewMenuService()
	detail, err := menuService.GetMenuByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get menu detail"))
	}
	return c.JSON(http.StatusOK, response.OK(detail))
}

func UpdateMenu(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request for id"))
	}
	req := new(request.CreateMenuReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	updateM := &model.Menu{
		ParentID: req.ParentID,
		Name:     req.Name,
		Path:     req.Path,
		Sort:     req.Sort,
		Status:   req.Status,
	}
	menuService := menu.NewMenuService()
	if err := menuService.UpdateMenu(id, updateM); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to update menu"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func DeleteMenu(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	menuService := menu.NewMenuService()
	if err := menuService.DeleteMenu(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to delete menu"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}
