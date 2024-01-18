package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/login")
}
