package Master

import (
	"apipos/database"
	"apipos/models/master"
	"github.com/gofiber/fiber/v2"
)

func Bankresource(c *fiber.Ctx) error {
	var banks []master.Bankmodel
	query := `SELECT prm_balance.*,users.nama_user from prm_balance inner join users on users.nik = prm_balance.create_user order by balance_name asc`
	rows, err := database.DBMASTER.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"respon": 501,
			"data":   "",
		})
	}
	defer rows.Close()
	for rows.Next() {
		var bk master.Bankmodel
		err := rows.Scan(&bk.Idbalance, &bk.Nomor, &bk.BalanceName, &bk.PaymentType, &bk.Status, &bk.LocationState, &bk.AddOns, &bk.Brands, &bk.StatusUpdate, &bk.CreateUser, &bk.Nama_user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"respon": 501,
				"data":   "",
			})
		}
		banks = append(banks, bk)
	}
	return c.JSON(fiber.Map{
		"data":   banks,
		"respon": 200,
	})
}
