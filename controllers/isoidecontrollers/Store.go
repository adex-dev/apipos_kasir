package isoidecontrollers

import (
	"apipos/resources/Master"
	"github.com/gofiber/fiber/v2"
)

func Storeisoide(c *fiber.Ctx) error {
	brands := "ISOIDE"
	status := c.FormValue("status")
	store, err := Master.StoreResources(brands, status)
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
	var response []map[string]interface{}
	for _, store := range store {
		storeMap := map[string]interface{}{
			"id_outlet":   store.Idoutlet,
			"nama_outlet": store.Namaoutlet,
		}
		response = append(response, storeMap)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   response,
		"respon": 200,
	})
}
