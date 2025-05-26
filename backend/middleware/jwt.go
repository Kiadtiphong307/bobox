package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	
)

// คือฟังก์ชันที่จะจัดการกับการปกป้องข้อมูลผู้ใช้
func Protected() fiber.Handler {
	// สร้างข้อมูลผู้ใช้ใหม่
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey:   "user",   // เก็บ token ลงใน c.Locals("user")
		ErrorHandler: jwtError, // เรียกเมื่อ token ผิดหรือไม่มี โดยใช้ฟังก์ชัน jwtError
	})
}

// คือฟังก์ชันที่จะจัดการกับการผิดพลาดของการปกป้องข้อมูลผู้ใช้
func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "error",
		"message": "Unauthorized - invalid or missing token",
		"detail":  err.Error(),
	})
}

