package marketingcontroller

import (
	"apipos/database"
	"apipos/models/marketing"
	"github.com/gofiber/fiber/v2"
)

func Voucherredem(c *fiber.Ctx) error {
	var voucher []marketing.Voucherredemmodel
	result, err := database.DBMEMBER.Query("SELECT vouchercode,vouchername,tanggal,claim_date,claim_store,status,persentase,valuenominal from prm_reedem_voucher_temp where type_reedem='REGULER' AND status='COMPLETED' order by tanggal DESC")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer result.Close()
	for result.Next() {
		var vc marketing.Voucherredemmodel
		err := result.Scan(&vc.Vouchercode, &vc.Vouchername, &vc.Tanggal, &vc.Claimdate, &vc.Claimstore, &vc.Status, &vc.Persentase, &vc.Valuenominal)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": err, "respon": 201})
		}
		voucher = append(voucher, vc)
	}
	return c.JSON(fiber.Map{"data": voucher, "respon": 200})
}

func Voucherget(c *fiber.Ctx) error {
	var voucher []marketing.Vouchermarketingmodel
	result, err := database.DBMEMBER.Query("SELECT vouchercode,namevoucher,qty,expired,disvalue,dispersent from prm_voucher order by expired DESC")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer result.Close()
	for result.Next() {
		var vc marketing.Vouchermarketingmodel
		err := result.Scan(&vc.Vouchercode, &vc.Vouchername, &vc.Qty, &vc.Expired, &vc.Disvalue, &vc.Dispersent)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": err.Error(), "respon": 201})
		}
		voucher = append(voucher, vc)
	}
	return c.JSON(fiber.Map{"data": voucher, "respon": 200})

}

func Voucherultah(c *fiber.Ctx) error {
	var voucherulta []marketing.Voucherultahmodel
	total, err := CountClaimvoucherulta()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	results, err := database.DBMEMBER.Query("SELECT vc.vouchercode,vc.birthday,vc.validvoucher,vc.status,mb.nama,mb.gender,mb.email,mb.no_hp,mb.joindate from prm_voucher_birtday as vc inner join prm_member as mb on mb.idmember = vc.memberid group by vc.memberid order by vc.birthday DESC")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "respon": 501})
	}
	defer results.Close()
	for results.Next() {
		var vcr marketing.Voucherultahmodel
		err := results.Scan(&vcr.Vouchercode, &vcr.Birthday, &vcr.Validvoucher, &vcr.Status, &vcr.Nama, &vcr.Gender, &vcr.Email, &vcr.Nohps, &vcr.Joindate)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": err.Error(), "respon": 201})
		}
		voucherulta = append(voucherulta, vcr)
	}
	return c.JSON(fiber.Map{
		"data":   voucherulta,
		"total":  total,
		"respon": 200})
}

func CountClaimvoucherulta() (float64, error) {
	var total float64

	err := database.DBMEMBER.QueryRow("SELECT COUNT(vouchercode) as totalclaim from prm_voucher_birtday WHERE status=?", "CLAIM").Scan(&total)
	if err != nil {
		return 0, nil
	}
	return total, nil
}
