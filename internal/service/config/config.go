package config

import (
	"blog_r/internal/model"
	"blog_r/internal/repository/config"
	"blog_r/internal/request"
)

type ConfigService struct {
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (s *ConfigService) GetConfigList() ([]model.Config, error) {
	var configs []model.Config
	configs, err := config.GetConfigList()
	if err != nil {
		return []model.Config{}, err
	}
	return configs, nil
}

func (s *ConfigService) GetConfigAbout(key string) (model.Config, error) {
	var Config model.Config

	Config, err := config.GetConfigByKey(key)
	if err != nil {
		return model.Config{}, err
	}
	return Config, nil

}

func (s *ConfigService) UpdateConfig(req request.UpdateConfigReq) error {
	err := req.Validate()
	if err != nil {
		return err
	}
	err = config.UpdateConfig(req.Key, req.Value)
	if err != nil {
		return err
	}
	return nil
}
