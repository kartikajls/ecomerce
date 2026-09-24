package cli

import (
	"e-comerce/internal/domain"
	"fmt"
	"strconv"
	"strings"
)

func (c *App) handleCheckout() {
	fmt.Println("\n--- PROSES CHECKOUT ---")
	var items []domain.OrderDetail
	for {

		fmt.Print("Masukkan ID Produk (Ketik '0' jika selesai memilih): ")

		c.scanner.Scan()

		id, _ := strconv.Atoi(strings.TrimSpace(c.scanner.Text()))

		if id == 0 {

			break

		}
		fmt.Print("Jumlah Kuantitas: ")

		c.scanner.Scan()

		qty, _ := strconv.Atoi(strings.TrimSpace(c.scanner.Text()))

		items = append(items, domain.OrderDetail{ProductID: id, Quantity: qty})

	}

	if len(items) > 0 {
		// --- TAMBAHKAN PEMILIHAN METODE PEMBAYARAN DI SINI ---
		fmt.Println("\nPilih Metode Pembayaran:")
		fmt.Println("1. Transfer Bank")
		fmt.Println("2. QRIS")
		fmt.Println("3. COD")
		fmt.Print("Pilih [1-3]: ")
		c.scanner.Scan()
		pilihan := strings.TrimSpace(c.scanner.Text())

		metodeMap := map[string]string{
			"1": "transfer_bank",
			"2": "qris",
			"3": "cod",
		}

		metode, ok := metodeMap[pilihan]
		if !ok {
			fmt.Println("❌ Metode pembayaran tidak valid, default ke 'transfer_bank'")
			metode = "transfer_bank"
		}

		// Panggil Checkout dengan metode yang sudah dipilih user
		err := c.orderUC.Checkout(c.user.ID, items, metode)
		if err != nil {
			fmt.Println("❌ Checkout Gagal:", err.Error())
		} else {
			fmt.Println("🎉 Berhasil! Pesanan dan pemotongan stok sukses dilakukan.")
		}
	}
}
