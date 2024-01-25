package mastercontrollers

import (
	"apipos/resources/Master"
	"github.com/gofiber/fiber/v2"
)

func Cashlist(c *fiber.Ctx) error {
	return Master.CashResource(c)
}
