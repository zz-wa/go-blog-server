package operation_log

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	"errors"
)

func InsertOperationLog(log *model.OperationLog) error {
	return global.DB.Create(log).Error
}

func ListOperationLog(page, pageSize int, userID *int) ([]model.OperationLog, int64, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, 0, errors.New("invalid pagination parameters")
	}
	var logs []model.OperationLog
	var total int64
	db := global.DB.Model(&model.OperationLog{})
	if userID != nil && *userID > 0 {
		db = db.Where("user_id = ?", *userID)
	}
	db.Count(&total)
	offset := (page - 1) * pageSize
	err := db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error
	if err != nil {
		return nil, total, err
	}
	return logs, total, nil
}
