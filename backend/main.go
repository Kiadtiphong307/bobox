package main

import (
	"backend/database"
	"backend/routes"
	admin "backend/routes/admin"
	adminSeed "backend/seed/admin"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// เชื่อมต่อกับฐานข้อมูล
	database.ConnectDB()

	// เรียกใช้งาน routes ของ admin product
	admin.AdminProductRoutes(app) // เรียกใช้งาน routes ของ admin product

	// เรียกใช้งาน routes
	routes.AuthRoutes(app)    // เรียกใช้งาน routes ของ auth
	routes.ProductRoutes(app) // เรียกใช้งาน routes ของ product

	// เรียกใช้งาน seed admin user
	adminSeed.SeedAdminUser()

	app.Listen(":8080")
}
