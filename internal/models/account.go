package models

import (
	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID `json:"id"`
	Name          string 	`json:"name"`
}
