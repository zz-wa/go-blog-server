package user

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	userRepo "blog_r/internal/repository/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InitAdmin(conf *global.Config) error {
	if conf.Admin.Email == "" || conf.Admin.Password == "" {
		return nil
	}

	// 如果管理员已存在（邮箱或用户名），不再创建
	if _, err := userRepo.GetByEmail(conf.Admin.Email); err == nil {
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if conf.Admin.Username != "" {
		if _, err := userRepo.GetByUsername(conf.Admin.Username); err == nil {
			return nil
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(conf.Admin.Password), 14)
	if err != nil {
		return err
	}

	admin := model.User{
		Username: conf.Admin.Username,
		Email:    conf.Admin.Email,
		Password: string(hashed),
		Role:     1, // 管理员
		Status:   1,
	}

	return userRepo.CreateUser(&admin)
}
