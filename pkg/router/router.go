package router

import (
	"github.com/SassoStorTo/studenti-italici/pkg/handlers"
	"github.com/SassoStorTo/studenti-italici/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

type PageData struct {
	Title   string
	Content string
}

func SetUpRoutes(app *fiber.App) {
	app.Static("/", "./public")

	app.Get("/par", func(c *fiber.Ctx) error {
		return c.Render("tables", fiber.Map{})
	})
	// app.Get("/mar", func(c *fiber.Ctx) error {
	// 	buf := new(bytes.Buffer)
	// 	views.Testazzo().Render(c.Context(), buf)
	// 	html := buf.String()
	// 	return c.Render("render_helper", fiber.Map{"content": html}, "template")
	// })

	api := app.Group("/api", middlewares.IsLogged)

	authGroup := app.Group("/auth")
	authGroup.Get("/callback", handlers.HandleCallback)

	app.Get("/wait-accept", handlers.WaitToAccept)
	app.Get("/login", handlers.HandleLogin)
	app.Get("/refresh-access-token", handlers.RefreshAccessToken)

	user := app.Group("/", middlewares.IsLogged)
	user.Get("/", func(c *fiber.Ctx) error {
		return c.Render("i", fiber.Map{}, "template")
	})

	user.Get("/classes", handlers.GetAllClasses)
	user.Get("/classes/create", handlers.GetCreateClassForm)
	user.Post("/classes/create", handlers.AddNewClass)

	user.Get("/majors", handlers.GetTableMajors)
	user.Get("/majors/create", handlers.GetCreateMajorForm)
	user.Post("/majors", handlers.AddNewMajor)
	user.Delete("/majors/:id", handlers.DeleteMajor)

	user.Get("/students", handlers.GetTablesStudents)
	user.Get("/students/create", handlers.GetCreateStuduentForm)
	user.Post("/students/create", handlers.AddNewStudent)
	user.Put("/students/edit/:id", handlers.SaveEditStudent)
	user.Get("/students/:id", handlers.GetStudentInfo)
	api.Get("/compoent/student-edit/:id", handlers.GetFomrComponentEditStudent)
	api.Get("/compoent/student-display/:id", handlers.GetFomrComponentDisplayStudent)

	////////////////////
	////////////////////
	////////////////////
	////////////////////

	// // user.Post("/student", students.Create)
	// // user.Post("/major", majors.Create)
	// user.Post("/studentclass", studentclass.Create)

	// user.Delete("/student", students.Delete)
	// user.Delete("/class", classes.Delete)
	// // user.Delete("/major", majors.Delete)
	// user.Delete("/studentclass", studentclass.Delete)

	// user.Put("/student", students.Edit)
	// user.Put("/class", classes.Edit)
	// user.Put("/major", majors.Edit)

	////////////////////
	////////////////////
	////////////////////
	////////////////////

	admin := user.Group("/admin", middlewares.IsAdmin) // Todo: add check

	admin.Get("/users", handlers.GetUserPage)
	admin.Get("/negro", func(c *fiber.Ctx) error {
		return c.Render("i", fiber.Map{}, "template")
	})

	adminApi := user.Group("/api/admin", middlewares.IsAdmin)
	adminApi.Get("/compoent/user-row-edit", handlers.GetUserEditRow)
	adminApi.Get("/compoent/user-row-edit-partial", handlers.GetUserEditRowPartialEdited)
	adminApi.Get("/compoent/user-row", handlers.GetUserRow)

	adminApi.Post("/change-status", handlers.SetStatus)

	/////////////

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
			"ciao": "Hello, World!",
		}, "template")
	})

	app.Get("/par", func(c *fiber.Ctx) error {
		return c.Render("tables", fiber.Map{})
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

// Use this tool to generate front-end code snippets in different languages.

// 1. 🪄 Add **/expand** to any of the prompts to front-end code snippets.
// 2. 🛠️ Customize existing prompts or add your own.
// 3. 🔘 Click **Insert** to use the generated chart or **Regenerate** to start again.
// 4. 🤖 (optional) Type **/assistant** to see all [Taskade AI commands](https://www.taskade.com/blog/introducing-taskade-ai/).

// # 🟡 HTML Code Generator

// - Generate HTML code for a basic webpage structure.
// - Generate HTML code for a webpage with a header, main section, and footer.
// - Generate HTML code for a list of users where for each one I can choose their permissions
// - Generate HTML code for a table for viewing every student in a class.
// - Generate HTML code for a table for viewing every class of the school.
// - Generate HTML code for a table for viewing every student in the school.
// - Generate HTML code for a view of the pasts years situations, like the total number of students and the number of students gained and lost by a class.
// - Generate HTML code for a responsive navigation bar.

// # 🟢 CSS Code Generator

// - Generate CSS code for a responsive navigation bar.
// - Generate CSS code for a webpage with a grid layout.
// - Generate CSS code for a webpage with a flexbox layout.
// - Generate CSS code for styling form elements.
// - Generate CSS code for a webpage with dynamic themes (light/dark mode).
// - Generate CSS code for a webpage with smooth scrolling animations.
// - Generate CSS code for an interactive hover effect on buttons.
// - Generate CSS code for a webpage with media queries for different screen sizes.
// - Generate CSS code for a website with a sticky header.
// - Generate CSS code for a parallax scrolling effect.
