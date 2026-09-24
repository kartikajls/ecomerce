package cli

import (
	"fmt"
	"strings"
)

func (c *App) handleReports() {
	fmt.Println("\n=== PUSAT LAPORAN WAJIB (MANDATORY REPORTS) ===")
	fmt.Println("1. User Reports | 2. Stock Reports | 3. Order Reports")
	fmt.Print("Pilih Laporan [1-3]: ")
	c.scanner.Scan()
	switch strings.TrimSpace(c.scanner.Text()) {
	case "1":
		users, _ := c.reportUC.GetUsers()
		fmt.Println("\n--- USER REPORTS ---")
		for _, u := range users {
			fmt.Printf("ID: %d | Email: %s (%s) | Nama: %s | Telp: %s\n", u.ID, u.Email, u.Role, u.FullName, u.Phone)
		}
	case "2":
		stocks, _ := c.reportUC.GetStocks()
		fmt.Println("\n--- STOCK REPORTS ---")
		for _, s := range stocks {
			fmt.Printf("[%s] SKU: %s | %s (%s) | Sisa Stok: %d\n", s.StatusStok, s.SKU, s.Name, s.Size, s.Stock)
		}
	case "3":
		orders, _ := c.reportUC.GetOrders()
		fmt.Println("\n--- ORDER REPORTS ---")
		for _, o := range orders {
			// Tambahkan tampilan payment_method jika sudah ada di struct OrderReport
			fmt.Printf("Order #%d | Pembeli: %s | Total: Rp%.2f | Status: %s | Metode: %s | Waktu: %s\n",
				o.OrderID, o.Pembeli, o.TotalAmount, o.Status, o.PaymentMethod, o.WaktuTransaksi)
		}
	}
}
