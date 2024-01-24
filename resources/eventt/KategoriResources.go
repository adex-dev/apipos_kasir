package eventt

import (
	"apipos/database"
	"apipos/models/nahm"
	"github.com/gofiber/fiber/v2"
)

func KategoriResource(c *fiber.Ctx) error {
	var kategori []nahm.Kategorilist

	rows, err := database.DBEVENT.Query(`select * from kategori order by nama_kategori ASC`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"respon": 501,
			"data":   "",
		})
	}
	defer rows.Close()
	for rows.Next() {
		var kt nahm.Kategorilist
		err := rows.Scan(&kt.Idkategori, &kt.Nomer, &kt.Namakategori, &kt.Addons, &kt.Status)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"respon": 501,
				"data":   "",
			})
		}
		kategori = append(kategori, kt)
	}
	return c.JSON(fiber.Map{
		"respon": 200,
		"data":   kategori,
	})
}
