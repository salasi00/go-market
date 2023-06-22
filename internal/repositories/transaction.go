package repositories

import (
	"github.com/salasi00/go-market/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction(userId int) ([]models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
}

func NewTransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction(userId int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id=?", userId).Preload("User").Preload("Product.User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Product").Preload("User").Create(&transaction).Error

	return transaction, err
}
