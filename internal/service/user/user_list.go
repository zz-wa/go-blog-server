package user

import (
	"blog_r/internal/repository/user"
	"blog_r/internal/request"
	"blog_r/internal/response"
	"errors"
)

type UserListService struct{}

func NewUserListService() *UserListService {
	return &UserListService{}
}

func (s *UserListService) GetUserList(req *request.UserListReq) ([]response.UserItem, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}
	req.SetDefault()
	list, total, err := user.GetUserList(req.Page, req.PageSize)
	if err != nil {
		return nil, total, errors.New("fail to get user list")
	}
	var items []response.UserItem
	for _, u := range list {
		items = append(items, response.UserItem{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			Status:    u.Status,
			Role:      u.Role,
			CreatedAt: u.CreatedAt, 
		})
	}
	return items, total, nil
}
