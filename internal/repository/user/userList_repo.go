package user

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
)

func GetUserList(page, pageSize int) ([]model.User, int64, error) {
	users, total, err := global.GetList[model.User](page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
