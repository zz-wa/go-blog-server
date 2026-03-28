package category

import (
	"blog_r/internal/model"
	categoryRepo "blog_r/internal/repository/category"
	"blog_r/internal/request"
	"errors"
)

type CategoryService struct {
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) CreateCategory(req *request.CreateCategoryReq) error {
	if req == nil {
		return errors.New("invalid request")
	}
	if err := req.Validate(); err != nil {
		return err
	}
	SetCategory := &model.Category{}
	SetCategory.Name = req.Name
	SetCategory.Desc = req.Desc

	if err := categoryRepo.CreateCategory(SetCategory); err != nil {
		return errors.New("fail to create category")
	}
	return nil
}

func (s *CategoryService) GetCategoryByID(id int) (model.Category, error) {
	if id <= 0 {
		return model.Category{}, errors.New("invalid category id")
	}
	Category, err := categoryRepo.GetCategoryByID(id)
	if err != nil {
		return model.Category{}, errors.New("fail to get category by id")
	}
	return Category, nil
}

func (s *CategoryService) GetCategoryList(req *request.CategoryListReq) ([]model.Category, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}
	req.SetDefault()

	CatrgoryList, total, err := categoryRepo.GetCategoryList(req.Page, req.PageSize)
	if err != nil {
		return nil, total, errors.New("fail to get category list")
	}
	return CatrgoryList, total, nil
}

func (s *CategoryService) UpdateCategory(id int, NewCategory *model.Category) error {
	if id <= 0 || NewCategory == nil {
		return errors.New("invalid request")
	}
	NewCategory.ID = id
	err := categoryRepo.UpdateCategory(NewCategory)
	if err != nil {
		return errors.New("failed to update category")
	}
	return nil
}

func (s *CategoryService) DeleteCategory(id int) error {
	if id <= 0 {
		return errors.New("invalid category ID")
	}
	err := categoryRepo.DeleteCategory(id)
	if err != nil {
		return errors.New("failed to delete category")
	}
	return nil
}
