package usecase

import "e-comerce/internal/domain"

type reportUC struct{ repo domain.ReportRepository }

func NewReportUseCase(r domain.ReportRepository) domain.ReportUseCase { return &reportUC{repo: r} }
func (u *reportUC) GetUsers() ([]domain.UserReportDTO, error)         { return u.repo.GetUserReports() }
func (u *reportUC) GetStocks() ([]domain.StockReportDTO, error)       { return u.repo.GetStockReports() }
func (u *reportUC) GetOrders() ([]domain.OrderReportDTO, error)       { return u.repo.GetOrderReports() }
