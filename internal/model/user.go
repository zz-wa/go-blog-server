package model

type User struct {
	Model
	Username string `gorm:"unique;type:varchar(256)" json:"username"`
	Email    string `gorm:"unique;type:varchar(256)" json:"email"`
	Password string `gorm:"type:varchar(256)" json:"password"`
	Status   int    `gorm:"type:int" json:"status"`
	Role     int    `gorm:"type:int" json:"role"`
}

