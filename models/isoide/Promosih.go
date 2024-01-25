package isoide

type PromosiHlist struct {
	IdPromotion string `json:"id_promotion"`
	Type        string `json:"type"`
	Namep       string `json:"nameP"`
	DiscPersen  string `json:"discPersen"`
	DiscValue   string `json:"discValue"`
	Minimal     string `json:"Minimal"`
	Maximal     string `json:"Maximal"`
	Stdate      string `json:"stdate"`
	Enddate     string `json:"enddate"`
	State       string `json:"state"`
	Sttime      string `json:"sttime"`
	Endtime     string `json:"entime"`
	Outlet      string `json:"outlet"`
	AddOns      string `json:"addOn"`
	Member      string `json:"member"`
	Status      string `json:"status"`
	NamaUser    string `json:"nama_user"`
}
