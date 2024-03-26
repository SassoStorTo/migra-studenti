package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ExecQuery(query string) {
	_, err := DB.Exec(query)

	if err != nil {
		log.Panic(err)
	}
}

func CheckIdInTable(table string, id int) bool { // Todo: controlla per injection
	rows, err := DB.Query(`
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

func ExistTable(name string) bool { // Todo: crate a check function
	rows, err := DB.Query(`
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

func FormatTimeForDb(date time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())
}

