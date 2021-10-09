package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"password-manager-backend/controllers"
	"password-manager-backend/errors"
	"password-manager-backend/logger"
	"password-manager-backend/middleware"
)

func Setup(app *fiber.App) {
	app.Use(middleware.Cors)

	app.Use(middleware.LogOnMiddleWare)
	app.Use(errors.HandleException)

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        50,
		Expiration: 0,
		LimitReached: func(c *fiber.Ctx) error {
			logger.LogError(fiber.ErrTooManyRequests)

			c.Status(fiber.StatusTooManyRequests)
			return c.JSON(errors.ErrorMessage{
				Error:   "TooManyRequests",
				Message: "",
			})
		},
	}))

	app.Get("/", controllers.Welcome)

	app.Use(func(c *fiber.Ctx) error {
		return nil
	})
}
