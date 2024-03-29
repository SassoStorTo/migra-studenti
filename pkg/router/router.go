package router

import (
	"github.com/SassoStorTo/studenti-italici/pkg/handlers"
	"github.com/SassoStorTo/studenti-italici/pkg/middlewares"
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
	app.Static("/", "./public")

	authGroup := app.Group("/auth")
	authGroup.Get("/callback", handlers.HandleCallback)

	app.Get("/wait-accept", handlers.WaitToAccept)
	app.Get("/login", handlers.HandleLogin)
	app.Get("/refresh-access-token", handlers.RefreshAccessToken)

	user := app.Group("/", middlewares.IsLogged)
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

	admin := user.Group("/admin", middlewares.IsAdmin) // Todo: add check
	admin.Post("/verify", handlers.SetVerify)

	admin.Get("/negro", func(c *fiber.Ctx) error {
		return c.Render("i", fiber.Map{}, "template")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
			"ciao": "Hello, World!",
		}, "template")
	})
}

// app.Get("/sandro", func(c *fiber.Ctx) error {
// 	user := &models.User{}
// 	err := utils.GetValue("sandro", user, c)
// 	if err != nil {
// 		return utils.ReturnError(err.Error(), c)
// 	}

// 	if user == nil {
// 		return c.SendString("Mario e' gays")
// 	}

// 	return c.SendString("ValuE: " + fmt.Sprintf("%s", user.Email))
// })
// app.Get("/gennaro", func(c *fiber.Ctx) error {
// 	gennaro := models.User{
// 		Id:            1,
// 		Email:         "sa@gmail.com",
// 		Name:          "sa",
// 		Hd:            "sas",
// 		VerifiedEmail: false,
// 		IsAdmin:       false,
// 		IsEditor:      false,
// 	}
// 	errs := utils.SetStore("sandro", gennaro, time.Second*30, c)
// 	if errs != nil {
// 		return utils.ReturnError(errs.Error(), c)
// 	}

// 	return c.SendString("SETTATTO TUTTO TUTTO")
// })

//////////////////////////////////

// admin := app.Group("/admin", middlewares.IsLogged, middlewares.IsAdmin) // Todo: add check
// admin.Post("/verify", handlers.SetVerify)

// admin.Get("/negro", func(c *fiber.Ctx) error {
// 	return c.Render("i", fiber.Map{}, "template")
// })

///////////////////////////////////////////

// app.Get("/par", func(c *fiber.Ctx) error {
// 	return c.Render("tables", fiber.Map{})
// })
