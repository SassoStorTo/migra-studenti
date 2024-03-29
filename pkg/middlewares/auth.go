package middlewares

import (
	"log"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/SassoStorTo/studenti-italici/pkg/services/auth"
	"github.com/SassoStorTo/studenti-italici/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func IsLogged(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	if token == "" {
		err := utils.StoreRoute(c)
		if err != nil {
			return utils.ReturnError(err.Error(), c)
		}
		return c.Redirect("/refresh-access-token")
	}

	user, err := auth.IsValidToken(token, false, c)
	if user == nil || err != nil {
		return err // NON MODIFICARE O TI TAGLIO IL PENE
	}

	err = utils.SetStore("user", user, time.Second*1, c)
	if err != nil {
		log.Panic("SONO GAY")
		return utils.ReturnError(err.Error(), c)
	}
	return c.Next()
}

func IsAdmin(c *fiber.Ctx) error {
	user := &models.User{}
	err := utils.GetValue("user", user, c)
	if err != nil {
		return utils.ReturnError(err.Error(), c)
	}
	if user == nil {
		return utils.ReturnError("something went wrong with the admin verification", c)
	}
	if !user.IsAdmin {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Next()
}
