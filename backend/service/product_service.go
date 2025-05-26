package service

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"strings"
)

// คือฟังก์ชันที่จะจัดการกับการสร้างสินค้า
func HandleCreateProduct(c *fiber.Ctx) error {
	var body models.Product
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid product format")
	}

	body.Slug = strings.ToLower(strings.ReplaceAll(body.Name, " ", "-"))

	if err := database.DB.Create(&body).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create product")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": body,
	})
}

// คือฟังก์ชันที่จะจัดการกับการดึงข้อมูลสินค้าเฉพาะตาม ID
func HandleGetProductBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var product models.Product
	if err := database.DB.Where("slug = ?", slug).First(&product).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}
	return c.JSON(fiber.Map{
		"message": "Product fetched successfully",
		"product": product,
	})
}

// คือฟังก์ชันที่จะจัดการกับการดึงข้อมูลสินค้าทั้งหมด
func HandleGetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch products")
	}
	return c.JSON(fiber.Map{
		"message":  "Products fetched successfully",
		"products": products,
	})
}

// คือฟังก์ชันที่จะจัดการกับการแก้ไขสินค้า
func HandleUpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// ฟิลด์ที่อนุญาตให้อัปเดต
	allowedFields := []string{"name", "price", "stock", "description", "category_id"}

	// กรองเฉพาะฟิลด์ที่อนุญาต
	updates := make(map[string]interface{})
	for _, field := range allowedFields {
		if val, ok := body[field]; ok {
			updates[field] = val
		}
	}

	if len(updates) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "No valid fields to update")
	}

	if err := database.DB.Model(&product).Updates(updates).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update product")
	}

	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})
}

// คือฟังก์ชันที่จะจัดการกับการลบสินค้า
func HandleDeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete product")
	}
	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
