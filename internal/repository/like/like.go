package like

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	"context"
	"errors"
	"fmt"
)

func CanLike(UserID, TargetID, LikeType int) bool {
	var key string
	ctx := context.Background()

	if LikeType == 0 {
		key = fmt.Sprintf("like:article:%d", TargetID)
	} else {
		key = fmt.Sprintf("like:comment:%d", TargetID)
	}
	result, err := global.Redis.SIsMember(ctx, key, UserID).Result()
	if err != nil {
		return false
	}
	if result {
		return false
	}
	db := global.DB.Model(&model.Like{}).Where("user_id=?", UserID).Where("type=?", LikeType)
	if LikeType == 0 {
		db = db.Where("article_id=?", TargetID)
	} else {
		db = db.Where("comment_id=?", TargetID)
	}
	var count int64
	db.Count(&count)
	return count == 0

}

func SetLike(UserID, TargetID, LikeType int) error {
	if CanLike(UserID, TargetID, LikeType) {
		var Like model.Like
		var key string
		ctx := context.Background()
		Like.UserID = UserID
		Like.Type = LikeType
		if Like.Type == 0 {
			Like.ArticleID = TargetID
			key = fmt.Sprintf("like:article:%d", TargetID)
		} else {
			Like.CommentID = TargetID
			key = fmt.Sprintf("like:comment:%d", TargetID)

		}

		global.Redis.SAdd(ctx, key, UserID)
		return global.DB.Create(&Like).Error

	}
	return errors.New("you a already  like")
}

func UndoLike(UserID, TargetID, LikeType int) error {
	if !CanLike(UserID, TargetID, LikeType) {
		var key string
		ctx := context.Background()
		db := global.DB.Where("user_id=?", UserID).Where("type=?", LikeType)
		if LikeType == 0 {
			db = db.Where("article_id=?", TargetID)
			key = fmt.Sprintf("like:article:%d", TargetID)
		} else {
			db = db.Where("comment_id=?", TargetID)
			key = fmt.Sprintf("like:comment:%d", TargetID)
		}
		global.Redis.SRem(ctx, key, UserID)
		return db.Delete(&model.Like{}).Error
	}
	return errors.New("undo like error")
}
