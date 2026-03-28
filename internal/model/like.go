package model

type Like struct {
	Model
	UserID    int `json:"user_id"`
	ArticleID int `json:"article_id"`
	CommentID int `json:"comment_id"`
	Type      int `json:"type"`
}
