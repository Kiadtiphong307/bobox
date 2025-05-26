package routes

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
	api := router.Group("/products")
	api.Get("/", controller.GetAllProducts)
	api.Get("/:slug", controller.GetProductBySlug)

}
