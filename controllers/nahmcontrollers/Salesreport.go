package nahmcontrollers

import (
	"apipos/database"
	"apipos/models/isoide"
	"github.com/gofiber/fiber/v2"
)

func Salesreportnahm(c *fiber.Ctx) error {
	var (
		startdate string = c.FormValue("starts")
		enddate   string = c.FormValue("enddate")
		idstore   string = c.FormValue("idstore")
		datanilai []isoide.DataNilai
	)
	rows, err := database.DBNAHM.Query("SELECT tp.no_trans,tp.gross_sales,tp.disc_item,tp.disc_other,tp.tax,tp.service_charge,tp.total_payment,tp.net_sales,tp.changed,tp.change_invisible,tp.status_payment,tp.close_payment,td.type_payment,td.payment_name,td.nominal_payment,td.jenis_payment from  trm_payment as tp inner join trm_payment_detail as td on td.no_trans = tp.no_trans inner join terminal_parameter as terminal on terminal.id_terminal = tp.terminal where (terminal.close_terminal between ? AND ? ) AND terminal.id_lokasi=? AND terminal.status_terminal='CLOSE' order by tp.no_trans ASC ", startdate, enddate, idstore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer rows.Close()
	for rows.Next() {
		var dt isoide.DataNilai
		err := rows.Scan(&dt.NoTrans, &dt.Gross, &dt.DiscItem, &dt.DiscOther, &dt.Tax, &dt.ServiceCharge, &dt.TotalPayment, &dt.NetSales, &dt.Changed, &dt.ChangeInvisible, &dt.StatusPayment, &dt.ClosePayment, &dt.TypePayment, &dt.PaymentName, &dt.NominalPayment, &dt.JenisPayment)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
		}
		datanilai = append(datanilai, dt)
	}

	var databaru = make(map[string]isoide.DataNilai)

	for _, vl := range datanilai {
		notrans := vl.NoTrans
		total := vl.NominalPayment
		gross := vl.Gross
		discother := vl.DiscOther
		discitem := vl.DiscItem
		serviceecharge := vl.ServiceCharge
		tax := vl.Tax
		netsales := vl.NetSales
		changed := vl.Changed
		changeinvisible := vl.ChangeInvisible
		jenispayment := vl.JenisPayment
		typepayment := vl.TypePayment
		paymentname := vl.PaymentName
		closepayment := vl.ClosePayment

		// Buat kunci unik untuk setiap entri dalam map
		key := notrans + "-" + paymentname

		// Periksa apakah entri sudah ada dalam map
		if _, ok := databaru[key]; !ok {
			// Buat entri baru dalam map
			databaru[key] = isoide.DataNilai{
				NominalPayment:  total,
				NoTrans:         notrans,
				ClosePayment:    closepayment,
				Gross:           gross,
				DiscOther:       discother,
				DiscItem:        discitem,
				ServiceCharge:   serviceecharge,
				Tax:             tax,
				JenisPayment:    jenispayment,
				TypePayment:     typepayment,
				PaymentName:     paymentname,
				NetSales:        netsales,
				Changed:         changed,
				ChangeInvisible: changeinvisible,
			}
		} else {
			existingData := databaru[key]
			existingData.NominalPayment += total
			databaru[key] = existingData
		}
	}

	// Konversi map menjadi slice
	var dataArray []map[string]isoide.DataNilai
	for key, value := range databaru {
		entry := make(map[string]isoide.DataNilai)
		entry[key] = value
		dataArray = append(dataArray, entry)
	}

	return c.JSON(fiber.Map{
		"data":   dataArray,
		"respon": 200,
	})

}
