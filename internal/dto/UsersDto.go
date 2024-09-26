package dto

type Logindto struct {
	Email	string	`json:"email"`
	Password string `json:"password"`
}
type SignUpdto struct {
	Logindto
	Phone string `json:"phone"`
}