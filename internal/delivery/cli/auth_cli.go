package cli

import (
	"e-comerce/internal/domain"
	"fmt"
	"strings"
)

func (c *App) loginScreen() {
	fmt.Println("\n[ Wajib Login Untuk Melanjutkan ]")
	fmt.Print("Email: ")
	c.scanner.Scan()
	email := strings.TrimSpace(c.scanner.Text())

	fmt.Print("Password: ")
	c.scanner.Scan()
	pass := strings.TrimSpace(c.scanner.Text())

	user, err := c.authUC.Login(email, pass)
	if err != nil {
		fmt.Println("❌ Gagal Login:", err.Error())
		return
	}
	c.user = user
	fmt.Printf("✅ Berhasil Login! Selamat datang, %s (%s)\n", user.Email, user.Role)
}

func (c *App) registerScreen() {
	fmt.Println("\n=== DAFTAR AKUN BARU ===")

	fmt.Print("Email: ")
	c.scanner.Scan()
	email := strings.TrimSpace(c.scanner.Text())

	fmt.Print("Password: ")
	c.scanner.Scan()
	password := strings.TrimSpace(c.scanner.Text())

	fmt.Print("Nama Lengkap: ")
	c.scanner.Scan()
	fullName := strings.TrimSpace(c.scanner.Text())

	fmt.Print("No. HP: ")
	c.scanner.Scan()
	phone := strings.TrimSpace(c.scanner.Text())

	fmt.Print("Alamat: ")
	c.scanner.Scan()
	address := strings.TrimSpace(c.scanner.Text())

	req := domain.CreateUserRequest{
		Email:    email,
		Password: password,
		Role:     "customer",
		FullName: fullName,
		Phone:    phone,
		Address:  address,
	}

	// Memanggil fungsi registrasi
	err := c.authUC.Register(req)
	if err != nil {
		fmt.Println("❌ Gagal Mendaftar:", err.Error())
		return
	}

	fmt.Println("✅ Registrasi Berhasil! Data sudah masuk ke MySQL.")
	fmt.Println("👉 Silakan Login menggunakan akun yang baru kamu daftarkan.")
}
