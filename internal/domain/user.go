package domain

import "time"

type User struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
}

type Profile struct {
	ID       int    `db:"id"`
	UserID   int    `db:"user_id"`
	FullName string `db:"full_name"`
	Phone    string `db:"phone"`
	Address  string `db:"address"`
}

type CreateUserRequest struct {
	Email    string
	Password string
	Role     string
	FullName string
	Phone    string
	Address  string
}

type UserRepository interface {
	Login(email, password string) (*User, error)
	GetAll() ([]UserReportDTO, error)
	Create(req CreateUserRequest) error
	UpdateRole(userID int, newRole string) error
	Delete(userID int) error
}

type AuthUseCase interface {
	Login(email, password string) (*User, error)
	Register(req CreateUserRequest) error // <-- TAMBAHKAN BARIS INI
}

type UserUseCase interface {
	GetAccountList() ([]UserReportDTO, error)
	RegisterNewAccount(req CreateUserRequest) error
	ChangeUserRole(userID int, newRole string) error
	RemoveAccount(userID int) error
}
