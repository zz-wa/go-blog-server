package article

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	articleRepo "blog_r/internal/repository/article"
	"blog_r/internal/repository/tag"
	"blog_r/internal/request"
	"errors"
	"time"
)

type ArticleService struct {
}

// 显m示的具体文章的结构体
type ArchiveArticleItem struct {
	ID        int       `json:"id"`
	Title     string    ` json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

// 这个月的文章列表结构体
type ArchiveGroup struct {
	Date     string               `json:"date"`
	Articles []ArchiveArticleItem `json:"articles" `
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) CreateArticle(req *request.CreateArticleReq) error {
	if req == nil {
		return errors.New("invalid request")
	}
	if err := req.Validate(); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return err
	}
	SetArticle := &model.Article{}
	SetArticle.Title = req.Title
	SetArticle.Content = req.Content
	SetArticle.Summary = req.Summary
	SetArticle.Cover = req.Cover
	SetArticle.CategoryID = req.CategoryID
	SetArticle.Status = req.Status

	if len(req.Tags) > 0 {
		tags, err := tag.GetByIDs(req.Tags)
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

func (s *ArticleService) GetArticleList(req *request.ArticleListReq) ([]model.Article, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}
	req.SetDefault()

	ArticleList, total, err := articleRepo.GetArticleList(req.Page, req.PageSize, req.Status, req.CategoryID, req.TagID, req.Keyword)
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
func (s *ArticleService) GetArticleArchive() ([]ArchiveGroup, error) {
	articles, err := articleRepo.GetPublishedArticleForArchive()
	var AllTimes []string
	if err != nil {
		return nil, err
	}
	if articles == nil {
		return nil, err
	}
	m := make(map[string][]ArchiveArticleItem)
	for _, article := range articles {
		key := article.CreatedAt.Format("2006-01")
		_, ok := m[key]
		if !ok {
			AllTimes = append(AllTimes, key)
		}
		m[key] = append(m[key], ArchiveArticleItem{article.ID, article.Title, article.CreatedAt})
	}
	var result []ArchiveGroup

	for _, month := range AllTimes {
		result = append(result, ArchiveGroup{
			month,
			m[month],
		})
	}

	return result, nil
}
