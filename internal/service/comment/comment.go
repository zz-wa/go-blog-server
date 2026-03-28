package comment

import (
	"blog_r/internal/model"
	"blog_r/internal/repository/comment"
	"blog_r/internal/request"
	"errors"
)

type CommentService struct {
}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (s *CommentService) CreateComment(req *request.CreateCommentReq) error {
	if req == nil {
		return errors.New("invalid request")
	}
	if err := req.Validate(); err != nil {
		return err
	}
	SetComment := &model.Comment{}
	SetComment.UserID = req.UserID
	SetComment.ArticleID = req.ArticleID
	SetComment.Content = req.Content
	SetComment.ReplyID = req.ReplyID

	if err := comment.CreateComment(SetComment); err != nil {
		return errors.New("failed to create commment")
	}
	return nil
}

func (s *CommentService) GetCommentList(req *request.CommentListReq) ([]model.Comment, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}

	req.SetDefault()
	if err := req.Validate(); err != nil {
		return nil, 0, err
	}
	CommentList, total, err := comment.GetCommentList(req.ArticleID, req.Page, req.PageSize)
	if err != nil {
		return nil, 0, errors.New("failde to get comment list")
	}
	return CommentList, total, nil
}

func (s *CommentService) DeleteComment(id int) error {
	if id <= 0 {
		return errors.New("invalid comment id")
	}
	if err := comment.DeleteComment(id); err != nil {
		return errors.New("fail to delete comment")
	}
	return nil
}
