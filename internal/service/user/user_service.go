package user

import (
	"blog_r/internal/model"
	"blog_r/internal/pkg/jwt"
	"blog_r/internal/repository/login_log"
	userRepo "blog_r/internal/repository/user"
	"blog_r/internal/request"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(req *request.RegisterReq) error {
	if req.Password == "" || req.Username == "" || req.Email == "" {
		return errors.New("invalid request")
	}
	if _, err := userRepo.GetByUsername(req.Username); err == nil {
		return errors.New("username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if _, err := userRepo.GetByEmail(req.Email); err == nil {
		return errors.New("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	var user model.User
	user.Username = req.Username
	user.Email = req.Email
	user.Role = 0
	user.Status = 1
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(password)
	if err := userRepo.CreateUser(&user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(req *request.LoginReq, ip string) (string, time.Time, error) {

	if req.Password == "" || req.Email == "" {
		_ = login_log.InsertLoginLog(&model.LoginLog{
			UserID:  0,
			Email:   req.Email,
			IP:      ip,
			Success: false,
			Msg:     "invalid request",
		})
		return "", time.Time{}, errors.New("invalid request")
	}
	user, err := userRepo.GetByEmail(req.Email)
	if err != nil {
		_ = login_log.InsertLoginLog(&model.LoginLog{
			UserID:  0,
			Email:   req.Email,
			IP:      ip,
			Success: false,
			Msg:     "email not found",
		})
		return "", time.Time{}, errors.New("email not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		_ = login_log.InsertLoginLog(&model.LoginLog{
			UserID:  user.ID,
			Email:   req.Email,
			IP:      ip,
			Success: false,
			Msg:     "password error",
		})
		return "", time.Time{}, errors.New("password error")
	}
	token, exp, err := jwt.GenerateToken(user.ID)

	if err != nil {
		_ = login_log.InsertLoginLog(&model.LoginLog{
			UserID:  user.ID,
			Email:   req.Email,
			IP:      ip,
			Success: false,
			Msg:     err.Error(),
		})
		return "", time.Time{}, err
	}
	_ = login_log.InsertLoginLog(&model.LoginLog{
		UserID:  user.ID,
		Email:   req.Email,
		IP:      ip,
		Success: true,
		Msg:     "",
	})
	return token, exp, nil
}

func (s *UserService) Profile(id int) (model.User, error) {
	user, err := userRepo.GetByID(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, req *request.UpdateUserReq) error {
	if id <= 0 {
		return errors.New("invalid request")
	}
	if req == nil {
		return errors.New("invalid request")
	}
	if err := req.Validate(); err != nil {
		return err
	}

	user, err := userRepo.GetByID(id)
	if err != nil {
		return err
	}

	if req.Username != user.Username {
		if u, err := userRepo.GetByUsername(req.Username); err == nil && u.ID != id {
			return errors.New("username already exists")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if req.Email != user.Email {
		if u, err := userRepo.GetByEmail(req.Email); err == nil && u.ID != id {
			return errors.New("email already exists")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	update := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
		Status:   req.Status,
	}
	return userRepo.UpdateUser(id, update)
}
func (s *UserService) ResetPassword(id int, req *request.ResetPasswordReq) error {
	if id <= 0 || req == nil {
		return errors.New("invalid request")
	}
	if err := req.Validate(); err != nil {
		return err
	}
	if _, err := userRepo.GetByID(id); err != nil {
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}
	return userRepo.ResetPassword(id, string(hashed))
}

func (s *UserService) ChangeUserStatus(id int, req *request.ChangeUserStatusReq) error {
	if id <= 0 || req == nil {
		return errors.New("invalid request")
	}
	if err := req.Validate(); err != nil {
		return err
	}

	if _, err := userRepo.GetByID(id); err != nil {
		return err
	}

	return userRepo.UpdateUserStatus(id, req.Status)
}
