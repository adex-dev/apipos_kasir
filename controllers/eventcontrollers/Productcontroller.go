package eventcontrollers

import (
	"apipos/resources/eventt"
	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {
	return eventt.ProductResource(c)
}
