package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vinl/internal/service"

	"github.com/gorilla/mux"
)

type BalanceHandler struct {
	service *service.BalanceService
}

func NewBalanceHandler(balanceService *service.BalanceService) *BalanceHandler {
	return &BalanceHandler{balanceService}
}

func (h *BalanceHandler) HandleGetBalanceOfAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]

		balance, err := h.service.GetTotalDollarsOfAccount(id)
		if err != nil {
			log.Printf("%v", fmt.Errorf("Error getting total cents of account with id %v: %q", id, err))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(balance)
	}
}

func (h *BalanceHandler) HandleGetBalanceOfAccounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}

		balances, err := h.service.GetTotalDollarsOfAccounts()
		if err != nil {
			log.Printf("%v", fmt.Errorf("Error gtting balances of all accounts: %q", err))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(balances)
	}
}
