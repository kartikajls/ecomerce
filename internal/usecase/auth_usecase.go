package usecase

import (
	"e-comerce/internal/domain"
	"errors"
	"strings"
)

type authUC struct{ repo domain.UserRepository }

func NewAuthUseCase(r domain.UserRepository) domain.AuthUseCase { return &authUC{repo: r} }

func (u *authUC) Login(e, p string) (*domain.User, error) {
	if e == "" || p == "" {
		return nil, errors.New("email dan password wajib diisi")
	}
	return u.repo.Login(e, p)
}

type userUC struct{ repo domain.UserRepository }

func NewUserUseCase(r domain.UserRepository) domain.UserUseCase { return &userUC{repo: r} }

func (u *userUC) GetAccountList() ([]domain.UserReportDTO, error) {
	return u.repo.GetAll()
}

func (u *userUC) RegisterNewAccount(req domain.CreateUserRequest) error {
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" || req.Password == "" || req.FullName == "" {
		return errors.New("email, password, dan nama lengkap wajib diisi")
	}
	if req.Role != "admin" && req.Role != "customer" {
		req.Role = "customer"
	}
	return u.repo.Create(req)
}

func (u *userUC) ChangeUserRole(userID int, newRole string) error {
	if userID <= 0 {
		return errors.New("ID user tidak valid")
	}
	if newRole != "admin" && newRole != "customer" {
		return errors.New("role hanya boleh 'admin' atau 'customer'")
	}
	return u.repo.UpdateRole(userID, newRole)
}

func (u *userUC) RemoveAccount(userID int) error {
	if userID <= 0 {
		return errors.New("ID user tidak valid")
	}
	return u.repo.Delete(userID)
}

func (u *authUC) Register(req domain.CreateUserRequest) error {
	// 1. Validasi sederhana agar input tidak kosong
	if req.Email == "" || req.Password == "" {
		return errors.New("email dan password tidak boleh kosong")
	}

	// 2. Panggil fungsi Create dari user repository yang sudah kamu buat sebelumnya
	return u.repo.Create(req)
}
