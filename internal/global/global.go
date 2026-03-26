package global

import (
	"errors"

	"github.com/casbin/casbin/v3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
	Redis    *redis.Client
)

func GetList[T any](page, pageSize int) ([]T, int64, error) {
	var results []T
	var total int64
	if page <= 0 || pageSize <= 0 {
		return nil, 0, errors.New("page 和 pageSize 必须大于 0")
	}
	DB.Model(new(T)).Count(&total)
	offset := (page - 1) * pageSize
	if err := DB.Offset(offset).Limit(pageSize).Find(&results).Error; err != nil {
		return nil, total, err
	}

	return results, total, nil

}
