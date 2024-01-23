package master

type Storelist struct {
	Idoutlet      string `json:"id_outlet"`
	Nomor         string `json:"nomor"`
	Namaoutlet    string `json:"nama_outlet"`
	Addons        string `json:"add_ons"`
	Createuser    string `json:"create_user"`
	Status        string `json:"status"`
	Token         string `json:"token"`
	Brands        string `json:"brands"`
	Prefix        string `json:"prefix"`
	ServiceCharge int    `json:"service_charge"`
	Tax           int    `json:"tax"`
	Alamat        string `json:"alamat"`
	Telp1         string `json:"telp_1"`
	Telp2         string `json:"telp_2"`
	Nama_user     string `json:"nama_user"`
}
