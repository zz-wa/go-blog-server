package model

type OperationLog struct {
	Model
	UserID int    `json:"user_id" gorm:"type:int"`
	Method string `json:"method" gorm:"type:varchar(16)"`
	Path   string `json:"path" gorm:"type:varchar(512)"`
	Body   string `json:"body" gorm:"type:text"`
	Status int    `json:"status" gorm:"type:int"`
}
