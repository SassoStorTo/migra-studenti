package services //Todo: magari qua mettere un nome piu' carinos

import (
	"log"

	"github.com/SassoStorTo/studenti-italici/api/database"
	"github.com/SassoStorTo/studenti-italici/api/tables"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
)

func ExistTable(name string) bool { // Todo: crate a check function
	rows, err := database.DB.Query(`
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_name = ($1)
		);
	`, name)
	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()
	rows.Next()
	var result bool
	err = rows.Scan(&result)
	if err != nil {
		log.Panic(err.Error())
	}
	return result
}

func SetupDb() {
	if !ExistTable("majors") {
		log.Printf("[handlers] setup: majors not found")
		ExecQuery(tables.QueryCreateMajors())
		models.NewMajor("Informatica").Save()
		models.NewMajor("Biotecnologie Ambientali").Save()
		models.NewMajor("Automazione").Save()
	}

	if !ExistTable("classes") {
		log.Printf("[handlers] setup: classes not found")
		ExecQuery(tables.QueryCreateClasses())
	}

	if !ExistTable("students") {
		log.Printf("[handlers] setup: students not found")
		ExecQuery(tables.QueryCreateStudents())
	}

	if !ExistTable("studentclass") {
		log.Printf("[handlers] setup: studentclass not found")
		ExecQuery(tables.QueryCreateStudentClass())
	}
}

func ExecQuery(query string) {
	_, err := database.DB.Exec(query)

	if err != nil {
		log.Panic(err)
	}
}

func CheckIdInTable(table string, id int) bool { // Todo: controlla per injection
	rows, err := database.DB.Query(`
		SELECT EXISTS (
			SELECT 1
			FROM ($1)
			WHERE Id = ($2)
		);`, table, id)
	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()
	rows.Next()
	var result bool
	err = rows.Scan(&result)
	if err != nil {
		log.Panic(err.Error())
	}
	return result
}

func Reset() {
	ExecQuery(`
		DROP TABLE IF EXISTS Classes CASCADE; 
		DROP TABLE IF EXISTS Majors CASCADE; 
		DROP TABLE IF EXISTS Students CASCADE; 
		DROP TABLE IF EXISTS StudentClass CASCADE; 
	`)
}
