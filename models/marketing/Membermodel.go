package marketing

type Membermodel struct {
	Idmember string `json:"idmember"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	NoHp     string `json:"no_hp"`
	Birthday string `json:"birthday"`
	Joindate string `json:"joindate"`
	Status   string `json:"status"`
}
