package marketingcontroller

import (
	"apipos/database"
	"apipos/models/marketing"
	"github.com/gofiber/fiber/v2"
)

func Poinlist(c *fiber.Ctx) error {
	var poin []marketing.Poinmodal
	result, err := database.DBMEMBER.Query("SELECT poinid,poin,minpayment,maxpayment,addon,status from prm_poin order by poin ASC ")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer result.Close()

	for result.Next() {
		var pt marketing.Poinmodal
		err := result.Scan(&pt.Poinid, &pt.Poin, &pt.Minpayment, &pt.Maxpayment, &pt.Addons, &pt.Status)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
		}
		poin = append(poin, pt)
	}

	return c.JSON(fiber.Map{"data": poin, "respon": 200})
}
