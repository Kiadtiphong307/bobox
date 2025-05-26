package database

import (
	"fmt"
	"log"
	"os"

	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// คือตัวแปรที่จะจัดการกับการเชื่อมต่อกับฐานข้อมูล
var DB *gorm.DB

// คือฟังก์ชันที่จะจัดการกับการเชื่อมต่อกับฐานข้อมูล
func ConnectDB() {
	// สร้างข้อมูลการเชื่อมต่อกับฐานข้อมูล
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// เชื่อมต่อกับฐานข้อมูล
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}
	// สร้างข้อมูลการเชื่อมต่อกับฐานข้อมูล
	DB = db
	// ส่งข้อความสำเร็จกลับไป
	log.Println("✅ Connected to MySQL")


	db.AutoMigrate(&models.User{})
}
