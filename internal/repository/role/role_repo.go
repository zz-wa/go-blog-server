package role

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
)

func CreateRole(role *model.Role) error {
	return global.DB.Create(role).Error
}

func GetRoleByID(id int) (model.Role, error) {
	role := model.Role{}
	db := global.DB.Where("id=?", id).First(&role)
	if db.Error != nil {
		return model.Role{}, db.Error
	}
	return role, nil
}

func GetRoleByName(name string) (model.Role, error) {
	role := model.Role{}
	db := global.DB.Where("name=?", name).First(&role)
	if db.Error != nil {
		return model.Role{}, db.Error
	}
	return role, nil
}

func GetRoleList(page, pageSize int) ([]model.Role, int64, error) {
	roles, total, err := global.GetList[model.Role](page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return roles, total, nil
}

func UpdateRole(role *model.Role) error {
	return global.DB.Save(role).Error
}

func DeleteRole(id int) error {
	return global.DB.Where("id=?", id).Delete(&model.Role{}).Error
}
