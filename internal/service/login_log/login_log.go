package login_log

import (
	"blog_r/internal/model"
	"blog_r/internal/repository/login_log"
	"blog_r/internal/request"
	"errors"
)

type LoginLogService struct {
}

func NewLoginLogService() *LoginLogService {
	return &LoginLogService{}
}

func (s *LoginLogService) LoginLogList(req *request.LogListReq) ([]model.LoginLog, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}
	var filterUserID *int
	if req.UserID > 0 {
		filterUserID = &req.UserID
	}
	loginList, total, err := login_log.ListLoginLog(req.Page, req.PageSize, filterUserID)
	if err != nil {
		return nil, total, err
	}
	return loginList, total, nil
}
