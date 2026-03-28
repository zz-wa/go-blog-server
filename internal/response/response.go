package response

import (
	"blog_r/internal/model"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type RegisterRes struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleListRes struct {
	List  []model.Article `json:"list"`
	Total int64           `json:"total"`
}

type ArticleDetailRes struct {
	Article model.Article `json:"article"`
}

type CategoryListRes struct {
	List  []model.Category `json:"list"`
	Total int64            `json:"total"`
}

type TagListRes struct {
	List  []model.Tag `json:"list"`
	Total int64       `json:"total"`
}

type RoleListRes struct {
	List  []model.Role `json:"list"`
	Total int64        `json:"total"`
}

type LoginLogListRes struct {
	List  []model.LoginLog `json:"list"`
	Total int64            `json:"total"`
}

type OperationLogListRes struct {
	List  []model.OperationLog `json:"list"`
	Total int64                `json:"total"`
}

type LoginRes struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type ProfileRes struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
type UserItem struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
type UserListRes struct {
	List  []UserItem `json:"list"`
	Total int64      `json:"total"`
}

type CommentListResp struct {
	List  []model.Comment `json:"list"`
	Total int64           `json:"total"`
}

const (
	CodeOk           = 0
	CodeBadRequest   = 4001
	CodeConflict     = 4090
	CodeUnauthorized = 4010
	CodeServerError  = 5000
	CodeForbidden    = 4030
)

func OK(data interface{}) Response {
	return Response{
		Code: CodeOk,
		Msg:  "ok",
		Data: data,
	}

}
func Fail(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}

}
