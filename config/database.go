package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // <--- Tanda _ ini WAJIB ada
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Gagal memuat .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		name,
	)

	// buka koneksi database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Tidak dapat terhubung dengan DB: %v", err)
	}

	// test koneksi
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Tidak bisa Ping DB: %v", err)
	}

	return db, nil

}
