package service

import (
	"log"
	"strconv"

	_"github.com/MinhDuy27/Ecommerce-App/domain"
	"github.com/MinhDuy27/Ecommerce-App/internal/helper"
	"github.com/MinhDuy27/Ecommerce-App/internal/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepository
	Auth helper.Auth
}

func (t *TransactionService) CreateTransaction(CartID,UserID string) error {
	uintCartID, err := strconv.ParseUint(CartID, 10, 64)
	if err != nil {
		return err
	}
	UintUserID, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		return err
	}
	if err := t.Repo.CreateTransaction(uint(uintCartID),uint(UintUserID)); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
	
}
func (t *TransactionService) DeleteTransaction(TransId string) error {
	trans_Uint, err := strconv.ParseUint(TransId, 10, 64)
	if err != nil {
		return err
	}
	if err = t.Repo.DeleteTransaction(uint(trans_Uint));err != nil {
		return err
	}
	return nil
}
// func (t* TransactionService) GetTransaction(TransId string) (domain.Transaction, error){
// 	trans_Uint, err := strconv.ParseUint(TransId, 10, 64)
// 	if err != nil {
// 		return domain.Transaction{},err
// 	}
// 	transaction,err = t.Repo.GetTransaction(uint(trans_Uint))
// 	if err != nil {
// 		return domain.Transaction{},err
// 	}
// 	return transaction,nil
// }
// func (t* TransactionService) GetAllTransaction(userID string,Amount int) ([]domain.Transaction, error){
// 	trans_Uint, err := strconv.ParseUint(userID, 10, 64)
// 	if err != nil {
// 		return[]domain.Transaction{},err
// 	}
// 	transaction,err = t.Repo.GetTransaction(uint(trans_Uint))
// 	if err != nil {
// 		return []domain.Transaction{},err
// 	}
// 	return transaction,nil
// }