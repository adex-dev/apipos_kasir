package master

type Outlet struct {
	Idoutlet string `json:"id_outlet "`
	Name     string `json:"nama_outlet"`
	Brands   string `json:"brands"`
}

type TransactionParameter struct {
	IDLokasi           int    `db:"id_lokasi"`
	ClosePayment       string `db:"close_transaksi"`
	Notrans            string `db:"no_trans"`
	StatusTransaksi    string `db:"status_transaksi"`
	TransactionBrands  string `db:"brands"`
	TransactionOutlets string `db:"nama_outlet"`
}
type StoreSales struct {
	NamaStore string `json:"namastore"`
	Jumlah    int    `json:"jumlah"`
}
type Result struct {
	DataIsoide  []StoreSales `json:"dataisoide"`
	DataNahm    []StoreSales `json:"datanahm"`
	TotalIsoide int          `json:"totalisoide"`
	TotalNahm   int          `json:"totalnahm"`
}

type SalesResult struct {
	Netsales int
}
