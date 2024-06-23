package routes

import (
	"github.com/BerkatPS/internal/domain/accounts"
	"github.com/BerkatPS/internal/domain/monthly_reports"
	"github.com/BerkatPS/internal/domain/transactions"
	"github.com/BerkatPS/internal/domain/users"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *users.UserController, accountController *accounts.AccountController, transactionController *transactions.TransactionController, monthlyreportcontroller *monthly_reports.MonthlyReportController) {
	api := app.Group("/api")

	// User routes
	api.Post("/users", userController.CreateUser)
	api.Get("/users", userController.GetAllUsers)
	api.Post("/users/login", userController.Login)

	// Account routes
	api.Post("/accounts", accountController.CreateAccount)
	api.Get("/accounts", accountController.GetAllAccounts)

	// Monthly Report routes
	api.Get("/monthly-reports", transactionController.GetAllTransactions)
	api.Get("/monthly-reports/:month/:year", monthlyreportcontroller.GetMonthlyReportsByMonthAndYear)
	api.Put("/monthly-reports/update", monthlyreportcontroller.UpdateMonthlyReport)

	// Transaction routes
	api.Post("/transactions", transactionController.CreateTransaction)
	api.Get("/transactions", transactionController.GetAllTransactions)
	api.Get("/daily-expense", transactionController.CountTotalExpenseDaily)
}
