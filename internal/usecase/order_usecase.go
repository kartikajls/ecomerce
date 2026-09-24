package usecase

import (
	"e-comerce/internal/domain"
	"errors"
)

type orderUC struct{ repo domain.OrderRepository }

func NewOrderUseCase(r domain.OrderRepository) domain.OrderUseCase { return &orderUC{repo: r} }

func (u *orderUC) Checkout(userID int, items []domain.OrderDetail, method string) error {
	if userID <= 0 {
		return errors.New("sesi tidak valid, silakan login ulang")
	}
	if len(items) == 0 {
		return errors.New("keranjang kosong")
	}
	if method == "" {
		method = "transfer_bank"
	}
	return u.repo.CreateCheckoutTx(userID, items, method)
}
