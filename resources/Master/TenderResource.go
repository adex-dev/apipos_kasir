package Master

import (
	"apipos/database"
	"apipos/models/master"
	"github.com/gofiber/fiber/v2"
)

func OtherTenderResource(c *fiber.Ctx) error {
	var tenders []master.Tendermodel
	query := `select pt.id_tender,pt.tender_name,pt.payment_type,pt.status,pt.location_state,pt.add_ons,pt.brands,pt.persen,pt.balance_nominal,pt.balance_minimal,pt.balance_maximal,pt.st_time,pt.st_date,user.nama_user from prm_balance_tender AS pt inner join users AS user on user.nik = pt.create_ons order by pt.tender_name Asc`
	result, err := database.DBMASTER.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer result.Close()
	for result.Next() {
		var td master.Tendermodel
		err := result.Scan(&td.IdTender, &td.TenderName, &td.PaymentType, &td.Status, &td.LocationState, &td.AddOns, &td.Brands, &td.Persen, &td.BalanceNominal, &td.BalanceMinimal, &td.BalanceMaximal, &td.StTime, &td.StDate, &td.NamaUser)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"data": err.Error(), "respon": 401})
		}
		tenders = append(tenders, td)
	}

	return c.JSON(fiber.Map{"data": tenders, "respon": 200})
}
