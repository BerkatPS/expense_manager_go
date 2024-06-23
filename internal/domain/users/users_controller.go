package users

import (
	"github.com/BerkatPS/internal/domain/model"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{Service: service}
}

func (controller *UserController) CreateUser(c *fiber.Ctx) error {
	user := new(model.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	user, err := controller.Service.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(user)
}

func (controller *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := controller.Service.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(users)
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	user := new(model.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(err.Error())
	}
	user, err := controller.Service.Login(user.Username, user.Password)
	if err != nil {
		return c.Status(401).JSON(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "login success",
	})
}
