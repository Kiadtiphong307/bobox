package admin

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"log"
)

// คือฟังก์ชัน seed user admin สำหรับทดสอบระบบ
func SeedAdminUser() {
	var count int64
	database.DB.Model(&models.User{}).Where("email = ?", "admin@example.com").Count(&count)
	if count == 0 {
		hashedPassword := utils.HashPassword("admin123") // ใช้ฟังก์ชันเข้ารหัสรหัสผ่าน

		adminUser := models.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: hashedPassword,
			Role:     "admin",
		}

		if err := database.DB.Create(&adminUser).Error; err != nil {
			log.Println("❌ Failed to create admin user:", err)
		} else {
			log.Println("✅ Admin user created successfully")
		}
	} else {
		log.Println("ℹ️ Admin user already exists, skipping seed")
	}
}
