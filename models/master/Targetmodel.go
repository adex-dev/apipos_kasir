package master

type TargetModel struct {
	Date       string `json:"date"`
	Amount     string `json:"amount"`
	StoreId    string `json:"store_id"`
	NamaOutlet string `json:"nama_outlet"`
}
