package Master

import (
	"apipos/database"
	"apipos/models/master"
	"github.com/gofiber/fiber/v2"
)

func CashResource(c *fiber.Ctx) error {
	var Cashs []master.Cashmodel
	query := `SELECT c.id_balance,c.balance_name,c.payment_type,c.status,c.location_state,c.add_ons,user.nama_user from prm_balance_cash as c inner join users as user on user.nik = c.create_user order by c.balance_name ASC`
	Rs, err := database.DBMASTER.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": "", "respon": 500})
	}
	defer Rs.Close()
	for Rs.Next() {
		var cs master.Cashmodel
		err := Rs.Scan(&cs.IdBalance, &cs.BalanceName, &cs.PaymentType, &cs.Status, &cs.LocationState, &cs.AddOns, &cs.NamaUser)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": "", "respon": 404})
		}
		Cashs = append(Cashs, cs)
	}
	return c.JSON(fiber.Map{"data": Cashs, "respon": 200})
}
