package server

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"vinl/internal/models"
	"vinl/internal/transfer"

	"github.com/gorilla/mux"
)

func (s *Server) routes() {
	s.router.HandleFunc("/transactions", s.handleAddTransaction()).Methods(http.MethodPost)
	s.router.HandleFunc("/transactions", s.handleGetTransactions()).Methods(http.MethodGet)
	s.router.HandleFunc("/transactions/{id}", s.handleGetTransactionById()).Methods(http.MethodGet)
	//s.router.HandleFunc("/transactions/{id}/tofile", s.handleWriteTransactionToFile()).Methods(http.MethodGet)
	s.router.HandleFunc("/transactionstofile", s.handleWriteTransactionsToFile()).Methods(http.MethodGet)
	s.router.HandleFunc("/transactions/{id}", s.handleDeleteTransactionById()).Methods(http.MethodDelete)
	s.router.HandleFunc("/uploadfile", s.handleReadTransactionsFromFile()).Methods(http.MethodPost)
}

func (s *Server) handleAddTransaction() http.HandlerFunc {
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
		//log.Printf("%v", t)
		t.SaveTransaction(s.db)

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleGetTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		var transactions models.Transactions
		transactions = transactions.GetTransactions(s.db) // getTransactionsWithAccounts should return "error"
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transactions)
	}
}

func (s *Server) handleGetTransactionById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]
		//id, err := strconv.Atoi(vars["id"])
		//checkError(err)

		var transaction models.Transaction
		transaction = transaction.GetTransactionById(s.db, id)  // getTransactionsWithAccounts should return "error"
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transaction)
	}
}

func (s *Server) handleDeleteTransactionById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			log.Printf("405 Method no allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]

		var transaction models.Transaction
		transaction.DeleteTransactionById(s.db, id)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleWriteTransactionsToFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		var ts models.Transactions
		transactions := ts.GetTransactions(s.db)
		transfer.WriteTransactionsToFile(transactions)

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

func (s *Server) handleReadTransactionsFromFile() http.HandlerFunc {
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
		transfer.TransferTransactionFromFile(buf, s.db)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func checkError (err error) {
    if err != nil {
        log.Printf("%s", err)
    }
}
