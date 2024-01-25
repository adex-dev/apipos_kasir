package isoidecontrollers

import (
	"apipos/database"
	"apipos/models/isoide"
	"github.com/gofiber/fiber/v2"
)

func PromotionH(c *fiber.Ctx) error {
	var promosi []isoide.PromosiHlist
	rs, err := database.DB.Query(`SELECT p.id_promotion,p.type,p.nameP,p.discPersen,p.discValue,p.Minimal,p.Maximal,p.stdate,p.enddate,p.state,p.sttime,p.entime,p.outlet,p.addOn,p.member,p.status,user.nama_user from prm_discount_h as p inner join pos_master.users as user on user.nik = p.create_ons order by p.id_promotion DESC`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer rs.Close()
	for rs.Next() {
		var prm isoide.PromosiHlist
		err := rs.Scan(&prm.IdPromotion, &prm.Type, &prm.Namep, &prm.DiscPersen, &prm.DiscValue, &prm.Minimal, &prm.Maximal, &prm.Stdate, &prm.Enddate, &prm.State, &prm.Sttime, &prm.Endtime, &prm.Outlet, &prm.AddOns, &prm.Member, &prm.Status, &prm.NamaUser)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": "tidak tesediah", "respon": 404})
		}
		promosi = append(promosi, prm)
	}
	return c.JSON(fiber.Map{"data": promosi, "respon": 200})
}
