package isoide

import (
	"apipos/database"
	"apipos/models/isoide"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func KategoriResource(status string, c *fiber.Ctx) error {
	var kategoris []isoide.Kategorilist
	var query string
	var rows *sql.Rows
	var err error

	if status == "" {
		query = `select * from kategori order by nomer ASC`
		rows, err = database.DB.Query(query)
	} else {
		query = `select * from kategori where status=? order by nomer ASC`
		rows, err = database.DB.Query(query, status)
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"respon": 501,
			"data":   "",
		})
	}
	defer rows.Close()
	for rows.Next() {
		var kt isoide.Kategorilist
		err := rows.Scan(&kt.Idkategori, &kt.Nomer, &kt.Namakategori, &kt.Addons, &kt.Status)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"respon": 501,
				"data":   "",
			})
		}
		kategoris = append(kategoris, kt)
	}
	return c.JSON(fiber.Map{
		"respon": 200,
		"data":   kategoris,
	})
}
