package marketingcontroller

import (
	"apipos/database"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Countmember(c *fiber.Ctx) error {
	lw, err := lastweek()
	cr, err := currentweek()
	th, err := totaltahun()
	cm, err := currentmoon()
	lm, err := lastmoon()
	count, err := allmembercount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"respon": 501,
		})
	}

	return c.JSON(fiber.Map{
		"respon": 200,
		"LW":     lw,
		"CR":     cr,
		"TH":     th,
		"CM":     cm,
		"LM":     lm,
		"count":  count,
	})
}

func lastweek() (float64, error) {
	query := `SELECT COUNT(idmember) as LW FROM prm_member WHERE joindate >= CURRENT_DATE - INTERVAL 1 WEEK AND joindate < CURRENT_DATE`

	var lw float64
	err := database.DBMEMBER.QueryRow(query).Scan(&lw)
	if err != nil {
		return 0, err
	}
	return lw, nil
}
func allmembercount() (float64, error) {
	query := `SELECT COUNT(idmember) as ct FROM prm_member`
	var ct float64
	err := database.DBMEMBER.QueryRow(query).Scan(&ct)
	if err != nil {
		return 0, err
	}
	return ct, nil
}
func currentweek() (float64, error) {
	var cr float64
	err := database.DBMEMBER.QueryRow("SELECT COUNT(idmember) as cr FROM prm_member where joindate >= CURRENT_DATE - INTERVAL 6 DAY AND joindate <= CURRENT_DATE").Scan(&cr)
	if err != nil {
		return 0, nil
	}
	return cr, nil
}
func lastmoon() (float64, error) {
	var cr float64
	currentTime := time.Now()

	// Mengurangi satu bulan dari tanggal saat ini
	lastMonth := currentTime.AddDate(0, -1, 0)

	// Mendapatkan awal bulan dari bulan lalu
	lastMonthStart := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, lastMonth.Location())

	// Mendapatkan akhir bulan dari bulan lalu
	lastMonthEnd := lastMonthStart.AddDate(0, 1, 0).Add(-time.Second)

	// Format tanggal menjadi string dalam format 'YYYY-MM-DD'
	formattedStart := lastMonthStart.Format("2006-01-02")
	formattedEnd := lastMonthEnd.Format("2006-01-02")
	err := database.DBMEMBER.QueryRow("SELECT COUNT(idmember) as cm FROM prm_member where joindate >=? AND joindate <=?", formattedStart, formattedEnd).Scan(&cr)
	if err != nil {
		return 0, nil
	}
	return cr, nil
}
func currentmoon() (float64, error) {
	var lm float64
	//tanggal saat ini
	currentTime := time.Now()

	// Mendapatkan awal bulan saat ini
	currentMonthStart := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())

	// Mendapatkan akhir bulan saat ini
	currentMonthEnd := currentMonthStart.AddDate(0, 1, 0).Add(-time.Second)

	// Format tanggal menjadi string dalam format 'YYYY-MM-DD'
	formattedStart := currentMonthStart.Format("2006-01-02")
	formattedEnd := currentMonthEnd.Format("2006-01-02")

	err := database.DBMEMBER.QueryRow("SELECT COUNT(idmember) as lm FROM prm_member where joindate >=? AND joindate <=?", formattedStart, formattedEnd).Scan(&lm)
	if err != nil {
		return 0, nil
	}
	return lm, nil
}
func totaltahun() (float64, error) {
	var th float64
	year := time.Now().Year()
	yearStr := strconv.Itoa(year)
	query := `SELECT COUNT(idmember) as th FROM prm_member where joindate like ?`
	err := database.DBMEMBER.QueryRow(query, "%"+yearStr+"%").Scan(&th)

	if err != nil {
		return 0, nil
	}
	return th, nil
}
