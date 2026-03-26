package api

import (
	"blog_r/internal/model"
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/role"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func CreateRole(c *echo.Context) error {
	req := new(request.CreateRoleReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	roleService := role.NewRoleService()
	if err := roleService.CreateRole(req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, err.Error()))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func RoleList(c *echo.Context) error {
	req := new(request.RoleListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	roleService := role.NewRoleService()
	list, total, err := roleService.GetRoleList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get role list"))
	}
	return c.JSON(http.StatusOK, response.OK(response.RoleListRes{
		List:  list,
		Total: total,
	}))
}

func RoleDetail(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request for id"))
	}
	roleService := role.NewRoleService()
	detail, err := roleService.GetRoleByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get role detail"))
	}
	return c.JSON(http.StatusOK, response.OK(detail))
}

func UpdateRole(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request for id"))
	}
	req := new(request.CreateRoleReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	updateR := &model.Role{
		Name: req.Name,
		Desc: req.Desc,
	}
	roleService := role.NewRoleService()
	if err := roleService.UpdateRole(id, updateR); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to update role"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func DeleteRole(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	roleService := role.NewRoleService()
	if err := roleService.DeleteRole(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to delete role"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}
