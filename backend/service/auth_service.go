package service

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"backend/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

)

// คือฟังก์ชันที่จะจัดการกับการลงทะเบียนของผู้ใช้
func HandleRegister(c *fiber.Ctx) error {
	// ดึงข้อมูลจากร่างการส่งข้อมูลจากผู้ใช้
	var body validation.RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request format")
	}
	// ตรวจสอบข้อมูลที่ส่งมาว่ามีความถูกต้องหรือไม่
	if err := validation.ValidateRegister(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// ถ้าข้อมูลถูกต้องจะถูกลบรหัสผ่านออกจากข้อมูลที่ส่งมา
	hashed := utils.HashPassword(body.Password)
	// สร้างข้อมูลผู้ใช้ใหม่
	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashed,
	}
	// สร้างข้อมูลผู้ใช้ใหม่ในฐานข้อมูล
	if err := database.DB.Create(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Email or Username already in use")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}

// คือฟังก์ชันที่จะจัดการกับการลงชื่อเข้าใช้งานของผู้ใช้
func HandleLogin(c *fiber.Ctx) error {
	// ดึงข้อมูลจากร่างการส่งข้อมูลจากผู้ใช้
	var body validation.LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
	if err := validation.ValidateLogin(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// ค้นหาข้อมูลผู้ใช้ในฐานข้อมูล
	var user models.User
	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}
	// ตรวจสอบรหัสผ่านว่ามีความถูกต้องหรือไม่ โดยใช้ฟังก์ชัน CheckPassword จาก utils
	if !utils.CheckPassword(body.Password, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Incorrect password")
	}
	// สร้างข้อมูลผู้ใช้ใหม่ในฐานข้อมูล
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token")
	}
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}

// คือฟังก์ชันที่จะจัดการกับการดึงข้อมูลผู้ใช้
func HandleProfile(c *fiber.Ctx) error {

	// ดึง token จากการส่งมาจากผู้ใช้
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
	}

	// คือข้อมูลที่จะดึงออกมาจาก token
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
	}

	// แปลง user_id เป็น float64 เพื่อใช้ในการค้นหาข้อมูลผู้ใช้
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID in token")
	}

	userID := uint(userIDFloat)

	// ค้นหาข้อมูลผู้ใช้ในฐานข้อมูล
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(fiber.Map{
		"message": "Profile retrieved successfully",
		"user":    user,
	})
}

