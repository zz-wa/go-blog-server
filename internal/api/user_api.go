package api

import (
	user2 "blog_r/internal/repository/user"
	"blog_r/internal/request"
	"blog_r/internal/response"
	"blog_r/internal/service/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func Register(c *echo.Context) error {
	req := new(request.RegisterReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误 "))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, err.Error()))
	}

	user := user.NewUserService(user2.NewRepo())
	if err := user.Register(req); err != nil {
		if err.Error() == "username already exists" || err.Error() == "email already exists" {
			return c.JSON(http.StatusConflict, response.Fail(response.CodeConflict, err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "服务器内部错误"))
	}

	return c.JSON(http.StatusOK, response.OK(nil))
}

func Login(c *echo.Context) error {
	ip := c.RealIP()

	req := new(request.LoginReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误 "))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, err.Error()))
	}
	userServeice := user.NewUserService(user2.NewRepo())
	token, exp, err := userServeice.Login(req, ip)
	if err != nil {
		if err.Error() == "email not found" || err.Error() == "password error" {
			return c.JSON(http.StatusUnauthorized, response.Fail(response.CodeUnauthorized, "邮箱或密码错误"))
		}
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "服务器内部错误"))
	}
	return c.JSON(http.StatusOK, response.OK(response.LoginRes{
		Token:     token,
		ExpiredAt: exp,
	}))
}

func Profile(c *echo.Context) error {
	raw := c.Get("user_id")
	userID, ok := raw.(int)
	if !ok || userID <= 0 {
		return c.JSON(http.StatusUnauthorized, response.Fail(response.CodeUnauthorized, "未登录"))
	}
	userService := user.NewUserService(user2.NewRepo())
	profile, err := userService.Profile(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "服务器内部错误"))
	}
	return c.JSON(http.StatusOK, response.OK(response.ProfileRes{
		ID:        profile.ID,
		Username:  profile.Username,
		Email:     profile.Email,
		CreatedAt: profile.CreatedAt,
	}))
}

func UpdateUser(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid id"))
	}

	req := new(request.UpdateUserReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, err.Error()))
	}

	userService := user.NewUserService(user2.NewRepo())
	if err := userService.UpdateUser(id, req); err != nil {
		if err.Error() == "username already exists" || err.Error() == "email already exists" {
			return c.JSON(http.StatusConflict, response.Fail(response.CodeConflict, err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "服务器内部错误"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func ResetPassword(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid id"))
	}

	req := new(request.ResetPasswordReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, err.Error()))
	}

	userService := user.NewUserService(user2.NewRepo())
	if err := userService.ResetPassword(id, req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "服务器内部错误"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}

func ChangeUserStatus(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "invalid id"))
	}
	req := new(request.ChangeUserStatusReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, "参数错误"))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(response.CodeBadRequest, err.Error()))
	}
	userService := user.NewUserService(user2.NewRepo())
	if err := userService.ChangeUserStatus(id, req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Fail(response.CodeServerError, "服务器内部错误"))
	}
	return c.JSON(http.StatusOK, response.OK(nil))
}
