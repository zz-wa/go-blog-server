package user

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	"errors"
)

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) CreateUser(user *model.User) error {
	return global.DB.Create(user).Error
}

func (r *Repo) GetByUsername(name string) (model.User, error) {
	user := model.User{}
	db := global.DB.Where("username=?", name).First(&user)
	if db.Error != nil {
		return model.User{}, db.Error
	}
	return user, nil
}

func (r *Repo) GetByEmail(email string) (model.User, error) {
	user := model.User{}
	db := global.DB.Where("email=?", email).First(&user)
	if db.Error != nil {
		return model.User{}, db.Error
	}
	return user, nil
}

func (r *Repo) GetByID(id int) (model.User, error) {
	user := model.User{}
	db := global.DB.Where("id=?", id).First(&user)
	if db.Error != nil {
		return model.User{}, db.Error
	}
	return user, nil
}

func (r *Repo) UpdateUser(id int, req *model.User) error {
	if id <= 0 || req == nil {
		return errors.New("invalid request")
	}
	updates := map[string]interface{}{
		"username": req.Username,
		"email":    req.Email,
		"role":     req.Role,
		"status":   req.Status,
	}
	return global.DB.Model(&model.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *Repo) ResetPassword(id int, hashedPassword string) error {
	if id <= 0 || hashedPassword == "" {
		return errors.New("invalid request")
	}
	db := global.DB.Model(&model.User{}).Where("id = ?", id).Update("password", hashedPassword)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *Repo) UpdateUserStatus(id int, status int) error {
	if id <= 0 {
		return errors.New("invalid request")
	}
	if status != 0 && status != 1 {
		return errors.New("invalid status")
	}
	db := global.DB.Model(&model.User{}).Where("id = ?", id).Update("status", status)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
