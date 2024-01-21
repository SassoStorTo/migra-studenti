package main

import (
	"log"
	"time"

	"github.com/SassoStorTo/studenti-italici/api/database"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
	dbutils "github.com/SassoStorTo/studenti-italici/pkg/services/databaseutils"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Panic(err)
	}
	dbutils.Reset()
	dbutils.SetupDb()

	class := models.NewClass(5, "I", 2023, 1)
	err = class.Save()
	if err != nil {
		log.Panic(err)
	}

	t := time.Now()
	e := models.NewStuent("elia", "soldati", t)
	e.Save()
	models.NewStudentClass(1, 1, time.Now()).Save()

	(*e).Name = "paolo"
	(*e).Id = 1
	err = (*e).Update()
	if err != nil {
		log.Panic(err)
	}

	(*class).Section = "tre"
	(*class).Id = 1
	(*class).Update()

	major := models.Majors{Id: 1}
	err = major.Delete()
	if err != nil {
		log.Panic(err)
	}

	// err = e.Delete()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// err = (*class).Delete()
	// if err != nil {
	// 	log.Panic(err)
	// }

	//////////////////

	// app := fiber.New(fiber.Config{
	// 	Prefork:       true,
	// 	CaseSensitive: true,
	// 	StrictRouting: true,
	// 	AppName:       "italici",
	// })

	// router.SetUpRoutes(app)

	// app.Listen(":8080")
}
