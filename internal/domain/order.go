package domain

import "time"

type Order struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	OrderDate   time.Time `db:"order_date"`
	Status      string    `db:"status"`
	TotalAmount float64   `db:"total_amount"`
	CreatedAt   time.Time `db:"created_at"`
}

type OrderDetail struct {
	ID         int     `db:"id"`
	OrderID    int     `db:"order_id"`
	ProductID  int     `db:"product_id"`
	Quantity   int     `db:"quantity"`
	PriceAtBuy float64 `db:"price_at_buy"`
	Subtotal   float64 `db:"subtotal"`
}

type Payment struct {
	ID            int       `db:"id"`
	OrderID       int       `db:"order_id"`
	PaymentMethod string    `db:"payment_method"`
	PaymentStatus string    `db:"payment_status"`
	Amount        float64   `db:"amount"`
	PaidAt        time.Time `db:"paid_at"`
}

type OrderRepository interface {
	CreateCheckoutTx(userID int, items []OrderDetail, paymentMethod string) error
}

type OrderUseCase interface {
	Checkout(userID int, items []OrderDetail, paymentMethod string) error
}
