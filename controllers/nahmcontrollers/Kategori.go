package nahmcontrollers

import (
	"apipos/resources/nahm"
	"github.com/gofiber/fiber/v2"
)

func Kategori(c *fiber.Ctx) error {
	return nahm.KategoriResource(c)
}
