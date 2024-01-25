package routing

import (
	"apipos/controllers/authcontrollers"
	"apipos/controllers/eventcontrollers"
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
	app.Get("/banklist", mastercontrollers.Bank)
	app.Get("/wallet", mastercontrollers.Wallet)
	app.Get("/other-tender", mastercontrollers.OtherTender)
	app.Get("/cashlist", mastercontrollers.Cashlist)
	//isoide
	isoide := app.Group("/isoide")
	isoide.Get("/product", isoidecontrollers.Product)
	isoide.Get("/kategori/:status?", isoidecontrollers.Kategori)
	isoide.Get("/subkategori/:idsub?/:idkategori?", isoidecontrollers.Subkategori)
	isoide.Get("/promotionh", isoidecontrollers.PromotionH)
	isoide.Get("/promotion_i", isoidecontrollers.PromotionI)
	isoide.Get("/voucher", isoidecontrollers.Voucher)
	isoide.Post("/disc-item-temp", isoidecontrollers.Promosi_itemp)

	//	Nahm
	nahm := app.Group("/nahm")
	nahm.Get("/product", nahmcontrollers.Product)
	nahm.Get("/kategori", nahmcontrollers.Kategori)
	nahm.Get("/subkategori", nahmcontrollers.Subkategori)

	//	Event
	event := app.Group("/event")
	event.Get("/product", eventcontrollers.Product)
	event.Get("/kategori", eventcontrollers.Kategori)
	event.Get("/subkategori", eventcontrollers.Subkategori)
}
