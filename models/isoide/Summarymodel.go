package isoide

type KategoriSummary struct {
	Namakategori string `json:"nama_kategori"`
	Terminal     string `json:"terminal"`
	Total        int    `json:"total"`
}

type Storesummarymodel struct {
	NamaStore     string `json:"nama_store"`
	Telp1         string `json:"telp_1"`
	Telp2         string `json:"telp_2"`
	Alamat        string `json:"alamat"`
	ServiceCharge string `json:"service_charge"`
	Tax           string `json:"tax"`
}

type PaymentSummarycardmodel struct {
	TypePayment  string `json:"type_payment"`
	JenisPayment string `json:"jenis_payment"`
	PaymentName  string `json:"payment_name"`
	Total        int    `json:"total"`
}

type Plusummary struct {
	NoTrans        string `json:"no_trans"`
	ItemNama       string `json:"itemName"`
	NamaKategori   string `json:"nama_kategori"`
	NamaSubitem    string `json:"nama_subitem"`
	CloseTransaksi string `json:"close_transaksi"`
	TimeTransaksi  string `json:"time_transaksi"`
	Status         string `json:"status_transaksi"`
	IdProduct      string `json:"id_product"`
}
type Plusummarys struct {
	NoTrans      string  `json:"no_trans"`
	ItemNama     string  `json:"itemName"`
	NamaKategori string  `json:"nama_kategori"`
	NamaSubitem  string  `json:"nama_subitem"`
	Status       string  `json:"status_transaksi"`
	IdProduct    string  `json:"id_product"`
	Price        float64 `json:"price"`
	Qty          float64 `json:"qty"`
	Additional   string  `json:"additional"`
}
type Dataprinter struct {
	Price      float64 `json:"price"`
	Qty        float64 `json:"qty"`
	Additional string  `json:"additional"`
}

type DataNilai struct {
	NoTrans         string  `json:"no_trans"`
	Gross           float64 `json:"gross_sales"`
	DiscItem        float64 `json:"disc_item"`
	DiscOther       float64 `json:"disc_other"`
	Tax             float64 `json:"tax"`
	ServiceCharge   float64 `json:"service_charge"`
	TotalPayment    float64 `json:"total_payment"`
	NetSales        float64 `json:"net_sales"`
	Changed         float64 `json:"changed"`
	ChangeInvisible float64 `json:"change_invisible"`
	StatusPayment   string  `json:"status_payment"`
	ClosePayment    string  `json:"close_payment"`
	TypePayment     string  `json:"type_payment"`
	PaymentName     string  `json:"payment_name"`
	NominalPayment  float64 `json:"nominal_payment"`
	JenisPayment    string  `json:"jenis_payment"`
}
