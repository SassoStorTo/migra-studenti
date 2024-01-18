package main

import (
	"github.com/studenti-italici/api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// database.ConnectDB()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "italici",
	})

	router.SetUpRoutes(app)

	app.Listen(":8080")
}
