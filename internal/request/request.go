package request

import (
	"errors"
	"regexp"
	"strings"
)

type PageReq struct {
	Page     int `json:"page" query:"page"`
	PageSize int `json:"page_size" query:"page_size"`
}

func (r *PageReq) SetDefault() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PageSize <= 0 {
		r.PageSize = 10
	}
	if r.PageSize > 50 {
		r.PageSize = 50
	}
}

var emailRe = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

func validateEmail(email string) error {
	if !emailRe.MatchString(email) {
		return errors.New("邮箱格式不正确")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 || len(password) > 32 {
		return errors.New("密码长度必须为6-32")
	}
	return nil
}

func validateUsername(username string) error {
	if len(username) < 3 || len(username) > 32 {
		return errors.New("用户名长度必须为3-32")
	}
	return nil
}

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterReq) Validate() error {
	r.Username = strings.TrimSpace(r.Username)
	r.Email = strings.TrimSpace(r.Email)
	if r.Username == "" || r.Email == "" || r.Password == "" {
		return errors.New("参数不能为空")
	}
	if err := validateUsername(r.Username); err != nil {
		return err
	}
	if err := validatePassword(r.Password); err != nil {
		return err
	}
	return validateEmail(r.Email)
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginReq) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)
	if r.Email == "" || r.Password == "" {
		return errors.New("参数不能为空")
	}
	if err := validateEmail(r.Email); err != nil {
		return err
	}
	return validatePassword(r.Password)
}

type UserListReq struct {
	PageReq
}

type UpdateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Status   int    `json:"status"`
}

func (r *UpdateUserReq) Validate() error {
	r.Username = strings.TrimSpace(r.Username)
	r.Email = strings.TrimSpace(r.Email)
	if r.Username == "" || r.Email == "" {
		return errors.New("参数不能为空")
	}
	if err := validateUsername(r.Username); err != nil {
		return err
	}
	if err := validateEmail(r.Email); err != nil {
		return err
	}
	if r.Role != 0 && r.Role != 1 {
		return errors.New("role 只能是 0/1")
	}
	if r.Status != 0 && r.Status != 1 {
		return errors.New("status 只能是 0/1")
	}
	return nil
}

type ResetPasswordReq struct {
	Password string `json:"password"`
}

func (r *ResetPasswordReq) Validate() error {
	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.New("password can't be empty")
	}
	return validatePassword(r.Password)
}

type CreateArticleReq struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Summary    string `json:"summary"`
	Cover      string `json:"cover"`
	CategoryID int    `json:"category_id"`
	Tags       []int  `json:"tags"`
	Status     int    `json:"status"`
}

func (r *CreateArticleReq) Validate() error {
	r.Title = strings.TrimSpace(r.Title)
	r.Content = strings.TrimSpace(r.Content)
	if r.Title == "" || r.Content == "" {
		return errors.New("标题和内容不能为空")
	}
	if len(r.Title) > 256 {
		return errors.New("标题最长256字符")
	}
	if r.Status != 0 && r.Status != 1 && r.Status != 2 {
		return errors.New("status 只能是 0/1/2")
	}
	return nil
}

type ArticleListReq struct {
	PageReq
	Status     *int   `json:"status" query:"status"`
	CategoryID int    `json:"category_id" query:"category_id"`
	TagID      int    `json:"tag_id" query:"tag_id"`
	Keyword    string `json:"keyword" query:"keyword"`
}

func (r *ArticleListReq) SetDefault() {
	r.PageReq.SetDefault()
	r.Keyword = strings.TrimSpace(r.Keyword)
}

type CreateCategoryReq struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CategoryListReq struct {
	PageReq
}

type CreateTagReq struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type TagListReq struct {
	PageReq
}

type CreateRoleReq struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type RoleListReq struct {
	PageReq
}

type CreateMenuReq struct {
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Sort     int    `json:"sort"`
	Status   int    `json:"status"`
}

type MenuListReq struct {
	PageReq
}
type LogListReq struct {
	PageReq
	UserID int `json:"user_id" query:"user_id"`
}

func (r *LogListReq) Validate() error {
	r.PageReq.SetDefault()
	if r.UserID < 0 {
		return errors.New("user_id < 0")
	}
	return nil
}

type ChangeUserStatusReq struct {
	Status int `json:"status"`
}

func (r *ChangeUserStatusReq) Validate() error {
	if r.Status != 0 && r.Status != 1 {
		return errors.New("status 只能是0 /1")
	}
	return nil
}
