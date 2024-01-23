package isoidecontrollers

import (
	"apipos/resources/isoide"
	"github.com/gofiber/fiber/v2"
)

func Kategori(c *fiber.Ctx) error {
	status := c.Params("status")
	return isoide.KategoriResource(status, c)

}
