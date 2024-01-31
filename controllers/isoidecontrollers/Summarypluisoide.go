package isoidecontrollers

import (
	"apipos/database"
	"apipos/models/isoide"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func Summarypluisoide(c *fiber.Ctx) error {
	var (
		startdate string = c.FormValue("starts")
		enddate   string = c.FormValue("enddate")
		idstore   string = c.FormValue("idstore")
	)
	storesummarys, err := storesummarys(idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"data":   "outlet",
			"respon": 501,
		})
	}

	dt, err := SalerReportSummaryPLUisoide(startdate, enddate, idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"data":   "Plu",
			"respon": 501,
		})
	}
	var response []map[string]interface{}

	// Iterasi melalui hasil laporan penjualan
	for _, vl := range dt {
		variables, err := AmbilDetailPrinter(vl.NoTrans, vl.IdProduct)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err,
				"data":   "dataprinter",
				"respon": 501,
			})
		}

		// Iterasi melalui detail printer
		for _, value := range variables {
			response = append(response, map[string]interface{}{
				"price":         value.Price,
				"qty":           value.Qty,
				"grandprice":    value.Price * value.Qty,
				"id_store":      idstore,
				"nama_store":    storesummarys[0].NamaStore,
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
func storesummarys(idstore string) ([]isoide.Storesummarymodel, error) {
	var store []isoide.Storesummarymodel
	st, err := database.DBMASTER.Query("SELECT nama_outlet, telp_1, telp_2, alamat, service_charge, tax FROM prm_outlet WHERE id_outlet=?", idstore)
	if err != nil {
		return nil, err
	}
	defer st.Close()
	for st.Next() {
		var sts isoide.Storesummarymodel
		err := st.Scan(&sts.NamaStore, &sts.Telp1, &sts.Telp2, &sts.Alamat, &sts.ServiceCharge, &sts.Tax)
		if err != nil {
			return nil, err
		}
		store = append(store, sts)
	}
	return store, nil
}

func SalerReportSummaryPLUisoide(startdate, enddate, idstore string) ([]isoide.Plusummary, error) {
	var Result []isoide.Plusummary
	rows, err := database.DB.Query("SELECT tp.no_trans,tp.close_transaksi,ti.id_product,tp.time_transaksi,pl.itemName,kt.nama_kategori,sb.nama_subitem,tp.status_transaksi from transaction_parameter as tp inner join terminal_parameter as terminal on terminal.id_terminal = tp.id_terminal inner join trm_item_order as ti on ti.no_trans = tp.no_trans inner join plu_items as pl on pl.id_item = ti.id_product inner join subkategori as sb on sb.id_subitem = pl.itemsub inner join kategori as kt on kt.id_kategori = pl.itemKategory where (terminal.close_terminal BETWEEN  ? AND ?) AND terminal.id_lokasi=? AND terminal.status_terminal='CLOSE' AND tp.status_transaksi='CLOSED' order by tp.no_trans", startdate, enddate, idstore)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rs isoide.Plusummary
		err := rows.Scan(&rs.NoTrans, &rs.CloseTransaksi, &rs.IdProduct, &rs.TimeTransaksi, &rs.ItemNama, &rs.NamaKategori, &rs.NamaSubitem, &rs.Status)
		if err != nil {
			return nil, err
		}
		Result = append(Result, rs)
	}

	return Result, nil
}
func AmbilDetailPrinter(NoTrans, IdProduct string) ([]isoide.Dataprinter, error) {
	var dta []isoide.Dataprinter
	rows, err := database.DB.Query("SELECT price,SUM(qty),additional from printing_order where no_trans=? AND id_product=?", NoTrans, IdProduct)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var additional sql.NullString
	for rows.Next() {
		var dt isoide.Dataprinter
		err := rows.Scan(&dt.Price, &dt.Qty, &additional)
		if err != nil {
			return nil, err
		}
		var _ string
		if additional.Valid {
			_ = dt.Additional
		} else {
			_ = ""
		}
		dta = append(dta, dt)
	}
	return dta, nil
}
