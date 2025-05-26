package controller

import (
	"backend/service"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	return service.HandleCreateProduct(c)
}

func GetAllProducts(c *fiber.Ctx) error {
	return service.HandleGetAllProducts(c)
}

func GetProductBySlug(c *fiber.Ctx) error {
	return service.HandleGetProductBySlug(c)
}

func UpdateProduct(c *fiber.Ctx) error {
	return service.HandleUpdateProduct(c)
}

func DeleteProduct(c *fiber.Ctx) error {
	return service.HandleDeleteProduct(c)
}
