package mastercontrollers

import (
	"apipos/resources/Master"
	"github.com/gofiber/fiber/v2"
)

func Wallet(c *fiber.Ctx) error {
	return Master.WalletResource(c)
}
