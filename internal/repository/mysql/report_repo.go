package mysql

import (
	"database/sql"
	"e-comerce/internal/domain"
)

type reportRepo struct{ db *sql.DB }

func NewReportRepository(db *sql.DB) domain.ReportRepository { return &reportRepo{db: db} }

func (r *reportRepo) GetUserReports() ([]domain.UserReportDTO, error) {
	rows, err := r.db.Query(`
	SELECT
		u.id, 
		u.email, 
		u.role, 
		COALESCE(p.full_name, '-'), COALESCE(p.phone, '-') 
	FROM users u 
	LEFT JOIN profiles p ON u.id = p.user_id 
	ORDER BY u.id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.UserReportDTO
	for rows.Next() {
		var d domain.UserReportDTO
		rows.Scan(&d.ID, &d.Email, &d.Role, &d.FullName, &d.Phone)
		list = append(list, d)
	}

	// cek error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *reportRepo) GetStockReports() ([]domain.StockReportDTO, error) {
	rows, err := r.db.Query(`
	SELECT 
		sku, 
		name, 
		size, 
		color, 
		stock, 
	CASE WHEN stock = 0 THEN 'HABIS' WHEN stock < 5 THEN 'MENIPIS' ELSE 'AMAN' END 
	FROM products ORDER BY stock ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.StockReportDTO
	for rows.Next() {
		var d domain.StockReportDTO
		rows.Scan(
			&d.SKU,
			&d.Name,
			&d.Size,
			&d.Color,
			&d.Stock,
			&d.StatusStok)
		list = append(list, d)
	}

	// cek error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *reportRepo) GetOrderReports() ([]domain.OrderReportDTO, error) {
	rows, err := r.db.Query(`
	SELECT 
		o.id, 
		u.email, 
		o.total_amount, 
		o.status, 
		p.payment_method,
	DATE_FORMAT(o.order_date, '%Y-%m-%d %H:%i') 
	FROM orders o 
	JOIN users u ON o.user_id = u.id 
	LEFT JOIN payments p ON p.order_id = o.id
	ORDER BY o.order_date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.OrderReportDTO
	for rows.Next() {
		var d domain.OrderReportDTO
		rows.Scan(
			&d.OrderID,
			&d.Pembeli,
			&d.TotalAmount,
			&d.Status,
			&d.PaymentMethod,
			&d.WaktuTransaksi)
		list = append(list, d)
	}

	// cek error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
