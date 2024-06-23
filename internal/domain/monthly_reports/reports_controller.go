package monthly_reports

import (
	"github.com/BerkatPS/internal/domain/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type MonthlyReportController struct {
	Service MonthlyReportService
}

func NewMonthlyReportController(service MonthlyReportService) *MonthlyReportController {
	return &MonthlyReportController{Service: service}
}

func (controller *MonthlyReportController) CreateMonthlyReport(c *fiber.Ctx) error {
	report := new(model.MonthlyReport)
	if err := c.BodyParser(report); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	report, err := controller.Service.CreateMonthlyReport(report)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(report)
}

func (controller *MonthlyReportController) GetAllMonthlyReports(c *fiber.Ctx) error {
	reports, err := controller.Service.GetAllMonthlyReports()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(reports)
}

func (controller *MonthlyReportController) GetMonthlyReportsByMonthAndYear(c *fiber.Ctx) error {
	month, err := strconv.Atoi(c.Params("month"))
	if err != nil {
		return c.Status(400).SendString("Invalid month")
	}

	year, err := strconv.Atoi(c.Params("year"))
	if err != nil {
		return c.Status(400).SendString("Invalid year")
	}

	reports, err := controller.Service.GetMonthlyReportsByMonthAndYear(month, year)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(reports)
}

func (controller *MonthlyReportController) UpdateMonthlyReport(c *fiber.Ctx) error {
	var data struct {
		AccountID       int     `json:"account_id"`
		Amount          float64 `json:"amount"`
		TransactionType int     `json:"transaction_type"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	err := controller.Service.UpdateMonthlyReport(data.AccountID, data.Amount, data.TransactionType)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}
