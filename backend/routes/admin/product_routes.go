package admin

import (
	"backend/controller"
	admin "backend/middleware/admin"
	"backend/middleware" // <-- เพิ่ม middleware ปกติด้วย
	"github.com/gofiber/fiber/v2"
)

func AdminProductRoutes(router fiber.Router) {
	api := router.Group("/admin/products",
		middleware.Protected(), // ✅ ตรวจ JWT token
		admin.CheckAdmin(),     // ✅ ตรวจว่าเป็น role = admin
	)

	api.Post("/", controller.CreateProduct)
	api.Put("/:id", controller.UpdateProduct)
	api.Delete("/:id", controller.DeleteProduct)
}
