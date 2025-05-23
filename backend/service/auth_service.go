package service

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"backend/validation"

	"github.com/gofiber/fiber/v2"
)

func HandleRegister(c *fiber.Ctx) error {
	var body validation.RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request format")
	}

	if err := validation.ValidateRegister(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	hashed := utils.HashPassword(body.Password)
	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashed,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Email or Username already in use")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func HandleLogin(c *fiber.Ctx) error {
	var body validation.LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	if err := validation.ValidateLogin(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !utils.CheckPassword(body.Password, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Incorrect password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
