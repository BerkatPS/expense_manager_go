package main

import (
	"context"
	"fmt"
	"github.com/BerkatPS/database"
	"github.com/BerkatPS/internal/domain/accounts"
	"github.com/BerkatPS/internal/domain/monthly_reports"
	"github.com/BerkatPS/internal/domain/transactions"
	"github.com/BerkatPS/internal/domain/users"
	"github.com/BerkatPS/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {

	fmt.Println(runtime.GOMAXPROCS(1))
	err := godotenv.Load()

	if err != nil {
		return
	}
	db, err := database.Connect()
	if err != nil {
		panic("failed to connect database")
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Expense Tracking Apps",
	})

	userRepository := users.NewUserRepository(db)
	accountRepository := accounts.NewAccountRepository(db)
	transactionRepository := transactions.NewTransactionRepository(db)
	monthlyReportRepository := monthly_reports.NewMonthlyReportRepository(db)

	userService := users.NewUserService(userRepository)
	accountService := accounts.NewAccountService(accountRepository)
	transactionService := transactions.NewTransactionService(transactionRepository)
	monthlyReportService := monthly_reports.NewMonthlyReportService(monthlyReportRepository)

	userController := users.NewUserController(userService)
	accountController := accounts.NewAccountController(accountService)
	transactionController := transactions.NewTransactionController(transactionService)
	monthlyreportcontroller := monthly_reports.NewMonthlyReportController(monthlyReportService)

	routes.SetupRoutes(app, userController, accountController, transactionController, monthlyreportcontroller)
	// create channel to listen for interupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":4000"); err != nil {
			log.Fatalf("Error Starting Server: %v", err)
		}
	}()

	// wait for interupt signal
	<-stop

	// create deadline for shutting down the server
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("server stopped gracefully")
}
