package menu

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
)

func CreateMenu(menu *model.Menu) error {
	return global.DB.Create(menu).Error
}

func GetMenuByID(id int) (model.Menu, error) {
	menu := model.Menu{}
	db := global.DB.Where("id=?", id).First(&menu)
	if db.Error != nil {
		return model.Menu{}, db.Error
	}
	return menu, nil
}

func GetMenuByPath(path string) (model.Menu, error) {
	menu := model.Menu{}
	db := global.DB.Where("path=?", path).First(&menu)
	if db.Error != nil {
		return model.Menu{}, db.Error
	}
	return menu, nil
}

func GetMenuList(page, pageSize int) ([]model.Menu, int64, error) {
	menus, total, err := global.GetList[model.Menu](page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return menus, total, nil
}

func UpdateMenu(menu *model.Menu) error {
	return global.DB.Save(menu).Error
}

func DeleteMenu(id int) error {
	return global.DB.Where("id=?", id).Delete(&model.Menu{}).Error
}
