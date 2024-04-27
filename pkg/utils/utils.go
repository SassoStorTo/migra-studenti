package utils

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ReturnError(message string, c *fiber.Ctx) error {
	errStr := fmt.Sprintf("qualcosa e' andato storto [%s] ... contattare il programmmatore", message)
	return c.Status(fiber.StatusInternalServerError).SendString(errStr)
}

func ConvertNullInt64ToInt(n sql.NullInt64) int {
	if n.Valid {
		return int(n.Int64) // Cast to int, be cautious about the potential overflow on 32-bit systems.
	}
	return 0 // Default value or you could choose another default or even return an error.
}

func ConvertNullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String // Directly return the string if it is valid.
	}
	return "" // Return an empty string as default, or you can customize this part.
}
