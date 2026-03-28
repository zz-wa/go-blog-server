package model

type Comment struct {
	Model
	UserID    int    `json:"user_id" gorm:"type:int"`
	ArticleID int    `json:"article_id" gorm:"int"`
	Content   string `json:"content" gorm:"content"`
	ReplyID   int    `json:"reply_id" gorm:"type:int"`
}
