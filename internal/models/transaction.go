package models

import (
	"github.com/google/uuid"
)

type Transaction struct {
	Id        uuid.UUID `json:"id"`
	Date      string    `json:"date"`
	Payee     string    `json:"payee"`
	PayeeComment string `json:"payee_comment"`
	Comment   string    `json:"comment"` // omitempty
	Postings  []Posting `json:"postings"`
	IsComment bool      `json:"is_comment"`
}

type Transactions []Transaction
