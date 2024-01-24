package nahm

type Kategorilist struct {
	Idkategori   string `json:"id_kategori"`
	Nomer        string `json:"nomer"`
	Namakategori string `json:"nama_kategori"`
	Addons       string `json:"add_ons"`
	Status       string `json:"status"`
}
