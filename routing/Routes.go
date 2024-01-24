package routing

import (
	"apipos/controllers/authcontrollers"
	"apipos/controllers/isoidecontrollers"
	"apipos/controllers/mastercontrollers"
	"apipos/controllers/nahmcontrollers"
	"github.com/gofiber/fiber/v2"
)

func Routes(c *fiber.App) {
	app := c.Group("api")
	//auth
	app.Post("auth/cekuser", authcontrollers.Loginauth)
	app.Post("loadsales", mastercontrollers.Loadsales)
	//master
	app.Get("pagination", mastercontrollers.Pagination)
	app.Get("/outlet", mastercontrollers.Storeshow)
	//isoide
	isoide := app.Group("/isoide")
	isoide.Get("/product", isoidecontrollers.Product)
	isoide.Get("/kategori/:status?", isoidecontrollers.Kategori)
	isoide.Get("/subkategori/:idsub?/:idkategori?", isoidecontrollers.Subkategori)

	//	Nahm
	nahm := app.Group("/nahm")
	nahm.Get("/product", nahmcontrollers.Product)
	nahm.Get("/kategori", nahmcontrollers.Kategori)
	nahm.Get("/subkategori", nahmcontrollers.Subkategori)

}
