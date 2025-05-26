package controller

import (
	"backend/service"
	"github.com/gofiber/fiber/v2"
)

// คือฟังก์ชันที่จะจัดการกับการลงทะเบียนของผู้ใช้ 
// จะเรียกฟังก์ชัน HandleRegister จาก service เพื่อจัดการกับการลงทะเบียนของผู้ใช้
func Register(c *fiber.Ctx) error {
	return service.HandleRegister(c)
}

// คือฟังก์ชันที่จะจัดการกับการลงชื่อเข้าใช้งานของผู้ใช้
// จะเรียกฟังก์ชัน HandleLogin จาก service เพื่อจัดการกับการลงชื่อเข้าใช้งานของผู้ใช้
func Login(c *fiber.Ctx) error {
	return service.HandleLogin(c)
}
