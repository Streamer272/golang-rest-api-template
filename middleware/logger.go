package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/logger"
	"time"
)

func LogOnMiddleWare(c *fiber.Ctx) error {
	startTime := time.Now()

	err := c.Next()
	if err != nil {
		panic(err)
	}

	logger.Log(fmt.Sprintf("[%v LOG %v] %v ---> %v\n", time.Now().Format("02-01-2006 15:04:05"), time.Since(startTime), c.Route().Method, c.Path()))

	return nil
}
