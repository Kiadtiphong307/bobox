package controller

import (
	"backend/service"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return service.HandleRegister(c)
}

func Login(c *fiber.Ctx) error {
	return service.HandleLogin(c)
}
