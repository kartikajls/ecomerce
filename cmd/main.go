package main

import (
	"e-comerce/config"
	"e-comerce/internal/delivery/cli"
	"e-comerce/internal/repository/mysql"
	"e-comerce/internal/usecase"
	"fmt"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("error: Gagal terhubung ke MySQL:", err)
		return
	}
	defer db.Close()

	fmt.Println("Database connected successfully")

	userRepo := mysql.NewUserRepository(db)
	prodRepo := mysql.NewProductRepository(db)
	orderRepo := mysql.NewOrderRepository(db)
	reportRepo := mysql.NewReportRepository(db)

	authUC := usecase.NewAuthUseCase(userRepo)
	userUC := usecase.NewUserUseCase(userRepo)
	prodUC := usecase.NewProductUseCase(prodRepo)
	orderUC := usecase.NewOrderUseCase(orderRepo)
	reportUC := usecase.NewReportUseCase(reportRepo)

	app := cli.NewApp(authUC, userUC, prodUC, orderUC, reportUC)
	app.Start()
}
