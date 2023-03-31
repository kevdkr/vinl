package storage

import (
	"vinl/internal/models"
)

type TransactionStorage interface {
	GetTransactions() (*models.Transactions, error)
	GetTransactionById(id string) (*models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) (error)
	DeleteTransactionById(id string) (error)
}

//type AccountStorage interface {}
