package isoidecontrollers

import (
	"apipos/resources/isoide"
	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {
	return isoide.ProductResource(c)
}
