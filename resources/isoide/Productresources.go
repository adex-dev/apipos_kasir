package isoide

import (
	"apipos/database"
	"apipos/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ProductResource(c *fiber.Ctx) error {
	var products []models.Productlist
	query := `
		SELECT plu_items.*, kategori.nama_kategori, subkategori.nama_subitem, users.nama_user
		FROM plu_items
		INNER JOIN kategori ON kategori.id_kategori = plu_items.itemKategory
		INNER JOIN subkategori ON subkategori.id_subitem = plu_items.itemSub
		INNER JOIN pos_master.users ON users.nik = plu_items.create_ons order by plu_items.id_item DESC 
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"respon": 401,
		})
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Productlist
		err := rows.Scan(
			&product.Id_item,
			&product.Nomor_items,
			&product.ItemName,
			&product.ItemKategori,
			&product.ItemSub,
			&product.ItemPrice,
			&product.ItemOptional,
			&product.ItemStatus,
			&product.Adds_date,
			&product.Adds_time,
			&product.State,
			&product.Stdate,
			&product.Endate,
			&product.Dist_state,
			&product.Infor,
			&product.Create_ons,
			&product.NamaKategori,
			&product.NamaSubitem,
			&product.Nama_user,
		)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"respon": 401,
			})
		}

		products = append(products, product)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":   products,
		"respon": 200,
	})
}
