package model

type Role struct {
	Model
	Name string `json:"name" gorm:"uniqueIndex;type:varchar(64)"`
	Desc string `json:"desc" gorm:"type:varchar(256)"`
}
