package mastercontrollers

import (
	"apipos/resources/Master"
	"github.com/gofiber/fiber/v2"
)

func Loadsales(c *fiber.Ctx) error {
	start := c.FormValue("start")
	end := c.FormValue("end")
	result, err := Master.SalesResource(start, end)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"respon":  500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    result,
		"respon":  200,
	})
}
