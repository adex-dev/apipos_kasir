package Master

import (
	"apipos/database"
	"apipos/models/master"
	"github.com/gofiber/fiber/v2"
)

func WalletResource(c *fiber.Ctx) error {
	var wallet []master.Walletmodel
	query := `SELECT prm_balance_wallet.id_wallet,prm_balance_wallet.wallet_name,prm_balance_wallet.payment_type,prm_balance_wallet.status,prm_balance_wallet.location_state,prm_balance_wallet.add_ons,prm_balance_wallet.brands,users.nama_user from prm_balance_wallet inner join users on users.nik = prm_balance_wallet.create_user order by prm_balance_wallet.wallet_name asc`
	result, err := database.DBMASTER.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": "", "respon": 501})
	}
	defer result.Close()
	for result.Next() {
		var wl master.Walletmodel
		err := result.Scan(&wl.IdWallet, &wl.WalletName, &wl.PaymentType, &wl.Status, &wl.LocationState, &wl.AddOns, &wl.Brands, &wl.NamaUser)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"data": "", "respon": 401})
		}
		wallet = append(wallet, wl)
	}
	return c.JSON(fiber.Map{"data": wallet, "respon": 200})

}
