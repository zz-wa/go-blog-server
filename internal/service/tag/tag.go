package tag

import (
	"blog_r/internal/model"
	tagRepo "blog_r/internal/repository/tag"
	"blog_r/internal/request"
	"errors"
)

type TagService struct {
}

func NewTagService() *TagService {
	return &TagService{}
}

func (s *TagService) CreateTag(req *request.CreateTagReq) error {
	if req == nil {
		return errors.New("invalid request")
	}
	if req.Name == "" {
		return errors.New("name is empty")
	}
	newTag := &model.Tag{
		Name:  req.Name,
		Color: req.Color,
	}
	if err := tagRepo.CreateTag(newTag); err != nil {
		return errors.New("fail to create tag")
	}
	return nil
}

func (s *TagService) GetTagByID(id int) (model.Tag, error) {
	if id <= 0 {
		return model.Tag{}, errors.New("invalid tag id")
	}
	tag, err := tagRepo.GetTagByID(id)
	if err != nil {
		return model.Tag{}, errors.New("fail to get tag by id")
	}
	return tag, nil
}

func (s *TagService) GetTagList(req *request.TagListReq) ([]model.Tag, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}
	req.SetDefault()

	list, total, err := tagRepo.GetTagList(req.Page, req.PageSize)
	if err != nil {
		return nil, total, errors.New("fail to get tag list")
	}
	return list, total, nil
}

func (s *TagService) UpdateTag(id int, newTag *model.Tag) error {
	if id <= 0 || newTag == nil {
		return errors.New("invalid request")
	}
	newTag.ID = id
	if err := tagRepo.UpdateTag(newTag); err != nil {
		return errors.New("failed to update tag")
	}
	return nil
}

func (s *TagService) DeleteTag(id int) error {
	if id <= 0 {
		return errors.New("invalid tag id")
	}
	if err := tagRepo.DeleteTag(id); err != nil {
		return errors.New("failed to delete tag")
	}
	return nil
}
