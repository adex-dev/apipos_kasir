package Master

import (
	"apipos/database"
	"github.com/gofiber/fiber/v2"
)

type paginat struct {
	Starts string `json:"starts"`
	Endded string `json:"endded"`
}

func Paginationresources(c *fiber.Ctx) error {
	var pagi []paginat
	rows, err := database.DBMASTER.Query("select * from pagination")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error internal Server",
			"respon":  401,
		})
	}
	defer rows.Close()
	for rows.Next() {
		var pg paginat
		errr := rows.Scan(&pg.Starts, &pg.Endded)
		if errr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error internal Server",
				"respon":  401,
			})
		}
		pagi = append(pagi, pg)
	}
	return c.JSON(fiber.Map{
		"data":    pagi,
		"message": "Success",
		"respon":  200,
	})
}
