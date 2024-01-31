package master

type Usermodel struct {
	Nik      string `json:"nik"`
	NamaUser string `json:"nama_user"`
	AddOns   string `json:"add_ons"`
	Level    string `json:"level"`
	State    string `json:"state"`
}
