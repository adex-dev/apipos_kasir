package mastercontrollers

import (
	"apipos/resources/Master"
	"github.com/gofiber/fiber/v2"
)

func Storeshow(c *fiber.Ctx) error {
	store, err := Master.StoreResources(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"respon":  500,
		})
	}
	if store == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Getdata",
			"respon":  401,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   store,
		"respon": 200,
	})
}
