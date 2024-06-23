package transactions

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type TransactionController struct {
	Service TransactionService
}

func NewTransactionController(service TransactionService) *TransactionController {
	return &TransactionController{Service: service}
}

func (controller *TransactionController) CreateTransaction(c *fiber.Ctx) error {
	accountID, err := strconv.Atoi(c.FormValue("accountID"))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	categoryID, err := strconv.Atoi(c.FormValue("categoryID"))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	transactionType, err := strconv.Atoi(c.FormValue("transactionType"))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	amount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		return c.Status(400).JSON(err.Error())

	}

	err = controller.Service.CreateTransaction(accountID, categoryID, transactionType, amount)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON("Transaction created successfully")
}

func (controller *TransactionController) GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := controller.Service.GetAllTransactions()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(transactions)
}

func (controller *TransactionController) CountTotalExpenseDaily(c *fiber.Ctx) error {

	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid date format",
		})

	}
	totalExpense, err := controller.Service.CountTotalExpenseDaily(date)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"total_expense": totalExpense,
	})
}
