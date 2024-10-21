package domain

import "time"

type Product struct {
	ID          uint    `json:"id" gorm:"Primarikey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image_url   string  `json:"image_url" gorm:"unique,default:null"`
	Price       float64 `json:"price"`
	Quantity   uint    `json:"quantity"`
	UserID      uint    `json:"user_id" `
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`

}