package dashboard

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
)

type Count struct {
	ArticleCount  int64 `json:"article_count" gorm:"type:int"`
	DraftCount    int64 `json:"draft_count" gorm:"type:int"`
	CategoryCount int64 `json:"category_count" gorm:"type:int"`
	TagCount      int64 `json:"tag_count" gorm:"type:int"`
	UserCount     int64 `json:"user_count" gorm:"type:int"`
	TotalViews    int64 `json:"total_views" gorm:"type:int"`
}
type DashboardService struct {
}

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}

func (s *DashboardService) Stats() (*Count, error) {
	count := &Count{}
	db := global.DB
	if err := db.Model(&model.Article{}).Count(&count.ArticleCount).Error; err != nil {
		return &Count{}, err
	}

	if err := db.Model(&model.Article{}).Where("status=?", 0).Count(&count.DraftCount).Error; err != nil {
		return &Count{}, err

	}

	if err := db.Model(&model.Category{}).Count(&count.CategoryCount).Error; err != nil {
		return &Count{}, err

	}
	if err := db.Model(&model.Tag{}).Count(&count.TagCount).Error; err != nil {
		return &Count{}, err

	}
	if err := db.Model(&model.User{}).Count(&count.UserCount).Error; err != nil {
		return &Count{}, err
	}

	if err := global.DB.Model(&model.Article{}).Select("COALESCE(SUM(views), 0)").Scan(&count.TotalViews).Error; err != nil {
		return &Count{}, err
	}

	return count, nil

}
