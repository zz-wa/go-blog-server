package article

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	"errors"
)

func CreateArticle(article *model.Article) error {
	return global.DB.Create(article).Error
}
func GetArticleByID(id int) (model.Article, error) {
	article := model.Article{}
	db := global.DB.Preload("Tags").Where("id=?", id).First(&article)
	if db.Error != nil {
		return model.Article{}, db.Error
	}
	return article, nil
}

func GetArticleList(page, pageSize int, status *int, categoryID, tagID int, keyWord string) ([]model.Article, int64, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, 0, errors.New("invalid pagination parameters")
	}
	var articles []model.Article
	var total int64
	db := global.DB.Model(&model.Article{})
	if status != nil {
		db = db.Where("status=?", *status)
	}
	if categoryID > 0 {
		db = db.Where("category_id=?", categoryID)
	}
	if keyWord != "" {
		like := "%" + keyWord + "%"
		db = db.Where("title LIKE ? OR content LIKE ?", like, like)
	}
	if tagID > 0 {
		db = db.Joins("JOIN article_tags ON article_tags.article_id = article.id").
			Where("article_tags.tag_id = ?", tagID)

	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset).Limit(pageSize).Preload("Tags").Find(&articles)
	if db.Error != nil {
		return nil, total, db.Error
	}
	return articles, total, nil
}
func UpdateArticle(article *model.Article) error {
	return global.DB.Save(article).Error
}
func DeleteArticle(id int) error {
	return global.DB.Where("id=?", id).Delete(&model.Article{}).Error
}
