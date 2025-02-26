package service

import (

	"github.com/MinhDuy27/Ecommerce-App/internal/helper"
	"github.com/MinhDuy27/Ecommerce-App/internal/repository"
)

type CartService struct {
	Repo repository.CartRepository
	Auth  helper.Auth
}

// func(c* CartService) CreateCart(UserID string) error {
// 	Id_uint,err := strconv.ParseUint(UserID, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	if err := c.Repo.CreateCart(uint(Id_uint)); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c* CartService) DeleteCart(CartID string) error {
// 	Id_uint,err := strconv.ParseUint(CartID, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	if err := c.Repo.DeleteCart(uint(Id_uint)); err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (c* CartService) AddToCart(CartID,ProductID string,quantity int) error {
// 	CartID_uint,err := strconv.ParseUint(CartID, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	ProductID_uint,err := strconv.ParseUint(ProductID, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	if err := c.Repo.AddToCart(uint(CartID_uint),uint(ProductID_uint),quantity); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c* CartService) RemoveFromCart(CartID,ProductID string,quantity int) error {
// 	CartID_uint,err := strconv.ParseUint(CartID, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	ProductID_uint,err := strconv.ParseUint(ProductID, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	if err := c.Repo.RemoveFromCart(uint(CartID_uint),uint(ProductID_uint),quantity); err != nil {
// 		return err
// 	}
// 	return nil
// }