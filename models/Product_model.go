package models

type Productlist struct {
	Id_item      string `json:"id_item"`
	Nomor_items  string `json:"nomor_items"`
	ItemName     string `json:"itemName"`
	ItemKategori string `json:"itemKategori"`
	ItemSub      string `json:"itemSub"`
	ItemPrice    string `json:"itemPrice"`
	ItemOptional string `json:"itemOptional"`
	ItemStatus   string `json:"itemStatus"`
	Adds_date    string `json:"adds_date"`
	Adds_time    string `json:"adds_time"`
	State        string `json:"state"`
	Stdate       string `json:"stdate"`
	Endate       string `json:"endate"`
	Dist_state   string `json:"dist_state"`
	Infor        string `json:"infor"`
	Create_ons   string `json:"create_ons"`
	Nama_user    string `db:"nama_user"`
	NamaKategori string `db:"nama_kategori"`
	NamaSubitem  string `db:"nama_subitem"`
}
