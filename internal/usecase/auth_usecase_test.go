package usecase

import (
	"e-comerce/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authUC := NewAuthUseCase(mockRepo)

	req := domain.CreateUserRequest{
		Email:    "test@mail.com",
		Password: "password123",
		Role:     "customer",
		FullName: "User Test",
	}

	mockRepo.On("Create", req).Return(nil)

	err := authUC.Register(req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegister_EmptyEmail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authUC := NewAuthUseCase(mockRepo)

	req := domain.CreateUserRequest{
		Email:    "",
		Password: "password123",
	}

	err := authUC.Register(req)

	assert.Error(t, err)
	assert.Equal(t, "email dan password tidak boleh kosong", err.Error())
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authUC := NewAuthUseCase(mockRepo)

	expectedUser := &domain.User{
		ID:    1,
		Email: "test@mail.com",
		Role:  "customer",
	}

	mockRepo.On("Login", "test@mail.com", "password123").Return(expectedUser, nil)

	user, err := authUC.Login("test@mail.com", "password123")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@mail.com", user.Email)
	mockRepo.AssertExpectations(t)
}
