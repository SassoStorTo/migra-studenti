package dbutils

import (
	"log"

	"github.com/SassoStorTo/studenti-italici/api/database"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/SassoStorTo/studenti-italici/pkg/services/classes"
	"github.com/SassoStorTo/studenti-italici/pkg/services/majors"
	"github.com/SassoStorTo/studenti-italici/pkg/services/studentclass"
	"github.com/SassoStorTo/studenti-italici/pkg/services/students"
)

func SetupDb() {
	if !database.ExistTable("majors") {
		log.Printf("[handlers] setup: majors not found")
		database.ExecQuery(majors.QueryCreate())
		models.NewMajor("Informatica").Save()
		models.NewMajor("Biotecnologie Ambientali").Save()
		models.NewMajor("Automazione").Save()
	}

	if !database.ExistTable("classes") {
		log.Printf("[handlers] setup: classes not found")
		database.ExecQuery(classes.QueryCreate())
	}

	if !database.ExistTable("students") {
		log.Printf("[handlers] setup: students not found")
		database.ExecQuery(students.QueryCreate())
	}

	if !database.ExistTable("studentclass") {
		log.Printf("[handlers] setup: studentclass not found")
		database.ExecQuery(studentclass.QueryCreate())
	}
}

func Reset() {
	database.ExecQuery(`
		DROP TABLE IF EXISTS Classes CASCADE; 
		DROP TABLE IF EXISTS Majors CASCADE; 
		DROP TABLE IF EXISTS Students CASCADE; 
		DROP TABLE IF EXISTS StudentClass CASCADE; 
	`)
}
