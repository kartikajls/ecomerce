package cli

import (
	"e-comerce/internal/domain"
	"fmt"
	"strconv"
)

func (c *App) showCatalog() {
	products, _ := c.prodUC.GetCatalog()
	fmt.Println("\n--- KATALOG PAKAIAN TERSEDIA ---")
	for _, p := range products {
		fmt.Printf("[%d] %s | Ukuran: %s | Warna: %s | Rp%.2f | Stok: %d\n", p.ID, p.Name, p.Size, p.Color, p.Price, p.Stock)
	}
}

func (c *App) handleAddProduct() {
	fmt.Println("\n--- TAMBAH PRODUK ---")
	fmt.Print("ID Kategori (1: Atasan Pria, 2: Outerwear, 3: Celana): ")
	c.scanner.Scan()
	catID, _ := strconv.Atoi(c.scanner.Text())
	fmt.Print("SKU Unik: ")
	c.scanner.Scan()
	sku := c.scanner.Text()
	fmt.Print("Nama Baju: ")
	c.scanner.Scan()
	name := c.scanner.Text()
	fmt.Print("Ukuran (S/M/L/XL): ")
	c.scanner.Scan()
	size := c.scanner.Text()
	fmt.Print("Warna: ")
	c.scanner.Scan()
	color := c.scanner.Text()
	fmt.Print("Harga: ")
	c.scanner.Scan()
	price, _ := strconv.ParseFloat(c.scanner.Text(), 64)
	fmt.Print("Stok Awal: ")
	c.scanner.Scan()
	stock, _ := strconv.Atoi(c.scanner.Text())

	err := c.prodUC.AddProduct(domain.Product{CategoryID: catID, SKU: sku, Name: name, Size: size, Color: color, Price: price, Stock: stock})
	if err != nil {
		fmt.Println("❌ Gagal:", err)
	} else {
		fmt.Println("✅ Baju baru berhasil ditambahkan!")
	}
}

func (c *App) handleUpdateProduct() {
	fmt.Print("\nID Produk yang akan diupdate: ")
	c.scanner.Scan()
	id, _ := strconv.Atoi(c.scanner.Text())
	fmt.Print("Stok Baru: ")
	c.scanner.Scan()
	stock, _ := strconv.Atoi(c.scanner.Text())
	fmt.Print("Harga Baru: ")
	c.scanner.Scan()
	price, _ := strconv.ParseFloat(c.scanner.Text(), 64)

	if err := c.prodUC.UpdateProduct(id, stock, price); err != nil {
		fmt.Println("❌ Gagal:", err)
	} else {
		fmt.Println("✅ Produk diupdate!")
	}
}

func (c *App) handleDeleteProduct() {
	fmt.Print("\nID Produk yang akan dihapus: ")
	c.scanner.Scan()
	id, _ := strconv.Atoi(c.scanner.Text())
	if err := c.prodUC.RemoveProduct(id); err != nil {
		fmt.Println("❌ Gagal:", err)
	} else {
		fmt.Println("✅ Produk dihapus!")
	}
}
