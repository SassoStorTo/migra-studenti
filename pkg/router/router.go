package router

import (
	"github.com/SassoStorTo/studenti-italici/pkg/handlers"
	"github.com/SassoStorTo/studenti-italici/pkg/services/auth"
	"github.com/SassoStorTo/studenti-italici/pkg/services/classes"
	"github.com/SassoStorTo/studenti-italici/pkg/services/majors"
	"github.com/SassoStorTo/studenti-italici/pkg/services/studentclass"
	"github.com/SassoStorTo/studenti-italici/pkg/services/students"
	"github.com/gofiber/fiber/v2"
)

type PageData struct {
	Title   string
	Content string
}

func SetUpRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Get("/login", handlers.Login)
	authGroup.Get("/callback", auth.HandleCallback)
	authGroup.Get("/login/google", auth.HandleLogin)

	app.Static("/", "./public")

	app.Get("/par", func(c *fiber.Ctx) error {
		return c.Render("tables", fiber.Map{})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("i", fiber.Map{}, "template")
	})

	// api := app.Group("/api")
	// v1 := api.Group("v1")
	app.Post("/student", students.Create)
	app.Post("/class", classes.Create)
	app.Post("/major", majors.Create)
	app.Post("/studentclass", studentclass.Create)

	app.Delete("/student", students.Delete)
	app.Delete("/class", classes.Delete)
	app.Delete("/major", majors.Delete)
	app.Delete("/studentclass", studentclass.Delete)

	app.Put("/student", students.Edit)
	app.Put("/class", classes.Edit)
	app.Put("/major", majors.Edit)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
			"ciao": "Hello, World!",
		}, "template")
	})

}
