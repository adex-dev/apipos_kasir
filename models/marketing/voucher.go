package marketing

type Voucherredemmodel struct {
	Vouchercode  string `json:"vouchercode"`
	Vouchername  string `json:"vouchername"`
	Tanggal      string `json:"tanggal"`
	Claimdate    string `json:"claim_date"`
	Claimstore   string `json:"claim_store"`
	Status       string `json:"status"`
	Persentase   string `json:"persentase"`
	Valuenominal string `json:"valuenominal"`
}

type Vouchermarketingmodel struct {
	Vouchercode string `json:"vouchercode"`
	Vouchername string `json:"namevoucher"`
	Qty         string `json:"qty"`
	Expired     string `json:"expired"`
	Disvalue    string `json:"disvalue"`
	Dispersent  string `json:"dispersent"`
}

type Voucherultahmodel struct {
	Vouchercode  string `json:"vouchercode"`
	Birthday     string `json:"birthday"`
	Validvoucher string `json:"validvoucher"`
	Status       string `json:"status"`
	Nama         string `json:"nama"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Nohps        string `json:"no_hp"`
	Joindate     string `json:"joindate"`
}
