package nahmcontrollers

import (
	"apipos/database"
	"apipos/models/isoide"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func Summarynahm(c *fiber.Ctx) error {
	var (
		startdate string = c.FormValue("starts")
		enddate   string = c.FormValue("enddate")
		idstore   string = c.FormValue("idstore")
		katesum   []isoide.KategoriSummary
	)
	storesummary, err := storesummary(idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"data":   "outlet",
			"respon": 501,
		})
	}
	paymentcardsummary, err := paymentcardsummary(startdate, enddate, idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"data":   "paymentsummary",
			"respon": 501,
		})
	}
	paymentothersummary, err := paymentothersummary(startdate, enddate, idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"data":   "paymentothersummary",
			"respon": 501,
		})
	}
	salesbreakdown, err := salesbreakdown(startdate, enddate, idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"data":   "Sales Breakdown",
			"respon": 501,
		})
	}
	avgbill, err := avgbill(startdate, enddate, idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"data":   "Avg Bill",
			"respon": 501,
		})
	}

	ktsum, errkate := database.DBNAHM.Query("SELECT DISTINCT kategori.nama_kategori,trm_item_order.terminal, SUM(trm_item_order.grandprice) as total from terminal_parameter as terminal inner join trm_item_order on trm_item_order.terminal = terminal.id_terminal inner join kategori on kategori.id_kategori = trm_item_order.kategori WHERE terminal.close_terminal >=?  AND terminal.close_terminal <=? AND terminal.status_terminal=? AND terminal.id_lokasi=? AND trm_item_order.status=? GROUP BY trm_item_order.kategori", startdate, enddate, "CLOSE", idstore, "PAID")
	if errkate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Data": errkate.Error(), "respon": 501})
	}
	defer ktsum.Close()
	for ktsum.Next() {
		var kts isoide.KategoriSummary
		var total float64
		if errkate := ktsum.Scan(&kts.Namakategori, &kts.Terminal, &total); errkate != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": "tidak ditemukan", "respon": 401})
		}
		kts.Total = int(total)
		katesum = append(katesum, kts)
	}
	return c.JSON(fiber.Map{
		"datastore":    storesummary,
		"datakategori": katesum,
		"datacard":     paymentcardsummary,
		"otherpayment": paymentothersummary,
		"breakdown":    salesbreakdown,
		"databill":     avgbill,
		"respon":       200,
	})
}
func storesummary(idstore string) ([]isoide.Storesummarymodel, error) {
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

func paymentcardsummary(startdate, enddate, idstore string) ([]isoide.PaymentSummarycardmodel, error) {
	var payment []isoide.PaymentSummarycardmodel
	rows, err := database.DBNAHM.Query("SELECT DISTINCT pd.jenis_payment, pd.type_payment, pd.payment_name, SUM(pd.nominal_payment) as total FROM terminal_parameter AS terminal INNER JOIN trm_payment_detail AS pd ON pd.terminal = terminal.id_terminal WHERE (terminal.close_terminal BETWEEN ? AND ?) AND terminal.status_terminal = ? AND terminal.id_lokasi = ? AND pd.status = ? AND (pd.type_payment = 'DEBIT' OR pd.type_payment = 'CREDIT' OR pd.type_payment = 'CASH') GROUP BY pd.type_payment, pd.payment_name, pd.jenis_payment", startdate, enddate, "CLOSE", idstore, "PAID")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pt isoide.PaymentSummarycardmodel
		var total float64
		err := rows.Scan(&pt.JenisPayment, &pt.TypePayment, &pt.PaymentName, &total)
		pt.Total = int(total)
		if err != nil {
			return nil, err
		}
		payment = append(payment, pt)
	}
	return payment, nil
}
func paymentothersummary(startdate, enddate, idstore string) ([]isoide.PaymentSummarycardmodel, error) {
	var payment []isoide.PaymentSummarycardmodel
	rows, err := database.DBNAHM.Query("SELECT DISTINCT pd.jenis_payment, pd.type_payment, pd.payment_name, SUM(pd.nominal_payment) as total FROM terminal_parameter AS terminal INNER JOIN trm_payment_detail AS pd ON pd.terminal = terminal.id_terminal WHERE (terminal.close_terminal BETWEEN ? AND ?) AND terminal.status_terminal = ? AND terminal.id_lokasi = ? AND pd.status = ? AND pd.type_payment NOT IN  ('DEBIT', 'CREDIT', 'CASH') GROUP BY pd.type_payment, pd.payment_name, pd.jenis_payment", startdate, enddate, "CLOSE", idstore, "PAID")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pt isoide.PaymentSummarycardmodel
		var total float64
		err := rows.Scan(&pt.JenisPayment, &pt.TypePayment, &pt.PaymentName, &total)
		pt.Total = int(total)
		if err != nil {
			return nil, err
		}
		payment = append(payment, pt)
	}
	return payment, nil
}

func salesbreakdown(startdate, enddate, idstore string) (map[string]float64, error) {
	rows, err := database.DBNAHM.Query("SELECT SUM(pt.changed) as changed,SUM(pt.gross_sales) as grosssales,SUM(pt.disc_item) as disitem,SUM(pt.disc_other) as disc_other,SUM(pt.tax) as tax,SUM(pt.service_charge) as service_charge,SUM(pt.net_sales) as total_payment,SUM(pt.change_invisible) as change_invisible from trm_payment as pt INNER JOIN terminal_parameter as terminal on terminal.id_terminal = pt.terminal WHERE (terminal.close_terminal BETWEEN ? AND ?) AND terminal.status_terminal='CLOSE' AND terminal.id_lokasi=? AND pt.status_payment='PAID' GROUP BY pt.status_payment", startdate, enddate, idstore)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var paymentSummary = make(map[string]float64)
	if rows.Next() {
		var (
			changed         float64
			grossSales      float64
			discountItem    float64
			discountOther   float64
			tax             float64
			serviceCharge   float64
			totalPayment    float64
			changeInvisible float64
		)
		err := rows.Scan(&changed, &grossSales, &discountItem, &discountOther, &tax, &serviceCharge, &totalPayment, &changeInvisible)
		if err != nil {
			return nil, err
		}

		paymentSummary["changed"] = changed
		paymentSummary["grosssales"] = grossSales
		paymentSummary["disitem"] = discountItem
		paymentSummary["disc_other"] = discountOther
		paymentSummary["tax"] = tax
		paymentSummary["service_charge"] = serviceCharge
		paymentSummary["total_payment"] = totalPayment
		paymentSummary["change_invisible"] = changeInvisible
	}
	return paymentSummary, nil
}

func avgbill(startdate, enddate, idstore string) (map[string]float64, error) {
	rows, err := database.DB.Query("SELECT COUNT(no_trans) as totalavg,SUM(custumer_count) as avgcover FROM transaction_parameter WHERE (close_transaksi between ? AND ?) AND status_transaksi='CLOSED' AND id_lokasi=?", startdate, enddate, idstore)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var avgbills = make(map[string]float64)
	if rows.Next() {
		var totalAv sql.NullFloat64
		var avCover sql.NullFloat64
		err := rows.Scan(&totalAv, &avCover)
		if err != nil {
			return nil, err
		}

		// Menggunakan nilai default jika kolom memiliki nilai NULL
		if totalAv.Valid {
			avgbills["totalavg"] = totalAv.Float64
		} else {
			avgbills["totalavg"] = 0 // Nilai default jika NULL
		}

		if avCover.Valid {
			avgbills["avgcover"] = avCover.Float64
		} else {
			avgbills["avgcover"] = 0 // Nilai default jika NULL
		}
	}
	return avgbills, nil
}
