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

func GetProductById(c *fiber.Ctx) error {
	return service.HandleGetProductById(c)
}

func UpdateProduct(c *fiber.Ctx) error {
	return service.HandleUpdateProduct(c)
}

func DeleteProduct(c *fiber.Ctx) error {
	return service.HandleDeleteProduct(c)
}
