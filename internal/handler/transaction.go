package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"vinl/internal/models"
	"vinl/internal/transfer"
	"vinl/internal/service"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) HandleAddTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			log.Printf("405 Method not allowed")
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error processing body in request: %s", err)
		}
		var t models.Transaction
		err = json.Unmarshal(body, &t)
		if err != nil {
			log.Printf("Error unmarshalling json: %s", err)
		}
		err = h.service.CreateTransaction(&t)
		if err != nil {
			log.Printf("error saving transaction")
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *TransactionHandler) HandleGetTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		transactions, err := h.service.GetTransactions()// getTransactionsWithAccounts should return "error"
		if err != nil {
			log.Printf("error getting transactions")
			return //TODO what to return here?
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transactions)
	}
}

func (h *TransactionHandler) HandleGetTransactionById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]
		//checkError(err)

		transaction, err := h.service.GetTransactionById(id)  // getTransactionsWithAccounts should return "error"
		if err != nil {
			log.Printf("error getting transaction with id %v", id)
			return //TODO what to return here
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transaction)
	}
}

func (h *TransactionHandler) HandleDeleteTransactionById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			log.Printf("405 Method no allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]

		err := h.service.DeleteTransactionById(id)
		if err != nil {
			log.Printf("error deleting transaction %v", err)
			return //TODO what to return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *TransactionHandler) HandleWriteTransactionsToFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		transactions, err := h.service.GetTransactions()
		if err != nil {
			log.Printf("error getting transactions %v", err)
			return
		}
		transfer.WriteTransactionsToFile(*transactions)

		//w.WriteHeader(http.StatusOK)

		// if r.Method != http.MethodGet {
		// 	log.Printf("405 Method not allowed")
		// 	return
		// }
		//vars := mux.Vars(r)
		//id, err := strconv.Atoi(vars["id"])
		//checkError(err)
	}
}

func (h *TransactionHandler) HandleReadTransactionsFromFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			log.Printf("405 Method no allowed")
			return
		}

		r.ParseMultipartForm(32 << 20)

		var content []byte
		file, _, err := r.FormFile("file")
		if err != nil {
			log.Printf("%v", err)
		}
		defer file.Close()

		buf := bytes.NewBuffer(content)
		io.Copy(buf, file)
		h.service.TransferTransactionFromFile(buf)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func checkError (err error) {
    if err != nil {
        log.Printf("%s", err)
    }
}
