package model

type Menu struct {
	Model
	ParentID int    `json:"parent_id" gorm:"type:int;default:0"`
	Name     string `json:"name" gorm:"type:varchar(64)"`
	Path     string `json:"path" gorm:"uniqueIndex;type:varchar(255)"`
	Sort     int    `json:"sort" gorm:"type:int;default:0"`
	Status   int    `json:"status" gorm:"type:int;default:1"`
}
