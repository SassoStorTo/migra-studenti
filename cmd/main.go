package main

import (
	"log"

	"github.com/SassoStorTo/studenti-italici/pkg/database"
	"github.com/SassoStorTo/studenti-italici/pkg/middlewares"
	"github.com/SassoStorTo/studenti-italici/pkg/router"
	dbutils "github.com/SassoStorTo/studenti-italici/pkg/services/databaseutils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	middlewares.InitStoreSess()
	err := database.ConnectDB()
	if err != nil {
		log.Panic(err)
	}
	dbutils.Reset()
	dbutils.SetupDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Prefork:       false, // questo tipo spwna tante volte il processo
		CaseSensitive: false,
		StrictRouting: true,
		AppName:       "italici",
		Views:         engine,
		// ViewsLayout:   "frontends/mypages/template",
	})

	router.SetUpRoutes(app)

	app.Listen(":8080")
}

// func srtupStuff() {
// dbutils.Reset()

// class := models.NewClass(5, "I", 2023, 1)
// err = class.Save()
// if err != nil {
// 	log.Panic(err)
// }

// t := time.Now()
// e := models.NewStuent("elia", "soldati", t)
// e.Save()
// models.NewStudentClass(1, 1, time.Now()).Save()

// (*e).Name = "paolo"
// (*e).Id = 1
// err = (*e).Update()
// if err != nil {
// 	log.Panic(err)
// }

// (*class).Section = "tre"
// (*class).Id = 1
// (*class).Update()

// major := models.Majors{Id: 1}
// err = major.Delete()
// if err != nil {
// 	log.Panic(err)
// }

// err = e.Delete()
// if err != nil {
// 	log.Panic(err)
// }

// err = (*class).Delete()
// if err != nil {
// 	log.Panic(err)
// }

// }
