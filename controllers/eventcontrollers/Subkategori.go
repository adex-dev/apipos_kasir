package eventcontrollers

import (
	"apipos/resources/eventt"
	"github.com/gofiber/fiber/v2"
)

func Subkategori(c *fiber.Ctx) error {
	return eventt.SubkategoriResource(c)
}
