package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userVal := c.Locals("user")
		if userVal == nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing or invalid token")
		}

		userToken, ok := userVal.(*jwt.Token)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
		}

		claims, ok := userToken.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
		}

		roleVal, ok := claims["role"]
		if !ok {
			return fiber.NewError(fiber.StatusForbidden, "Role not found in token")
		}

		role, ok := roleVal.(string)
		if !ok || role != "admin" {
			return fiber.NewError(fiber.StatusForbidden, "Admin access only")
		}

		return c.Next()
	}
}


