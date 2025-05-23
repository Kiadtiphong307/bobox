package main

import (
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() 

	database.ConnectDB()

	routes.AuthRoutes(app)

	app.Listen(":8080")
}