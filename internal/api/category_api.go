package api

import (
	"blog_r/internal/model"
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/category"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func CreateCategory(c *echo.Context) error {
	req := new(request.CreateCategoryReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	// 修改：移除无效注释块，避免干扰阅读
	categoryService := category.NewCategoryService()
	if err := categoryService.CreateCategory(req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "服务器错误"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func CategoryList(c *echo.Context) error {
	req := new(request.CategoryListReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}

	categoryService := category.NewCategoryService()
	list, total, err := categoryService.GetCategoryList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "error for get details"))
	}
	return c.JSON(http.StatusOK, response.OK(response.CategoryListRes{
		List:  list,
		Total: total,
	}))
}

func UpdateCategory(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 { // 修改：补充 id<=0 校验
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request for id"))
	}
	req := new(request.CreateCategoryReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(400, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	UpdateC := &model.Category{
		Name: req.Name,
		Desc: req.Desc,
	}
	categorySevice := category.NewCategoryService()
	err = categorySevice.UpdateCategory(id, UpdateC)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to update category")) // 修改：错误信息
	}
	return c.JSON(http.StatusOK, response.OK(nil))

}

func DeleteCategory(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request"))
	}
	categorySevice := category.NewCategoryService()
	err = categorySevice.DeleteCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to delete category")) // 修改：错误信息
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func CategoryDetail(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid request for id"))
	}
	// 修改：新增分类详情接口
	categoryService := category.NewCategoryService()
	detail, err := categoryService.GetCategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "fail to get category detail"))
	}
	return c.JSON(http.StatusOK, response.OK(detail))
}
