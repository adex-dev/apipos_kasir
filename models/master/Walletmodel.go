package master

type Walletmodel struct {
	IdWallet      string `json:"id_wallet"`
	WalletName    string `json:"wallet_name"`
	PaymentType   string `json:"payment_type"`
	Status        string `json:"status"`
	LocationState string `json:"location_state"`
	AddOns        string `json:"add_ons"`
	Brands        string `json:"brands"`
	NamaUser      string `json:"nama_user"`
}
