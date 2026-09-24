package mysql

import (
	"database/sql"
	"e-comerce/internal/domain"
	"fmt"
)

type orderRepo struct{ db *sql.DB }

func NewOrderRepository(db *sql.DB) domain.OrderRepository { return &orderRepo{db: db} }

func (r *orderRepo) CreateCheckoutTx(userID int, items []domain.OrderDetail, paymentMethod string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var totalAmount float64 = 0

	res, err := tx.Exec(`INSERT INTO orders (user_id, status, total_amount) VALUES (?, 'pending', 0)`, userID)
	if err != nil {
		return fmt.Errorf("gagal membuat pesanan: %v", err)
	}
	orderID64, _ := res.LastInsertId()
	orderID := int(orderID64)

	for _, item := range items {
		var stock int
		var price float64
		var name string

		err := tx.QueryRow(`SELECT name, price, stock FROM products WHERE id = ? FOR UPDATE`, item.ProductID).
			Scan(&name, &price, &stock)
		if err != nil {
			return fmt.Errorf("produk ID %d tidak ditemukan", item.ProductID)
		}

		if stock < item.Quantity {
			return fmt.Errorf("stok '%s' tidak cukup (Sisa: %d)", name, stock)
		}

		item.PriceAtBuy = price
		item.Subtotal = price * float64(item.Quantity)
		totalAmount += item.Subtotal

		_, err = tx.Exec(`INSERT INTO order_details (order_id, product_id, quantity, price_at_buy, subtotal) VALUES (?, ?, ?, ?, ?)`,
			orderID, item.ProductID, item.Quantity, item.PriceAtBuy, item.Subtotal)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`UPDATE products SET stock = stock - ? WHERE id = ?`, item.Quantity, item.ProductID)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(`UPDATE orders SET total_amount = ? WHERE id = ?`, totalAmount, orderID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO payments (order_id, payment_method, payment_status, amount) VALUES (?, ?, 'unpaid', ?)`,
		orderID, paymentMethod, totalAmount)
	if err != nil {
		return err
	}

	return tx.Commit()
}
