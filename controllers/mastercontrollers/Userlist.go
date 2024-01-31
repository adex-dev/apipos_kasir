package mastercontrollers

import (
	"apipos/database"
	"apipos/models"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func Userlist(c *fiber.Ctx) error {
	var user []models.Users
	rs, err := database.DBMASTER.Query("select nik,nama_user,add_ons,status,level,state from  users order by nama_user ASC ")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"messages": err.Error(), "respon": 501})
	}
	defer rs.Close()
	var state sql.NullString
	for rs.Next() {
		var us models.Users
		err := rs.Scan(&us.Nik, &us.Nama_user, &us.Add_ons, &us.Status, &us.Level, &state)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": "", "respon": 404})
		}
		if state.Valid {
			us.State = state.String
		} else {
			us.State = ""
		}

		user = append(user, us)
	}
	return c.JSON(fiber.Map{"data": user, "respon": 200})
}
