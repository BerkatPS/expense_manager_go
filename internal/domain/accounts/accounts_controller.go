package accounts

import (
	"github.com/BerkatPS/internal/domain/model"
	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	Service AccountService
}

func NewAccountController(service AccountService) *AccountController {
	return &AccountController{Service: service}
}

func (controller *AccountController) CreateAccount(c *fiber.Ctx) error {
	account := new(model.Accounts)
	if err := c.BodyParser(account); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	account, err := controller.Service.CreateAccount(account)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(account)
}

func (controller *AccountController) GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := controller.Service.GetAllAccounts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(accounts)
}
