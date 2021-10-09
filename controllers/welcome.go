package controllers

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	c.Status(200)
	c.SendString("Welcome!")

	return c.Next()
}
