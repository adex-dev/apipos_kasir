package isoidecontrollers

import (
	"apipos/database"
	"apipos/models/isoide"
	"github.com/gofiber/fiber/v2"
)

func Jurnalpluisoide(c *fiber.Ctx) error {
	var (
		startdate string = c.FormValue("starts")
		enddate   string = c.FormValue("enddate")
		idstore   string = c.FormValue("idstore")
	)

	dt, err := jurnalPLUisoide(startdate, enddate, idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"data":   "Plu",
			"respon": 501,
		})
	}
	var response []map[string]interface{}

	for _, vl := range dt {
		variables, err := Dataprinterisoide(vl.NoTrans)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err,
				"data":   "dataprinter",
				"respon": 501,
			})
		}
		for _, value := range variables {
			response = append(response, map[string]interface{}{
				"price":         value.Price,
				"qty":           value.Qty,
				"grandprice":    value.Price * value.Qty,
				"no_trans":      vl.NoTrans,
				"id_product":    vl.IdProduct,
				"itemName":      vl.ItemNama,
				"nama_kategori": vl.NamaKategori,
				"nama_subitem":  vl.NamaSubitem,
				"additional":    value.Additional,
				"status":        vl.Status,
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   response,
		"respon": 200,
	})

}
func jurnalPLUisoide(startdate, enddate, idstore string) ([]isoide.Plusummary, error) {
	var Result []isoide.Plusummary
	rows, err := database.DB.Query("SELECT tp.no_trans,ti.id_product,pl.itemName,kt.nama_kategori,sb.nama_subitem,tp.status_transaksi from transaction_parameter as tp inner join terminal_parameter as terminal on terminal.id_terminal = tp.id_terminal inner join trm_item_order as ti on ti.no_trans = tp.no_trans inner join plu_items as pl on pl.id_item = ti.id_product inner join subkategori as sb on sb.id_subitem = pl.itemsub inner join kategori as kt on kt.id_kategori = pl.itemKategory where (terminal.close_terminal BETWEEN  ? AND ?) AND terminal.id_lokasi=? AND terminal.status_terminal='CLOSE' order by tp.no_trans", startdate, enddate, idstore)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rs isoide.Plusummary
		err := rows.Scan(&rs.NoTrans, &rs.IdProduct, &rs.ItemNama, &rs.NamaKategori, &rs.NamaSubitem, &rs.Status)
		if err != nil {
			return nil, err
		}
		Result = append(Result, rs)
	}

	return Result, nil
}
