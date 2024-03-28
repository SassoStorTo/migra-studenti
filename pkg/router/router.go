package router

import (
	"fmt"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/handlers"
	"github.com/SassoStorTo/studenti-italici/pkg/middlewares"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/SassoStorTo/studenti-italici/pkg/services/auth"
	"github.com/SassoStorTo/studenti-italici/pkg/services/classes"
	"github.com/SassoStorTo/studenti-italici/pkg/services/majors"
	"github.com/SassoStorTo/studenti-italici/pkg/services/studentclass"
	"github.com/SassoStorTo/studenti-italici/pkg/services/students"
	"github.com/SassoStorTo/studenti-italici/pkg/utils"
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
	// authGroup.Get("/refresh-access-token", handlers.RefreshAccessToken)

	app.Static("/", "./public")

	app.Get("/par", func(c *fiber.Ctx) error {
		return c.Render("tables", fiber.Map{})
	})
	app.Get("/wait-accept", handlers.WaitToAccept)
	app.Get("/refresh-access-token", handlers.RefreshAccessToken)
	app.Get("/sandro", func(c *fiber.Ctx) error {
		value, err := utils.GetValue("sandro", c)
		if err != nil {
			return utils.ReturnError(err.Error(), c)
		}

		if value == nil {
			return c.SendString("Mario e' gays")
		}

		return c.SendString("ValuE: " + fmt.Sprintf("%v", value))
	})
	app.Get("/gennaro", func(c *fiber.Ctx) error {
		gennaro := models.User{
			Id:            1,
			Email:         "sa@gmail.com",
			Name:          "sa",
			Hd:            "sas",
			VerifiedEmail: false,
			IsAdmin:       false,
			IsEditor:      false,
		}
		errs := utils.SetStore("sandro", gennaro, time.Second*30, c)
		if errs != nil {
			return utils.ReturnError(errs.Error(), c)
		}

		return c.SendString("SETTATTO TUTTO TUTTO")
	})

	user := app.Group("/", middlewares.IsLogged)

	admin := user.Group("/", middlewares.IsAdmin) // Todo: add check
	admin.Post("/verify", handlers.SetVerify)

	user.Get("/", func(c *fiber.Ctx) error {
		return c.Render("i", fiber.Map{}, "template")
	})

	user.Post("/student", students.Create)
	user.Post("/class", classes.Create)
	user.Post("/major", majors.Create)
	user.Post("/studentclass", studentclass.Create)

	user.Delete("/student", students.Delete)
	user.Delete("/class", classes.Delete)
	user.Delete("/major", majors.Delete)
	user.Delete("/studentclass", studentclass.Delete)

	user.Put("/student", students.Edit)
	user.Put("/class", classes.Edit)
	user.Put("/major", majors.Edit)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
			"ciao": "Hello, World!",
		}, "template")
	})

}
