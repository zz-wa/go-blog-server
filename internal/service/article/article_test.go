package article

import (
	"blog_r/internal/model"
	"blog_r/internal/request"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockArticleRepo struct {
	mock.Mock
}

func (m *MockArticleRepo) ReplaceArticleTag(NewArticle *model.Article, newTags []model.Tag) error {
	panic("implement me")
}

func (m *MockArticleRepo) CreateArticle(article *model.Article) error {
	args := m.Called(article)
	return args.Error(0)
}
func (m *MockArticleRepo) GetArticleByID(id int) (model.Article, error) {
	args := m.Called(id)
	return model.Article{}, args.Error(0)
}
func (m *MockArticleRepo) GetArticleList(page, pageSize int, status *int, categoryID, tagID int, keyWord string) ([]model.Article, int64, error) {
	args := m.Called(page, pageSize, status, categoryID, tagID, keyWord)
	return []model.Article{}, 0, args.Error(0)
}
func (m *MockArticleRepo) GetPublishedArticleForArchive() ([]model.Article, error) {
	args := m.Called()
	return args.Get(0).([]model.Article), args.Error(1)
}
func (m *MockArticleRepo) UpdateArticle(article *model.Article) error {
	args := m.Called(article)
	return args.Error(0)
}
func (m *MockArticleRepo) DeleteArticle(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
func TestCreateArticle(t *testing.T) {
	svc := NewArticleService(nil)
	tests := []struct {
		name    string
		req     *request.CreateArticleReq
		wantErr string
	}{
		{
			name:    "req为nil",
			req:     nil,
			wantErr: "invalid request",
		},
		{
			name: "标题为空",
			req: &request.CreateArticleReq{
				Title:   "",
				Content: "some content",
			},
			wantErr: "标题和内容不能为空",
		},
		{
			name: "status非法",
			req: &request.CreateArticleReq{
				Title:   "标题",
				Content: "内容",
				Status:  99,
			},
			wantErr: "status 只能是 0/1/2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.CreateArticle(tt.req)
			assert.EqualError(t, err, tt.wantErr)
		})
	}

}

func TestCreatArticle_WithRepo(t *testing.T) {
	mockRepo := &MockArticleRepo{}
	svc := NewArticleService(mockRepo)

	mockRepo.On("CreateArticle", mock.Anything).Return(nil)
	req := &request.CreateArticleReq{
		Title:   "测试内容",
		Content: "测试内容",
		Status:  1,
	}
	err := svc.CreateArticle(req)
	assert.NoError(t, err)
}
func TestGetArticleArchive(t *testing.T) {
	tests := []struct {
		name      string
		articles  []model.Article
		repoErr   error
		wantErr   bool
		wantEmpty bool
	}{
		{
			name:      "repo报错",
			articles:  []model.Article{},
			repoErr:   errors.New("db error"),
			wantErr:   true,
			wantEmpty: true,
		},
		{
			name:      "空列表",
			articles:  []model.Article{},
			repoErr:   nil,
			wantErr:   false,
			wantEmpty: true,
		},
		{
			name: "正常分组",
			articles: []model.Article{
				{
					Title: "i",
					Model: model.Model{
						CreatedAt: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC),
					},
				},
				{
					Title: "o",
					Model: model.Model{
						CreatedAt: time.Date(2024, 1, 12, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockArticleRepo{}
			svc := NewArticleService(mockRepo)
			mockRepo.On("GetPublishedArticleForArchive").Return(tt.articles, tt.repoErr)
			result, err := svc.GetArticleArchive()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if tt.wantEmpty {
				assert.Empty(t, result)
			} else {
				assert.Len(t, result, 1)
				assert.Len(t, result[0].Articles, 2)
			}
		})
	}

}
