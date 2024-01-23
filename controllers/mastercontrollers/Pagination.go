package mastercontrollers

import (
	"apipos/resources/Master"
	"github.com/gofiber/fiber/v2"
)

func Pagination(c *fiber.Ctx) error {
	return Master.Paginationresources(c)
}
