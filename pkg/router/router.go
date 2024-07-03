package router

import (
	"github.com/SassoStorTo/migra-studenti/pkg/handlers"
	"github.com/SassoStorTo/migra-studenti/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

type PageData struct {
	Title   string
	Content string
}

func SetUpRoutes(app *fiber.App) {
	app.Static("/", "./public")

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

	////////////////////

	user.Get("/classes", handlers.GetAllClasses)
	user.Get("/classes/create", handlers.GetCreateClassForm)
	user.Post("/classes/create", handlers.AddNewClass)

	user.Post("/upload", handlers.UploadFile)
	user.Get("/upload", func(c *fiber.Ctx) error {
		return c.Render("classes/marko_gay", fiber.Map{}, "template")
	})

	user.Get("/classes/:id", handlers.GetClassInfo)

	user.Put("/classes/:id", handlers.SaveEditClass)
	api.Get("/compoent/classes-edit/:id", handlers.GetFomrComponentEditClass)
	api.Get("/compoent/classes-display/:id", handlers.GetFomrComponentDisplayClass)
	api.Get("/compoent/classes-migration/:id", handlers.GetStudentClassMigration)
	api.Post("/compoent/classes-migration-edit/:id", handlers.GetStudentClassMigrationEdit)
	api.Get("/compoent/classes-display-students/:id", handlers.GetTablesStudentsOfClass)
	api.Post("/compoent/classes-display-students-update/:id", handlers.ClassMigrationRefreshPage)

	////////////////////

	user.Get("/majors", handlers.GetTableMajors)
	user.Get("/majors/create", handlers.GetCreateMajorForm)
	user.Post("/majors", handlers.AddNewMajor)
	user.Delete("/majors/:id", handlers.DeleteMajor)

	////////////////////

	user.Get("/students", handlers.GetTablesStudents)
	user.Get("/students/create", handlers.GetCreateStuduentForm)
	user.Post("/students/create", handlers.AddNewStudent)
	user.Put("/students/edit/:id", handlers.SaveEditStudent)
	user.Delete("/students/delete/:id", handlers.DeleteStudent)
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

	admin := user.Group("/admin", middlewares.IsAdmin)
	admin.Get("/users", handlers.GetUserPage)

	adminApi := user.Group("/api/admin", middlewares.IsAdmin)
	adminApi.Get("/compoent/user-row-edit", handlers.GetUserEditRow)
	adminApi.Get("/compoent/user-row-edit-partial", handlers.GetUserEditRowPartialEdited)
	adminApi.Get("/compoent/user-row/:id", handlers.GetUserRow)

	adminApi.Post("/change-status", handlers.SetStatus)

	/////////////

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{}, "template")
	})
}
