package models

type Users struct {
	IdUsers int    `json:"idUsers"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Pwd     string `json:"pwd"`
	IsAdmin bool   `json:"isAdmin"`
}
