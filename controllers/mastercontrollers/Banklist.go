package mastercontrollers

import (
	"apipos/resources/Master"
	"github.com/gofiber/fiber/v2"
)

func Bank(c *fiber.Ctx) error {
	return Master.Bankresource(c)
}
