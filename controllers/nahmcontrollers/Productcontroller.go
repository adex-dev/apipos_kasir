package nahmcontrollers

import (
	"apipos/resources/nahm"
	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {
	return nahm.ProductResource(c)
}
