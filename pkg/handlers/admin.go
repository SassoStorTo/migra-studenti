package handlers

import (
	"fmt"
	"strconv"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func SetStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return fmt.Errorf("[Handlers] Verify: id field incorrect")
	}

	isVerified, err := strconv.Atoi(c.FormValue("is-verified"))
	if err != nil {
		return fmt.Errorf("[Handlers] Verify: is-verified field incorrect")
	}

	isAdmin, err := strconv.Atoi(c.FormValue("is-admin"))
	if err != nil {
		return fmt.Errorf("[Handlers] Verify: is-admin field incorrect")
	}

	user, err := models.GetUserById(id)
	if err != nil {
		return err
	}

	user.IsEditor = (isVerified != 0)
	user.IsAdmin = (isAdmin != 0)
	user.Update()

	return c.Render("admin/com_user_row", fiber.Map{"User": user})
}

func GetUserPage(c *fiber.Ctx) error {
	users, err := models.GetAllUsers()
	if err != nil {
		return err
	}
	for _, item := range users {
		fmt.Printf("%s %s \n", item.Name, item.Email)
	}

	return c.Render("admin/table_view_users", fiber.Map{"Users": users}, "template")
}

func GetUserEditRow(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return fmt.Errorf("[Handlers] GetUserEdit: id field incorrect")
	}
	user, err := models.GetUserById(id)
	if err != nil {
		return err
	}

	return c.Render("admin/com_user_row_edit", fiber.Map{"User": user})
}

func GetUserEditRowPartialEdited(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return fmt.Errorf("[Handlers] GetUserEdit: id field incorrect")
	}
	isVerified, err := strconv.Atoi(c.FormValue("is-verified"))
	if err != nil {
		return fmt.Errorf("[Handlers] Verify: is-verified field incorrect")
	}
	isAdmin, err := strconv.Atoi(c.FormValue("is-admin"))
	if err != nil {
		return fmt.Errorf("[Handlers] Verify: is-admin field incorrect")
	}

	user, err := models.GetUserById(id)
	if err != nil {
		return err
	}

	user.IsEditor = (isVerified != 0)
	user.IsAdmin = (isAdmin != 0)

	return c.Render("admin/com_user_row_edit", fiber.Map{"User": user})
}

func GetUserRow(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("[Handlers] GetUserEdit: id field incorrect")
	}
	user, err := models.GetUserById(id)
	if err != nil {
		return err
	}

	return c.Render("admin/com_user_row", fiber.Map{"User": user})
}
