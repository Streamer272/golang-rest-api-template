package middleware

import "github.com/gofiber/fiber/v2"

func Cors(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
	c.Set(fiber.HeaderAccessControlAllowCredentials, "true")

	c.Vary(fiber.HeaderAccessControlAllowOrigin)
	c.Vary(fiber.HeaderAccessControlAllowCredentials)
	c.Vary(fiber.HeaderOrigin)
	c.Vary(fiber.HeaderAccessControlRequestMethod)
	c.Vary(fiber.HeaderAccessControlRequestHeaders)

	return c.Next()
}
