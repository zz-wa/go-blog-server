package user

import (
	"blog_r/internal/model"
	"blog_r/internal/request"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user *model.User) error {
	return m.Called(user).Error(0)
}
func (m *MockUserRepo) GetByUsername(name string) (model.User, error) {
	args := m.Called(name)
	return args.Get(0).(model.User), args.Error(1)
}
func (m *MockUserRepo) GetByEmail(email string) (model.User, error) {
	args := m.Called(email)
	return args.Get(0).(model.User), args.Error(1)
}
func (m *MockUserRepo) GetByID(id int) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}
func (m *MockUserRepo) UpdateUser(id int, req *model.User) error {
	return m.Called(id, req).Error(0)
}
func (m *MockUserRepo) ResetPassword(id int, hashedPassword string) error {
	return m.Called(id, hashedPassword).Error(0)
}
func (m *MockUserRepo) UpdateUserStatus(id int, status int) error {
	return m.Called(id, status).Error(0)
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name    string
		req     *request.RegisterReq
		wantErr string
	}{
		{
			name:    "用户名为空",
			req:     &request.RegisterReq{Username: "", Email: "a@b.com", Password: "123456"},
			wantErr: "参数不能为空",
		},
		{
			name:    "邮箱格式错误",
			req:     &request.RegisterReq{Username: "abc", Email: "notanemail", Password: "123456"},
			wantErr: "邮箱格式不正确",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewUserService(nil)
			err := svc.Register(tt.req)
			assert.EqualError(t, err, tt.wantErr)
		})
	}
}

func TestRegister_UsernameExists(t *testing.T) {
	mockRepo := &MockUserRepo{}
	svc := NewUserService(mockRepo)

	mockRepo.On("GetByUsername", "existuser").Return(model.User{Model: model.Model{ID: 1}}, nil)

	err := svc.Register(&request.RegisterReq{
		Username: "existuser",
		Email:    "a@b.com",
		Password: "123456",
	})
	assert.EqualError(t, err, "username already exists")
}

func TestRegister_Success(t *testing.T) {
	mockRepo := &MockUserRepo{}
	svc := NewUserService(mockRepo)

	mockRepo.On("GetByUsername", "newuser").Return(model.User{}, gorm.ErrRecordNotFound)
	mockRepo.On("GetByEmail", "new@b.com").Return(model.User{}, gorm.ErrRecordNotFound)
	mockRepo.On("CreateUser", mock.Anything).Return(nil)

	err := svc.Register(&request.RegisterReq{
		Username: "newuser",
		Email:    "new@b.com",
		Password: "123456",
	})
	assert.NoError(t, err)
}
