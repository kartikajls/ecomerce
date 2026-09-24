package cli

import (
	"bufio"
	"e-comerce/internal/domain"
	"fmt"
	"os"
	"strings"
)

type App struct {
	authUC   domain.AuthUseCase
	userUC   domain.UserUseCase
	prodUC   domain.ProductUseCase
	orderUC  domain.OrderUseCase
	reportUC domain.ReportUseCase
	scanner  *bufio.Scanner
	user     *domain.User
}

func NewApp(a domain.AuthUseCase, u domain.UserUseCase, p domain.ProductUseCase, o domain.OrderUseCase, r domain.ReportUseCase) *App {
	return &App{authUC: a, userUC: u, prodUC: p, orderUC: o, reportUC: r, scanner: bufio.NewScanner(os.Stdin)}
}

func (c *App) Start() { // Atau Run() sesuai kode kamu
	for c.user == nil {
		fmt.Println("\n=== CLOTHING E-COMMERCE ===")
		fmt.Println("1. Login")
		fmt.Println("2. Daftar Akun Baru (Register)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		c.scanner.Scan()
		pilihan := strings.TrimSpace(c.scanner.Text())

		switch pilihan {
		case "1":
			c.loginScreen() // Kalau sukses, c.user tidak nil lagi dan loop berhenti
		case "2":
			c.registerScreen() // Daftar ke database, pas selesai balik lagi ke menu ini buat login
		case "0":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid.")
		}
	}

	// --- MASUK KE MENU UTAMA (Katalog, Checkout, Laporan) ---
	c.mainMenu()
}

func (c *App) mainMenu() {
	for {
		fmt.Println("\n--- MENU UTAMA ---")
		fmt.Println("1. Lihat Katalog Pakaian")
		fmt.Println("2. Checkout / Beli Pakaian")
		if c.user.Role == "admin" {
			fmt.Println("3. [Admin] Tambah Pakaian Baru")
			fmt.Println("4. [Admin] Update Stok & Harga")
			fmt.Println("5. [Admin] Hapus Pakaian")
		}
		fmt.Println("6. [Mandatory] Laporan Sistem (Users, Stock, Orders)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		c.scanner.Scan()
		choice := strings.TrimSpace(c.scanner.Text())

		switch choice {
		case "1":
			c.showCatalog()
		case "2":
			c.handleCheckout()
		case "3":
			if c.user.Role == "admin" {
				c.handleAddProduct()
			} else {
				fmt.Println("⚠️ Khusus Admin.")
			}
		case "4":
			if c.user.Role == "admin" {
				c.handleUpdateProduct()
			} else {
				fmt.Println("⚠️ Khusus Admin.")
			}
		case "5":
			if c.user.Role == "admin" {
				c.handleDeleteProduct()
			} else {
				fmt.Println("⚠️ Khusus Admin.")
			}
		case "6":
			c.handleReports()
		case "0":
			fmt.Println("Terima kasih!")
			os.Exit(0)
		default:
			fmt.Println("⚠️ Pilihan tidak valid.")
		}
	}
}
