package menu

import (
	"blog_r/internal/model"
	menuRepo "blog_r/internal/repository/menu"
	"blog_r/internal/request"
	"errors"
)

type MenuService struct {
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (s *MenuService) CreateMenu(req *request.CreateMenuReq) error {
	if req == nil {
		return errors.New("invalid request")
	}
	if req.Name == "" || req.Path == "" {
		return errors.New("name or path is empty")
	}
	if _, err := menuRepo.GetMenuByPath(req.Path); err == nil {
		return errors.New("menu path already exists")
	}
	newMenu := &model.Menu{
		ParentID: req.ParentID,
		Name:     req.Name,
		Path:     req.Path,
		Sort:     req.Sort,
		Status:   req.Status,
	}
	if err := menuRepo.CreateMenu(newMenu); err != nil {
		return errors.New("fail to create menu")
	}
	return nil
}

func (s *MenuService) GetMenuByID(id int) (model.Menu, error) {
	if id <= 0 {
		return model.Menu{}, errors.New("invalid menu id")
	}
	menu, err := menuRepo.GetMenuByID(id)
	if err != nil {
		return model.Menu{}, errors.New("menu not found")
	}
	return menu, nil
}

func (s *MenuService) GetMenuList(req *request.MenuListReq) ([]model.Menu, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}
	req.SetDefault()
	list, total, err := menuRepo.GetMenuList(req.Page, req.PageSize)
	if err != nil {
		return nil, total, errors.New("fail to get menu list")
	}
	return list, total, nil
}

func (s *MenuService) UpdateMenu(id int, newMenu *model.Menu) error {
	if id <= 0 || newMenu == nil {
		return errors.New("invalid request")
	}
	newMenu.ID = id
	if err := menuRepo.UpdateMenu(newMenu); err != nil {
		return errors.New("fail to update menu")
	}
	return nil
}

func (s *MenuService) DeleteMenu(id int) error {
	if id <= 0 {
		return errors.New("invalid menu id")
	}
	if err := menuRepo.DeleteMenu(id); err != nil {
		return errors.New("fail to delete menu")
	}
	return nil
}
