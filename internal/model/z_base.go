package model

import (
	"time"

	"gorm.io/gorm"
)

func MakeMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&Config{},
		&User{},
		&Article{},
		&Tag{},
		&Category{},
		&Role{},
		&Menu{},
		&LoginLog{},
		&OperationLog{},
		&Comment{})
}

type Model struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
