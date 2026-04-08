package comment

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	"errors"
)

type Repo struct{}

func NewRepo() *Repo { return &Repo{} }

func (r *Repo) CreateComment(comment *model.Comment) error {
	return global.DB.Create(comment).Error
}

func (r *Repo) GetCommentList(articleID, page, pageSize int) ([]model.Comment, int64, error) {
	db := global.DB.Model(&model.Comment{})
	var comments []model.Comment
	if page <= 0 || pageSize <= 0 {
		return nil, 0, errors.New("invalid pagination parameters")
	}
	var total int64
	db.Where("article_id=?", articleID).Count(&total)
	offest := (page - 1) * pageSize
	db = db.Where("article_id=?", articleID).Offset(offest).Limit(pageSize).Find(&comments)
	if db.Error != nil {
		return nil, total, db.Error
	}
	return comments, total, nil
}

func (r *Repo) DeleteComment(id int) error {
	return global.DB.Where("id=?", id).Delete(&model.Comment{}).Error
}
