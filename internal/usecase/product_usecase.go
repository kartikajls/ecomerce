package usecase

import (
	"e-comerce/internal/domain"
	"errors"
)

type prodUC struct{ repo domain.ProductRepository }

func NewProductUseCase(r domain.ProductRepository) domain.ProductUseCase { return &prodUC{repo: r} }

func (u *prodUC) GetCatalog() ([]domain.Product, error) { return u.repo.GetAll() }

func (u *prodUC) AddProduct(p domain.Product) error {
	if p.Name == "" || p.Price <= 0 {
		return errors.New("nama dan harga produk tidak valid")
	}
	return u.repo.Create(p)
}

func (u *prodUC) UpdateProduct(id int, stock int, price float64) error {
	return u.repo.Update(id, stock, price)
}

func (u *prodUC) RemoveProduct(id int) error { return u.repo.Delete(id) }
