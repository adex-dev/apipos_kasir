package master

type Bankmodel struct {
	Idbalance     string `json:"id_balance"`
	Nomor         string `json:"nomor"`
	BalanceName   string `json:"balance_name"`
	PaymentType   string `json:"payment_type"`
	Status        string `json:"status"`
	LocationState string `json:"location_state"`
	AddOns        string `json:"add_ons"`
	Brands        string `json:"brands"`
	StatusUpdate  string `json:"status_update"`
	CreateUser    string `json:"create_user"`
	Nama_user     string `db:"nama_user"`
}
