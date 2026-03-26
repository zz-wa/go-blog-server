package model

type Article struct {
	Model
	Title      string `json:"title" gorm:"type:varchar(256)"`
	Content    string `json:"content" gorm:"type:text"`
	Summary    string `json:"summary" gorm:"type:varchar(256)"`
	Cover      string `json:"cover" gorm:"type:varchar(256)"`
	Views      int    `json:"views" gorm:"type:int"`
	Status     int    `json:"status" gorm:"type:int"`
	CategoryID int    `json:"category_id" gorm:"type:int"`
	Tags       []Tag  `json:"tags" gorm:"many2many:article_tags;"`
}

type Category struct {
	Model
	Name string `json:"name" gorm:"type:varchar(256)"`
	Desc string `json:"desc" gorm:"type:varchar(256)"`
}

type Tag struct {
	Model
	Name  string `json:"name" gorm:"type:varchar(256)"`
	Color string `json:"color" gorm:"type:varchar(256)"`
}
