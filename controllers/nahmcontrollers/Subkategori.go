package nahmcontrollers

import (
	"apipos/resources/nahm"
	"github.com/gofiber/fiber/v2"
)

func Subkategori(c *fiber.Ctx) error {
	return nahm.SubkategoriResource(c)
}
