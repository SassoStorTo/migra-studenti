package dbutils

import (
	"log"

	"github.com/SassoStorTo/migra-studenti/pkg/database"
	"github.com/SassoStorTo/migra-studenti/pkg/models"
	"github.com/SassoStorTo/migra-studenti/pkg/services/classes"
	"github.com/SassoStorTo/migra-studenti/pkg/services/majors"
	"github.com/SassoStorTo/migra-studenti/pkg/services/studentclass"
	"github.com/SassoStorTo/migra-studenti/pkg/services/students"
	"github.com/SassoStorTo/migra-studenti/pkg/services/users"
)

func SetupDb() {
	if !database.ExistTable("majors") {
		log.Printf("[handlers] setup: majors not found")
		database.ExecQuery(majors.QueryCreate())
		// models.NewMajor("Informatica").Save()
		// models.NewMajor("Biotecnologie Ambientali").Save()
		// models.NewMajor("Automazione").Save()
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

	if !database.ExistTable("users") {
		log.Printf("[handlers] setup: users not found")
		database.ExecQuery(users.QueryCreate())
	}

	database.ExecQuery(`DROP VIEW IF EXISTS AllActiveStudentsClass;`)
	database.ExecQuery(`CREATE VIEW AllActiveStudentsClass AS
							SELECT S.Id, S.Name, S.LastName, S.DateOfBirth, SC.IdS, SC.IdC, SC.CreationDate
							FROM students AS S INNER JOIN
								studentclass AS SC ON S.Id = SC.IdS
							WHERE SC.CreationDate = (SELECT MAX(CreationDate)
														FROM studentclass
														WHERE IdS = S.Id);`)

	database.ExecQuery(`DROP VIEW IF EXISTS AllOldStudentsClass;`)
	database.ExecQuery(`CREATE VIEW AllOldStudentsClass AS
							SELECT S.Id, S.Name, S.LastName, S.DateOfBirth, SC.IdS, SC.IdC, SC.CreationDate
							FROM students AS S INNER JOIN
								studentclass AS SC ON S.Id = SC.IdS
							WHERE SC.CreationDate <> (SELECT MAX(CreationDate)
														FROM studentclass
														WHERE IdS = S.Id);`)

}

func SetupDbBono() {
	if !database.ExistTable("majors") {
		log.Printf("[handlers] setup: majors not found")
		database.ExecQuery(majors.QueryCreate())
		models.NewMajor("Informatica").Save()
		models.NewMajor("Biotecnologie Ambientali").Save()
		models.NewMajor("Automazione").Save()
	}

	log.Printf("[handlers] setup: classes not found")
	database.ExecQuery(classes.QueryCreate())

	log.Printf("[handlers] setup: students not found")
	database.ExecQuery(students.QueryCreate())

	log.Printf("[handlers] setup: studentclass not found")
	database.ExecQuery(studentclass.QueryCreate())

}

func Reset() {
	database.ExecQuery(`
		DROP TABLE IF EXISTS Classes CASCADE; 
		DROP TABLE IF EXISTS Majors CASCADE; 
		DROP TABLE IF EXISTS Students CASCADE; 
		DROP TABLE IF EXISTS StudentClass CASCADE; 
		DROP TABLE IF EXISTS Users CASCADE; 
	`)
}
