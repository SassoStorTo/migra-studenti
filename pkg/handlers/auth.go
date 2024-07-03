package handlers

import (
	"github.com/SassoStorTo/migra-studenti/pkg/services/auth"
	"github.com/SassoStorTo/migra-studenti/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// fai questo, poi vai da google e fixa i cookies la
func RefreshAccessToken(c *fiber.Ctx) error {
	token := c.Cookies("refresh_token")
	if token == "" {
		err := utils.StoreRoute(c)
		if err != nil {
			return utils.ReturnError(err.Error(), c)
		}
		return c.Redirect("/login") //Todo: check if it's the correct route
	}

	usr, err := auth.IsValidToken(token, true, c)
	if usr == nil || err != nil {
		return err // NON MODIFICARE O TI TAGLIO IL PENE
	}

	token, exp, err := auth.GetAccessToken(usr)
	if err != nil {
		return utils.ReturnError(err.Error(), c)
	}

	c.Cookie(&fiber.Cookie{
		Expires:  exp,
		Secure:   true,
		HTTPOnly: true, // accessible only by http (not js)
		Name:     "access_token",
		Value:    token,
	})

	return c.Redirect("/") // todo: sure about this?
}

func Login(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusInternalServerError)
}

func WaitToAccept(c *fiber.Ctx) error {
	cookieValue := c.Cookies("access_token")
	if cookieValue != "" {
		user, err := auth.ParseToken(cookieValue)
		if err == nil {
			if user.IsEditor {
				return c.Redirect("/")
			}
		}
	}
	return c.SendString("Aspettare l'autenticazione da parte di un admin")
}
