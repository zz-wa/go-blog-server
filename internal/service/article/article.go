package article

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	articleRepo "blog_r/internal/repository/article"
	"blog_r/internal/repository/tag"
	"blog_r/internal/request"
	"errors"
)

type ArticleService struct {
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) CreateArticle(Req *request.CreateArticleReq) error {
	if Req == nil {
		return errors.New("invalid request")
	}
	if err := Req.Validate(); err != nil {
		return err
	}

	if Req.Title == "" || Req.Content == "" {
		return errors.New("invalid request")
	}
	SetArticle := &model.Article{}
	SetArticle.Title = Req.Title
	SetArticle.Content = Req.Content
	SetArticle.Summary = Req.Summary
	SetArticle.Cover = Req.Cover
	SetArticle.CategoryID = Req.CategoryID
	SetArticle.Status = Req.Status

	if len(Req.Tags) > 0 {
		tags, err := tag.GetByIDs(Req.Tags)
		if err != nil {
			return errors.New("fail to load tags")
		}
		SetArticle.Tags = tags
	}
	if err := articleRepo.CreateArticle(SetArticle); err != nil {
		return errors.New("failed to create article")
	}
	return nil
}

func (s *ArticleService) GetArticleByID(id int) (model.Article, error) {
	if id <= 0 {
		return model.Article{}, errors.New("invalid article ID")
	}
	var Article model.Article
	Article, err := articleRepo.GetArticleByID(id)
	if err != nil {
		return model.Article{}, errors.New("article not found")
	}
	return Article, nil
}

func (s *ArticleService) GetArticleList(Req *request.ArticleListReq) ([]model.Article, int64, error) {
	if Req == nil {
		return nil, 0, errors.New("invalid request")
	}
	Req.SetDefault()

	ArticleList, total, err := articleRepo.GetArticleList(Req.Page, Req.PageSize, Req.Status, Req.CategoryID, Req.TagID, Req.Keyword)
	if err != nil {
		return nil, 0, errors.New("failed to get article list")
	}
	return ArticleList, total, nil
}

func (s *ArticleService) UpdateArticle(id int, NewArticle *model.Article, tagIDs []int) error {
	if id <= 0 || NewArticle == nil {
		return errors.New("invalid request")
	}
	NewArticle.ID = id
	if err := articleRepo.UpdateArticle(NewArticle); err != nil {
		return errors.New("failed to update article")
	}
	var newTags []model.Tag
	if len(tagIDs) > 0 {
		tags, err := tag.GetByIDs(tagIDs)
		if err != nil {
			return errors.New("fail to load tags")
		}
		newTags = tags
	}
	if err := global.DB.Model(NewArticle).Association("Tags").Replace(newTags); err != nil {
		return errors.New("failed to update article tags")
	}
	return nil
}

func (s *ArticleService) DeleteArticle(id int) error {
	if id <= 0 {
		return errors.New("invalid article ID")
	}
	err := articleRepo.DeleteArticle(id)
	if err != nil {
		return errors.New("failed to delete article")
	}
	return nil
}
