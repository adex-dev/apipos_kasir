package mastercontrollers

import (
	"apipos/resources/Master"
	"github.com/gofiber/fiber/v2"
)

func OtherTender(c *fiber.Ctx) error {
	return Master.OtherTenderResource(c)
}
