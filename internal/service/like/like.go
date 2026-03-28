package like

import (
	"blog_r/internal/repository/like"
	"blog_r/internal/request"
	"errors"
)

type LikeService struct {
}

func NewLikeService() *LikeService {
	return &LikeService{}
}

func (s *LikeService) ToggleLike(req *request.SetORUndoLikeReq) error {
	if err := req.Validate(); err != nil {
		return err
	}

	if like.CanLike(req.UserID, req.TargetID, req.LikeType) {
		if err := like.SetLike(req.UserID, req.TargetID, req.LikeType); err != nil {
			return errors.New("like error")
		}
	} else {
		if err := like.UndoLike(req.UserID, req.TargetID, req.LikeType); err != nil {
			return errors.New("undo like error")
		}
	}
	return nil
}
