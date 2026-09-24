package mysql

import (
	"database/sql"
	"e-comerce/internal/domain"
	"errors"
	"fmt"
)

type userRepo struct{ db *sql.DB }

func NewUserRepository(db *sql.DB) domain.UserRepository { return &userRepo{db: db} }

func (r *userRepo) Login(email, password string) (*domain.User, error) {
	var u domain.User
	err := r.db.QueryRow(`SELECT id, email, role FROM users WHERE email = ? AND password = ?`, email, password).
		Scan(&u.ID, &u.Email, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("email atau password salah")
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) GetAll() ([]domain.UserReportDTO, error) {
	query := `SELECT u.id, u.email, u.role, COALESCE(p.full_name, '-'), COALESCE(p.phone, '-') 
	          FROM users u LEFT JOIN profiles p ON u.id = p.user_id ORDER BY u.id ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.UserReportDTO
	for rows.Next() {
		var u domain.UserReportDTO
		rows.Scan(&u.ID, &u.Email, &u.Role, &u.FullName, &u.Phone)
		list = append(list, u)
	}

	// cek error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *userRepo) Create(req domain.CreateUserRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.Exec(`INSERT INTO users (email, password, role) VALUES (?, ?, ?)`, req.Email, req.Password, req.Role)
	if err != nil {
		return fmt.Errorf("email sudah terdaftar atau error: %v", err)
	}

	newID64, _ := res.LastInsertId()
	newID := int(newID64)

	_, err = tx.Exec(`INSERT INTO profiles (user_id, full_name, phone, address) VALUES (?, ?, ?, ?)`,
		newID, req.FullName, req.Phone, req.Address)
	if err != nil {
		return fmt.Errorf("gagal membuat profil: %v", err)
	}

	return tx.Commit()
}

func (r *userRepo) UpdateRole(userID int, newRole string) error {
	res, err := r.db.Exec(`UPDATE users SET role = ? WHERE id = ?`, newRole, userID)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("user ID tidak ditemukan")
	}
	return nil
}

func (r *userRepo) Delete(userID int) error {
	res, err := r.db.Exec(`DELETE FROM users WHERE id = ?`, userID)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("user ID tidak ditemukan")
	}
	return nil
}
