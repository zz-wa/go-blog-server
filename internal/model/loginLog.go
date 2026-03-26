package model

type LoginLog struct {
	Model
	UserID  int    `json:"user_id" gorm:"type:int"`
	Email   string `json:"email" gorm:"type:varchar(64)"`
	IP      string `json:"ip" gorm:"type:varchar(64)"`
	Success bool   `json:"success" gorm:"type:boolean"`
	Msg     string `json:"msg" gorm:"type:varchar(256)"`
}
