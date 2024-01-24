package eventcontrollers

import (
	"apipos/resources/eventt"
	"github.com/gofiber/fiber/v2"
)

func Kategori(c *fiber.Ctx) error {
	return eventt.KategoriResource(c)
}
