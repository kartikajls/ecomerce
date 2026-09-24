package domain

type UserReportDTO struct {
	ID       int    `db:"user_id"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	FullName string `db:"full_name"`
	Phone    string `db:"phone"`
}

type StockReportDTO struct {
	SKU        string `db:"sku"`
	Name       string `db:"name"`
	Size       string `db:"size"`
	Color      string `db:"color"`
	Stock      int    `db:"stock"`
	StatusStok string `db:"status_stok"`
}

type OrderReportDTO struct {
	OrderID        int     `db:"order_id"`
	Pembeli        string  `db:"pembeli"`
	TotalAmount    float64 `db:"total_amount"`
	Status         string  `db:"status"`
	PaymentMethod  string  `db:"payment_method"`
	WaktuTransaksi string  `db:"waktu_transaksi"`
}

type ReportRepository interface {
	GetUserReports() ([]UserReportDTO, error)
	GetStockReports() ([]StockReportDTO, error)
	GetOrderReports() ([]OrderReportDTO, error)
}

type ReportUseCase interface {
	GetUsers() ([]UserReportDTO, error)
	GetStocks() ([]StockReportDTO, error)
	GetOrders() ([]OrderReportDTO, error)
}
