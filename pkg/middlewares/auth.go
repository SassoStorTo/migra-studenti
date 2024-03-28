package middlewares

import (
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/services/auth"
	"github.com/SassoStorTo/studenti-italici/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func IsLogged(c *fiber.Ctx) error {
	// log.Printf("Entrato dentro IsLogged %s\n", c.Route().Path)
	token := c.Cookies("access_token")
	if token == "" {
		err := utils.SetStore("original-route", c.Route().Path, time.Minute*2, c)
		if err != nil {
			return utils.ReturnError(err.Error(), c)
		}
		return c.Redirect("/refresh-access-token")
	}

	user, err := auth.IsValidToken(token, false, c)
	if err != nil {
		return err
	}
	// Todo: fixa questa cosa che esplode forte & aggiungila ad RefreshaccessToken
	err = utils.SetStore("user", user, time.Second*2, c)
	if err != nil {
		return utils.ReturnError(err.Error(), c)
	}
	return c.Next()
}

func IsAdmin(c *fiber.Ctx) error {
	return nil
}
