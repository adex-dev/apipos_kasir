package master

type Cashmodel struct {
	IdBalance     string `json:"id_balance"`
	BalanceName   string `json:"balance_name"`
	PaymentType   string `json:"payment_type"`
	Status        string `json:"status"`
	LocationState string `json:"location_state"`
	AddOns        string `json:"add_ons"`
	NamaUser      string `json:"nama_user"`
}
