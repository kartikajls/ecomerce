package usecase

import (
	"e-comerce/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

// ---> PASTIKAN FUNGSI CREATE INI ADA DAN TIDAK TERHAPUS <---
func (m *MockUserRepository) Create(req domain.CreateUserRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockUserRepository) Login(email, password string) (*domain.User, error) {
	args := m.Called(email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetAll() ([]domain.UserReportDTO, error) {
	args := m.Called()
	return args.Get(0).([]domain.UserReportDTO), args.Error(1)
}

func (m *MockUserRepository) UpdateRole(userID int, newRole string) error {
	args := m.Called(userID, newRole)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}
