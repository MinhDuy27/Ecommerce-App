package repository

import "gorm.io/gorm"

type CartRepository interface {
	// CreateCart(UserID uint) error
	// DeleteCart(CartID uint) error
	// AddToCart(CartID, ProductID uint, quantity int) error
	// RemoveFromCart(CartID, ProductID uint, quantity int) error
}

func GetCartImage(db *gorm.DB) cartRepository {
	return cartRepository{
		DB: db,
	}
}
type cartRepository struct {
	DB *gorm.DB
}

