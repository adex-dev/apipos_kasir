package marketingcontroller

import (
	"apipos/database"
	"apipos/models/marketing"
	"github.com/gofiber/fiber/v2"
)

func Members(c *fiber.Ctx) error {
	var member []marketing.Membermodel
	result, err := database.DBMEMBER.Query("SELECT idmember,nama,email,gender,no_hp,birthday,joindate,status from prm_member order by  nama ASC ")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer result.Close()

	for result.Next() {
		var mb marketing.Membermodel
		err := result.Scan(&mb.Idmember, &mb.Nama, &mb.Email, &mb.Gender, &mb.NoHp, &mb.Birthday, &mb.Joindate, &mb.Status)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
		}

		member = append(member, mb)
	}
	return c.JSON(fiber.Map{"data": member, "respon": 200})
}
