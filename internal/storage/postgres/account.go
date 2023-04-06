package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"vinl/internal/models"

	"github.com/google/uuid"
)

type PostgresAccountStorage struct {
	db *sql.DB
}

func NewPostgresAccountStorage(db *sql.DB) *PostgresAccountStorage {
	return &PostgresAccountStorage{db: db}
}

func (storage *PostgresAccountStorage) CreateAccount(a *models.Account) (error) {
	accountQuery := "INSERT INTO accounts (name) VALUES ($1) RETURNING id"
	var accountId uuid.UUID
	err := storage.db.QueryRow(accountQuery, a.Name).Scan(&accountId)
	if err != nil {
		return fmt.Errorf("Error inserting account into database: ", err)
	}
	return nil
}

func (storage *PostgresAccountStorage) GetAccounts() (*[]models.Account, error) {
	accountQuery := "SELECT id, name FROM accounts"
	var accounts []models.Account
	accountRows, err := storage.db.Query(accountQuery)
	if err != nil {
		return nil, fmt.Errorf("Error querying accounts from database: ", err)
	}
	defer accountRows.Close()

	for accountRows.Next() {
		var id uuid.UUID
		var name string

		err = accountRows.Scan(&id, &name)
		if err != nil {
			return nil, fmt.Errorf("Error scanning account rows from database: ", err)
		}
		a := models.Account{
			Id: id,
			Name: name,
		}
		accounts = append(accounts, a)
	}
	if accounts == nil {
		return &[]models.Account{}, nil
	}
	return &accounts, nil
}

func (storage *PostgresAccountStorage) GetAccountById(id string) (*models.Account, error) {
	accountQuery := "SELECT id, name FROM account WHERE id = $1"
	accountRow := storage.db.QueryRow(accountQuery, id)
	var a *models.Account
	var accountid uuid.UUID
	var name string
	err := accountRow.Scan(&id, &name)
	if err != nil {
		return nil, fmt.Errorf("Error scanning account row onto vars: %q", err)
	}
	a = &models.Account{
		Id: accountid,
		Name: name,
	}
	return a, nil
}

func (storage *PostgresAccountStorage) DeleteAccountById(id string) error {
	accountDeleteQuery := "DELETE FROM accounts WHERE id = $1" //TODO what is account is referenced by postings
	res, err := storage.db.Exec(accountDeleteQuery, id)
	if err != nil {
		return fmt.Errorf("Error deleting account with id %q from accounts table: %q", id, err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting count deleted from accounts table: %q", err)
	}
	if count != 1 {
		return fmt.Errorf("Error: 1 row was supposed to be deleted from the accounts table (account with id %q), but %d rows were deleted", id, count)
	} else {
		log.Printf("Deleted account with id %q from accounts table", id)
	}
	return nil
}
