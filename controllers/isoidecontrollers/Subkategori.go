package isoidecontrollers

import (
	"apipos/resources/isoide"
	"github.com/gofiber/fiber/v2"
)

func Subkategori(c *fiber.Ctx) error {
	idsub := c.Params("idsub")
	idkategori := c.Params("idkategori")
	return isoide.SubkategoriResource(idsub, idkategori, c)
}
