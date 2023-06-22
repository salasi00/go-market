package services

import (
	"github.com/salasi00/go-market/internal/models"
	"github.com/salasi00/go-market/internal/repositories"
)

type TransactionService struct {
	transactionRepository repositories.TransactionRepository
}

func NewTransactionService(transactionRepository repositories.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository}
}

func (s *TransactionService) FindTransaction(userId int) ([]models.Transaction, error) {
	transactions, err := s.transactionRepository.FindTransaction(userId)
	if err != nil {
		return []models.Transaction{}, err
	}

	return transactions, nil
}

func (s *TransactionService) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	createdTransaction, err := s.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return models.Transaction{}, err
	}

	return createdTransaction, nil
}
