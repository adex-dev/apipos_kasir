package master

type Tendermodel struct {
	IdTender       string `json:"id_tender"`
	TenderName     string `json:"tender_name"`
	PaymentType    string `json:"payment_type"`
	Status         string `json:"status"`
	LocationState  string `json:"location_state"`
	AddOns         string `json:"add_ons"`
	Brands         string `json:"brands"`
	Persen         string `json:"persen"`
	BalanceNominal string `json:"balance_nominal"`
	BalanceMinimal string `json:"balance_minimal"`
	BalanceMaximal string `json:"balance_maximal"`
	StTime         string `json:"st_time"`
	StDate         string `json:"st_date"`
	NamaUser       string `json:"nama_user"`
}
