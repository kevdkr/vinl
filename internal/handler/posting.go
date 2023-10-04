package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"vinl/internal/service"

	"github.com/gorilla/mux"
)

type PostingHandler struct {
	service *service.PostingService
}

func NewPostingHandler(postingService *service.PostingService) *PostingHandler {
	return &PostingHandler{postingService}
}

func (h *PostingHandler) HandleGetPostings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		postings, err := h.service.GetPostings()// getTransactionsWithAccounts should return "error"
		if err != nil {
			log.Printf("error getting postings")
			return //TODO what to return here?
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(postings)
	}
}

func (h *PostingHandler) HandleGetPostingsByTransactionId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]

		postings, err := h.service.GetPostingsByTransactionId(id)
		if err != nil {
			log.Printf("error getting postings with transaction id %v", id)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(postings)
	}
}

func (h *PostingHandler) HandleGetPostingsByAccountId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]

		postings, err := h.service.GetPostingsByAccountId(id)
		if err != nil {
			log.Printf("error getting postings with account id %v", id)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(postings)
	}
}
