package models

import (
	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID `json:"-"`
	Name          string 	`json:"name"`
}
