package service

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"backend/validation"

	"github.com/gofiber/fiber/v2"
)

// คือฟังก์ชันที่จะจัดการกับการลงทะเบียนของผู้ใช้
func HandleRegister(c *fiber.Ctx) error {
	// ดึงข้อมูลจากร่างการส่งข้อมูลจากผู้ใช้
	var body validation.RegisterRequest
	// ถ้ามีข้อมูลที่ส่งมาผิดพลาดจะส่งข้อความผิดพลาดกลับไป
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
		// ถ้ามีข้อมูลที่ส่งมาผิดพลาดจะส่งข้อความผิดพลาดกลับไป
		return fiber.NewError(fiber.StatusInternalServerError, "Email or Username already in use")
	}
	// ส่งข้อความสำเร็จกลับไป
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// คือฟังก์ชันที่จะจัดการกับการลงชื่อเข้าใช้งานของผู้ใช้
func HandleLogin(c *fiber.Ctx) error {
	// ดึงข้อมูลจากร่างการส่งข้อมูลจากผู้ใช้
	var body validation.LoginRequest
	// ถ้ามีข้อมูลที่ส่งมาผิดพลาดจะส่งข้อความผิดพลาดกลับไป
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
	// ตรวจสอบข้อมูลที่ส่งมาว่ามีความถูกต้องหรือไม่
	if err := validation.ValidateLogin(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// ค้นหาข้อมูลผู้ใช้ในฐานข้อมูล
	var user models.User
	// ถ้ามีข้อมูลที่ส่งมาผิดพลาดจะส่งข้อความผิดพลาดกลับไป
	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}
	// ตรวจสอบรหัสผ่านว่ามีความถูกต้องหรือไม่ โดยใช้ฟังก์ชัน CheckPassword จาก utils
	if !utils.CheckPassword(body.Password, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Incorrect password")
	}
	// สร้างข้อมูลผู้ใช้ใหม่ในฐานข้อมูล
	token, err := utils.GenerateJWT(user.ID)
	// ถ้ามีข้อมูลที่ส่งมาผิดพลาดจะส่งข้อความผิดพลาดกลับไป
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token")
	}
	// ส่งข้อความสำเร็จกลับไป
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}
