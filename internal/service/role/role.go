package role

import (
	"blog_r/internal/model"
	roleRepo "blog_r/internal/repository/role"
	"blog_r/internal/request"
	"errors"
)

type RoleService struct {
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (s *RoleService) CreateRole(req *request.CreateRoleReq) error {
	if req == nil {
		return errors.New("invalid request")
	}
	if req.Name == "" {
		return errors.New("name is empty")
	}
	if _, err := roleRepo.GetRoleByName(req.Name); err == nil {
		return errors.New("role already exists")
	}
	newRole := &model.Role{
		Name: req.Name,
		Desc: req.Desc,
	}
	if err := roleRepo.CreateRole(newRole); err != nil {
		return errors.New("fail to create role")
	}
	return nil
}

func (s *RoleService) GetRoleByID(id int) (model.Role, error) {
	if id <= 0 {
		return model.Role{}, errors.New("invalid role id")
	}
	role, err := roleRepo.GetRoleByID(id)
	if err != nil {
		return model.Role{}, errors.New("role not found")
	}
	return role, nil
}

func (s *RoleService) GetRoleList(req *request.RoleListReq) ([]model.Role, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}
	req.SetDefault()
	list, total, err := roleRepo.GetRoleList(req.Page, req.PageSize)
	if err != nil {
		return nil, total, errors.New("fail to get role list")
	}
	return list, total, nil
}

func (s *RoleService) UpdateRole(id int, newRole *model.Role) error {
	if id <= 0 || newRole == nil {
		return errors.New("invalid request")
	}
	newRole.ID = id
	if err := roleRepo.UpdateRole(newRole); err != nil {
		return errors.New("fail to update role")
	}
	return nil
}

func (s *RoleService) DeleteRole(id int) error {
	if id <= 0 {
		return errors.New("invalid role id")
	}
	if err := roleRepo.DeleteRole(id); err != nil {
		return errors.New("fail to delete role")
	}
	return nil
}
