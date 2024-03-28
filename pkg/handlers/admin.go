package handlers

import (
	"fmt"
	"strconv"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func SetVerify(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return fmt.Errorf("[Handlers] Verify: id field incorrect")
	}

	isVerified, err := strconv.Atoi(c.FormValue("is-verified"))
	if err != nil {
		return fmt.Errorf("[Handlers] Verify: is-verified field incorrect")
	}

	user, err := models.GetUserById(id)
	if err != nil {
		return err
	}

	user.IsEditor = (isVerified != 0)
	user.Update()

	return c.SendStatus(fiber.StatusAccepted)
}
