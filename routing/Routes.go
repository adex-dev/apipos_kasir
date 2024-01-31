package routing

import (
	"apipos/controllers/authcontrollers"
	"apipos/controllers/eventcontrollers"
	"apipos/controllers/isoidecontrollers"
	"apipos/controllers/marketingcontroller"
	"apipos/controllers/mastercontrollers"
	"apipos/controllers/nahmcontrollers"
	"github.com/gofiber/fiber/v2"
)

func Routes(c *fiber.App) {
	app := c.Group("api")
	//auth
	app.Post("cekuser", authcontrollers.Loginauth)
	app.Post("loadsales", mastercontrollers.Loadsales)
	//master
	app.Get("pagination", mastercontrollers.Pagination)
	app.Get("/outlet", mastercontrollers.Storeshow)
	app.Get("/banklist", mastercontrollers.Bank)
	app.Get("/wallet", mastercontrollers.Wallet)
	app.Get("/other-tender", mastercontrollers.OtherTender)
	app.Get("/cashlist", mastercontrollers.Cashlist)
	app.Get("/getuser", mastercontrollers.Userlist)
	app.Get("/target-sales", mastercontrollers.TargetSales)
	//isoide
	isoide := app.Group("/isoide")
	isoide.Get("/product", isoidecontrollers.Product)
	isoide.Get("/kategori/:status?", isoidecontrollers.Kategori)
	isoide.Get("/subkategori/:idsub?/:idkategori?", isoidecontrollers.Subkategori)
	isoide.Get("/promotionh", isoidecontrollers.PromotionH)
	isoide.Get("/promotion_i", isoidecontrollers.PromotionI)
	isoide.Get("/voucher", isoidecontrollers.Voucher)
	isoide.Post("/disc-item-temp", isoidecontrollers.Promosi_itemp)
	isoide.Post("/summary", isoidecontrollers.Summaryisoide)
	isoide.Post("/salessummaryplu", isoidecontrollers.Summarypluisoide)
	isoide.Post("/salesreport", isoidecontrollers.Salesreportisoide)
	isoide.Post("/salespluisoide", isoidecontrollers.Salessummarypluisoide)
	isoide.Post("/jurnalpluisoide", isoidecontrollers.Jurnalpluisoide)

	//	Nahm
	nahm := app.Group("/nahm")
	nahm.Get("/product", nahmcontrollers.Product)
	nahm.Get("/kategori", nahmcontrollers.Kategori)
	nahm.Get("/subkategori", nahmcontrollers.Subkategori)
	nahm.Get("/promotionh", nahmcontrollers.Promotionh)
	nahm.Get("/promotion_i", nahmcontrollers.Promotioni)
	nahm.Get("/voucher", nahmcontrollers.Voucher)
	nahm.Post("/disc-item-temp", nahmcontrollers.Promosiitempp)
	nahm.Post("/summary", nahmcontrollers.Summarynahm)
	nahm.Post("/salessummaryplu", nahmcontrollers.Summaryplunahms)
	nahm.Post("/salesplunahm", nahmcontrollers.Reportplunahm)
	nahm.Post("/salesreport", nahmcontrollers.Salesreportnahm)
	nahm.Post("/jurnalplu", nahmcontrollers.Jurnalplunahm)
	//	Event
	event := app.Group("/event")
	event.Get("/product", eventcontrollers.Product)
	event.Get("/kategori", eventcontrollers.Kategori)
	event.Get("/subkategori", eventcontrollers.Subkategori)
	event.Get("/promotionh", eventcontrollers.Promotionhevent)
	event.Get("/promotion_i", eventcontrollers.Promotionievent)
	event.Get("/voucher", eventcontrollers.Vouchereventt)
	event.Post("/disc-item-temp", eventcontrollers.Promosiitemppevent)

	//Marketing
	marketing := app.Group("/marketing")
	marketing.Get("/countmember", marketingcontroller.Countmember)
	marketing.Get("/getpoint", marketingcontroller.Poinlist)
	marketing.Get("/voucher-reg-reedem", marketingcontroller.Voucherredem)
	marketing.Get("/voucher", marketingcontroller.Voucherget)
	marketing.Get("/voucher-ulta", marketingcontroller.Voucherultah)
	marketing.Get("/members", marketingcontroller.Members)

}
