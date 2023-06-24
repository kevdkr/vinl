package service

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"vinl/internal/models"
	"vinl/internal/storage"

	"github.com/google/uuid"
)

type AccountService struct {
	storage storage.AccountStorage
}

func NewAccountService(accountStorage storage.AccountStorage) *AccountService {
	return &AccountService{accountStorage}
}

func (s *AccountService) GetAccountById(id string) (*models.Account, error) {
	account, err := s.storage.GetAccountById(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountService) GetAccountByName(name string) (*models.Account, error) {
	account, err := s.storage.GetAccountByName(name)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountService) GetAccounts() (*[]models.Account, error) {

	accounts, err := s.storage.GetAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *AccountService) CreateAccount(a *models.Account) (uuid.UUID, error) {
	accountId, err := s.storage.CreateAccount(a)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error creating account %q: %q", a.Name, err)
	}
	return accountId, nil
}

func (s *AccountService) CreateAccounts(accounts *[]models.Account) error {
	for _, a := range *accounts {
		_, err := s.CreateAccount(&a)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *AccountService) DeleteAccountById(id string) error {
	return s.storage.DeleteAccountById(id)
}

func (s *AccountService) TransferAccountsToFile(accounts *[]models.Account) error {
	return errors.New("Not implemented")
}

func (s *AccountService) TransferAccountsFromFile(buf *bytes.Buffer) error {

	content := buf.Bytes()
	reader := bytes.NewReader(content)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	accounts, err := parseAccountsFile(reader)
	if err != nil {
		return err
	}

	return s.CreateAccounts(accounts)
}

func parseAccountsFile(reader io.Reader) (*[]models.Account, error){
	return nil, errors.New("Not implemented")
}
