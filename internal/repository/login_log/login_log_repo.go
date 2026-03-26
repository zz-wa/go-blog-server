package login_log

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	"errors"
)

/*
再写 List（分页 + 可选过滤）
List 有三个关键点：分页、可选过滤、返回总数。
4.1 参数
page、pageSize
userID（建议用指针，表示“可选”）
4.2 逻辑流程
按这个顺序写逻辑：
校验分页参数
page <= 0 或 pageSize <= 0 就直接返回错误
初始化查询对象
基于 LoginLog 模型建立 db 查询对象
注意：后面 Count 和 Find 都用这个 db，保证条件一致
可选过滤
如果 userID 不为 nil 且 > 0
加 Where("user_id = ?", *userID)
统计总数
Count(&total)
一定在 Offset/Limit 之前执行
分页参数
offset = (page - 1) * pageSize
排序与查询
按 created_at DESC 排序（最新在前）
Offset + Limit 做分页
Find(&logs)
返回
返回 logs、total、error
*/
func InsertLoginLog(log *model.LoginLog) error {
	return global.DB.Create(log).Error
}

func ListLoginLog(page, pageSize int, UserID *int) ([]model.LoginLog, int64, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, 0, errors.New("invalid pagination parameters")
	}
	var loginLogs []model.LoginLog
	var total int64
	db := global.DB.Model(&model.LoginLog{})
	if UserID != nil && *UserID > 0 {
		db = db.Where("user_id = ?", *UserID)
	}
	db.Count(&total)
	offset := (page - 1) * pageSize
	err := db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&loginLogs).Error
	if err != nil {
		return nil, total, err
	}
	return loginLogs, total, nil

}
