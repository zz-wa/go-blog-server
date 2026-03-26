package tag

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
)

func GetByIDs(ids []int) ([]model.Tag, error) {
	if len(ids) == 0 {
		return []model.Tag{}, nil
	}
	var tags []model.Tag
	if err := global.DB.Where("id IN ?", ids).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func CreateTag(tag *model.Tag) error {
	return global.DB.Create(tag).Error
}

func GetTagByID(id int) (model.Tag, error) {
	tag := model.Tag{}
	db := global.DB.Where("id=?", id).First(&tag)
	if db.Error != nil {
		return model.Tag{}, db.Error
	}
	return tag, nil

}
func GetTagList(page, pageSize int) ([]model.Tag, int64, error) {
	tags, total, err := global.GetList[model.Tag](page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return tags, total, nil
}
func UpdateTag(tag *model.Tag) error {
	return global.DB.Save(tag).Error
}

func DeleteTag(id int) error {
	return global.DB.Where("id=?", id).Delete(&model.Tag{}).Error
}
