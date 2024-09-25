package domain

import "time"

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Code      int    `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verifired bool `json:"verifired"`
	UserType string `json:"usertype"`
}