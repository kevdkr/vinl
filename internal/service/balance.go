package service

import (
	"fmt"
	"strconv"
	"strings"
)

type BalanceService struct {
	postingService PostingService
	accountService AccountService
}

func NewBalanceService(postingService *PostingService, accountService *AccountService) *BalanceService {
	return &BalanceService{*postingService, *accountService}
}

func separateDollarsCents(currency string) (int64, int64, error) {

	if !strings.Contains(currency, ".") {
		dollars, err := strconv.ParseInt(currency, 10, 64)
		if err != nil {
			return 0, 0, err
		}
		return dollars, 0, nil
	}

	parts := strings.Split(currency, ".")
	// if len(parts) != 2 {
	// 	return 0, 0, fmt.Errorf("Invalid currency input")
	// }
	if len(parts) == 0 || len(parts) > 2 {
		return 0, 0, fmt.Errorf("Invalid currency input")
	}

	dollars, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	var cents int64
	if len(parts) == 2 {
		if len(parts[1]) != 2 {
			return 0, 0, fmt.Errorf("Invalid cents format")
		}
		cents, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return 0, 0, err
		}
	}

	return dollars, cents, nil
}

func calculateTotalCents(currency string) (int64, error) {
	dollars, cents, err := separateDollarsCents(currency)
	if err != nil {
		return 0, err
	}

	totalCents := dollars*100 + cents

	return totalCents, nil
}

func (s *BalanceService) GetTotalCentsOfAccount(id string) (int64, error) {
	postingsForAccount, err := s.postingService.GetPostingsByAccountId(id)
	if err != nil {
		return 0, err
	}

	var accountTotalCents int64
	for _, posting := range *postingsForAccount {
		totalCentsForPosting, err := calculateTotalCents(posting.Amount)
		if err != nil {
			return 0, err
		}
		accountTotalCents += totalCentsForPosting
	}

	return accountTotalCents, nil
}

func (s *BalanceService) GetTotalDollarsOfAccount(id string) (string, error) {
	cents, err := s.GetTotalCentsOfAccount(id)
	if err != nil {
		return "", err
	}

	var dollars int64
	dollars = cents / 100
	var newCents int64
	newCents = cents % 100

	dollarsString := strconv.Itoa(int(dollars))
	centsString := strconv.Itoa(int(newCents))

	if len(centsString) == 1 {
		if centsString == "0" {
			centsString += "0"
		}
		if centsString != "0" {
			centsString = "0" + centsString
		}

	}

	return dollarsString+"."+centsString, nil
}

func (s *BalanceService) GetTotalDollarsOfAccounts() ([]string, error) {
	accounts, err := s.accountService.GetAccounts()
	if err != nil {
		return nil, err
	}

	var balances []string
	for _, account := range *accounts {
		balance, err := s.GetTotalDollarsOfAccount(account.Id.String())
		if err != nil {
			return nil, err
		}
		balances = append(balances, balance)
	}
	return balances, nil
}

// func (s *TransactionService) GetBalanceOfAccount(id string) (float32, error) {
// 	transaction, err := s.storage.GetTransactionById(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return transaction, nil
// }
