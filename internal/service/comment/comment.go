package comment

import (
	"blog_r/internal/model"
	"blog_r/internal/request"
	"errors"
)

type CommentRepo interface {
	CreateComment(comment *model.Comment) error
	GetCommentList(articleID, page, pageSize int) ([]model.Comment, int64, error)
	DeleteComment(id int) error
}

type CommentService struct {
	repo CommentRepo
}

func NewCommentService(repo CommentRepo) *CommentService {
	return &CommentService{repo: repo}
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

	if err := s.repo.CreateComment(SetComment); err != nil {
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
	CommentList, total, err := s.repo.GetCommentList(req.ArticleID, req.Page, req.PageSize)
	if err != nil {
		return nil, 0, errors.New("failde to get comment list")
	}
	return CommentList, total, nil
}

func (s *CommentService) DeleteComment(id int) error {
	if id <= 0 {
		return errors.New("invalid comment id")
	}
	if err := s.repo.DeleteComment(id); err != nil {
		return errors.New("fail to delete comment")
	}
	return nil
}
