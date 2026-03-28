package config

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	"errors"

	"gorm.io/gorm"
)

//Config表--字典。他现在是空的。什么都没有，但是要插入相关的配置之后才能进行具体数据的读取修改
//key--字典的词条名
//value--词条内容
//desc --词条的解释

//所以这里的函数是一个大家都能用的函数。不针对某一个具体的固定项，取决于传进去的key!

func GetConfigList() ([]model.Config, error) {
	var configs []model.Config
	db := global.DB.Model(&model.Config{})
	if err := db.Find(&configs).Error; err != nil {
		return []model.Config{}, err
	}
	return configs, nil
}

func GetConfigByKey(key string) (model.Config, error) {
	var config model.Config
	db := global.DB.Model(&model.Config{})
	if err := db.Where("key=?", key).First(&config).Error; err != nil {
		return model.Config{}, err
	}
	return config, nil
}

func UpdateConfig(key, value string) error {
	if key == "" {
		return errors.New("invalid request")
	}
	db := global.DB.Model(&model.Config{})
	result := db.Where("key=?", key).Update("value", value)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
