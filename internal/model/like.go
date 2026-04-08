package model

type Like struct {
	Model
	ArticleID int `json:"article_id" gorm:"uniqueIndex:idx_user_article_type"`
	UserID    int `json:"user_id" gorm:"uniqueIndex:idx_user_article_type;uniqueIndex:idx_user_comment_type"`
	CommentID int `json:"comment_id" gorm:"uniqueIndex:idx_user_comment_type"`
	Type      int `json:"type" gorm:"uniqueIndex:idx_user_article_type;uniqueIndex:idx_user_comment_type"`
}
