package authcontrollers

import (
	"apipos/resources/auth"
	"github.com/gofiber/fiber/v2"
)

func Loginauth(c *fiber.Ctx) error {
	nik := c.FormValue("nik")
	password := c.FormValue("password")
	user, err := auth.ValidasiLogin(nik, password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"respon":  500,
		})
	}
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
			"respon":  401,
		})
	}
	if user.Status != "AKTIF" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Data Sudah tidak Aktif",
			"respon":  401,
		})
	}
	response := map[string]interface{}{
		"level": user.Level,
		"nik":   user.Nik,
		"nama":  user.Nama_user,
		"app":   "poscrm",
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"data":    response,
		"respon":  200,
	})
}
