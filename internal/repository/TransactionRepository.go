package repository

import (
	"github.com/MinhDuy27/Ecommerce-App/domain"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(CartID,UserID uint) error
	DeleteTransaction(TransId uint) error
	GetTransaction(TransId uint) (domain.Transaction, error)
	GetAllTransaction(userID uint, Amount int) ([]domain.Transaction, error)
}


func GetTransactionImage(db *gorm.DB) transactionRepository {
	return transactionRepository{
		Db: db,
	}
}

type transactionRepository struct {
	Db *gorm.DB
}

func (tr transactionRepository) CreateTransaction(CartID,UserID uint) error {
	transaction := domain.Transaction{
		
	}
	if err := tr.Db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}
func (tr transactionRepository) DeleteTransaction(TransId uint) error {
	var Transaction domain.Transaction
	if err := tr.Db.Delete(&Transaction,TransId).Error; err != nil {
		return err
	}
	return nil
}