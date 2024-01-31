package mastercontrollers

import (
	"apipos/database"
	"apipos/models/master"
	"github.com/gofiber/fiber/v2"
)

func TargetSales(c *fiber.Ctx) error {
	var target []master.TargetModel
	rows, err := database.DBMASTER.Query("SELECT tr.*, fk.nama_outlet FROM target_sales AS tr INNER JOIN prm_outlet AS fk ON fk.id_outlet = tr.store_id ORDER BY tr.date DESC")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer rows.Close()
	for rows.Next() {
		var tg master.TargetModel
		err := rows.Scan(&tg.Date, &tg.Amount, &tg.StoreId, &tg.NamaOutlet)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": err.Error(), "respon": 404})
		}
		target = append(target, tg)
	}
	return c.JSON(fiber.Map{"data": target, "respon": 200})
}
