package models

type Users struct {
	Nik       string `json:"nik"`
	Nama_user string `json:"nama_user"`
	Add_ons   string `json:"add_ons"`
	Status    string `json:"status"`
	Password  string `json:"password"`
	Level     string `json:"level"`
	State     string `json:"state"`
}
