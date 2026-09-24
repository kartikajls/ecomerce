package domain

import "time"

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Product struct {
	ID          int       `db:"id"`
	CategoryID  int       `db:"category_id"`
	SKU         string    `db:"sku"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Size        string    `db:"size"`
	Color       string    `db:"color"`
	Price       float64   `db:"price"`
	Stock       int       `db:"stock"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	Create(p Product) error
	Update(id int, stock int, price float64) error
	Delete(id int) error
}

type ProductUseCase interface {
	GetCatalog() ([]Product, error)
	AddProduct(p Product) error
	UpdateProduct(id int, stock int, price float64) error
	RemoveProduct(id int) error
}
