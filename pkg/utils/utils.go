package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ReturnError(message string, c *fiber.Ctx) error {
	errStr := fmt.Sprintf("qualcosa e' andato storto [%s] ... contattare il programmmatore", message)
	return c.Status(fiber.StatusInternalServerError).SendString(errStr)
}
