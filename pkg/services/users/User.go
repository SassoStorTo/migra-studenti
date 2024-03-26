package users

func QueryCreate() string {
	return `
		CREATE TABLE IF NOT EXISTS Users (
			Id SERIAL PRIMARY KEY,
			Name VARCHAR(50),
			Email VARCHAR(150) NOT NULL UNIQUE,
			Hd VARCHAR(100) NOT NULL,
			VerifiedEmail BOOL NOT NULL,
			IsAdmin BOOL DEFAULT FALSE,
			IsEditor BOOL DEFAULT FALSE
		)`
}
