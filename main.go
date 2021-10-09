package main

import (
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"password-manager-backend/database"
	"password-manager-backend/routes"
	"time"
)

func main() {
	database.Init()
	rand.Seed(time.Now().UnixNano())

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Password Manager",
		AppName:       "Password Manager",
	})

	app.Server().MaxConnsPerIP = 1

	routes.Setup(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
