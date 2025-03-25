package models

type User struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Mobile_No int64   `json:"mobile_no"`
	Address   Address `json:"address"`
}

type Address struct {
	Vill     string `json:"vill"`
	Post     string `json:"post"`
	Ps       string `json:"p_s"`
	Distt    string `json:"distt"`
	State    string `json:"state"`
	Zip_Code int64  `json:"zip_code"`
}
