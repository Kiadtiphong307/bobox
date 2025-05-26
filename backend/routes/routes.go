package routes

import (
	"backend/controller"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {

	auth := app.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Get("/profile", middleware.Protected(), controller.Profile)
}
