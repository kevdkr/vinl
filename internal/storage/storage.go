package storage

import (
	"vinl/internal/models"

	"github.com/google/uuid"
)

type TransactionStorage interface {
	GetTransactions() (*models.Transactions, error)
	GetTransactionById(id string) (*models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) (error)
	DeleteTransactionById(id string) (error)
}

type AccountStorage interface {
	GetAccounts() (*[]models.Account, error)
	GetAccountById(id string) (*models.Account, error)
	GetAccountByName(name string) (*models.Account, error)
	CreateAccount(account *models.Account) (uuid.UUID, error)
	DeleteAccountById(id string) error
}

type PostingStorage interface {
	GetPostings() (*[]models.Posting, error)
	GetPostingsByTransactionId(id string) (*[]models.Posting, error)
	GetPostingsByAccountId(id string) (*[]models.Posting, error)
}
