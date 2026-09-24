package mysql

import (
	"database/sql"
	"e-comerce/internal/domain"
	"errors"
)

type productRepo struct{ db *sql.DB }

func NewProductRepository(db *sql.DB) domain.ProductRepository { return &productRepo{db: db} }

func (r *productRepo) GetAll() ([]domain.Product, error) {
	rows, err := r.db.Query(`
	SELECT 
			id, 
			category_id, 
			sku, 
			name, 
			size, 
			color, 
			price, 
			stock 
	FROM products 
	ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(
			&p.ID,
			&p.CategoryID,
			&p.SKU,
			&p.Name,
			&p.Size,
			&p.Color,
			&p.Price,
			&p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	// cek error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// untuk menambah pakaian baru yang belum ada
func (r *productRepo) Create(p domain.Product) error {
	_, err := r.db.Exec(`
	INSERT INTO products (category_id, sku, name, description, size, color, price, stock) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		p.CategoryID, p.SKU, p.Name, p.Description, p.Size, p.Color, p.Price, p.Stock)
	return err
}

// untuk mengupdate stock dan harga
func (r *productRepo) Update(id int, stock int, price float64) error {
	res, err := r.db.Exec(`UPDATE products SET stock = ?, price = ? WHERE id = ?`, stock, price, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("produk tidak ditemukan")
	}
	return nil
}

func (r *productRepo) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM products WHERE id = ?`, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("produk tidak ditemukan")
	}
	return nil
}
