package category

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
)

func CreateCategory(cat *model.Category) error {
	return global.DB.Create(cat).Error
}

func GetCategoryByID(id int) (model.Category, error) {
	category := model.Category{}
	db := global.DB.Where("id=?", id).First(&category)
	if db.Error != nil {
		return model.Category{}, db.Error
	}
	return category, nil
}
func GetCategoryList(page, pageSize int) ([]model.Category, int64, error) {
	categories, total, err := global.GetList[model.Category](page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}

func UpdateCategory(cat *model.Category) error {
	return global.DB.Save(cat).Error
}

func DeleteCategory(id int) error {
	return global.DB.Where("id=?", id).Delete(&model.Category{}).Error
}
