package nahm

import (
	"apipos/database"
	"apipos/models/nahm"
	"github.com/gofiber/fiber/v2"
)

func SubkategoriResource(c *fiber.Ctx) error {
	var subkategori []nahm.Subkategorilist
	var query string
	query = `SELECT subkategori.id_subitem,subkategori.nomor,subkategori.nama_subitem,subkategori.id_kategori,subkategori.add_ons,subkategori.status,subkategori.location_state,subkategori.infor,kategori.nama_kategori,users.nama_user from subkategori INNER JOIN kategori ON kategori.id_kategori = subkategori.id_kategori INNER JOIN pos_master.users ON users.nik = subkategori.create_user order by subkategori.id_subitem ASC`
	rows, err := database.DBNAHM.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"respon": 501,
			"data":   "",
		})
	}
	defer rows.Close()
	for rows.Next() {
		var st nahm.Subkategorilist
		err := rows.Scan(&st.Idsub, &st.Nomor, &st.Namasub, &st.Idkate, &st.Addons, &st.Status, &st.Locationstate, &st.Infor, &st.Namakategori, &st.Namauser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"respon": 501,
				"data":   "",
			})
		}
		subkategori = append(subkategori, st)
	}
	return c.JSON(fiber.Map{
		"respon": 200,
		"data":   subkategori,
	})
}
