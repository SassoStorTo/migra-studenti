package students

func QueryCreate() string {
	return `
		CREATE TABLE IF NOT EXISTS Students (
			Id SERIAL PRIMARY KEY,
			Name varchar(50) NOT NULL,
			LastName varchar(50) NOT NULL,
			DateOfBirth TIMESTAMP NOT NULL
		);`
}
