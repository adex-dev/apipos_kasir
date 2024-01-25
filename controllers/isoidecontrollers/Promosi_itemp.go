package isoidecontrollers

import (
	"apipos/database"
	"apipos/models/isoide"
	"github.com/gofiber/fiber/v2"
)

func Promosi_itemp(c *fiber.Ctx) error {
	id := c.FormValue("id")
	var promositemp []isoide.Ptempmodel
	rs, err := database.DB.Query(`SELECT p.*,pt.nameP,pl.itemName  from  prm_discount_i_temp as p inner join prm_discount_i as pt on pt.id_promotion = p.idpromotion inner join plu_items as pl on pl.id_item=p.kd_product  where idpromotion=?`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "respon": 501})
	}
	defer rs.Close()
	for rs.Next() {
		var prm isoide.Ptempmodel
		err := rs.Scan(&prm.Idpromotion, &prm.KdProduct, &prm.StDate, &prm.EdDate, &prm.Namep, &prm.ItemName)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": "data tidak ditemukan", "respon": 404})
		}
		promositemp = append(promositemp, prm)
	}
	return c.JSON(fiber.Map{"data": promositemp, "respon": 200})
}
