package domain

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"Primarikey"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"index;unique;not null"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Code      int    `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verifired bool `json:"verifired" gorm:"default:false"`
	UserType string `json:"usertype" gorm:"default: buyer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`

}