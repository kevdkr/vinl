package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"vinl/internal/models"
	"vinl/internal/service"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService}
}

func (h *AccountHandler) HandleAddAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			log.Printf("405 Method not allowed")
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error processing body in request: %s", err)
		}
		var a models.Account
		err = json.Unmarshal(body, &a)
		if err != nil {
			log.Printf("Error unmarshalling json: %s", err)
		}
		_, err = h.service.CreateAccount(&a)
		if err != nil {
			log.Printf("error saving account")
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *AccountHandler) HandleGetAccounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		accounts, err := h.service.GetAccounts()// getTransactionsWithAccounts should return "error"
		if err != nil {
			log.Printf("error getting accounts")
			return //TODO what to return here?
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accounts)
	}
}

func (h *AccountHandler) HandleGetAccountById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("405 Method not allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]
		//checkError(err)

		account, err := h.service.GetAccountById(id)  // getTransactionsWithAccounts should return "error"
		if err != nil {
			log.Printf("error getting account with id %v", id)
			return //TODO what to return here
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(account)
	}
}

func (h *AccountHandler) HandleDeleteAccountById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			log.Printf("405 Method no allowed")
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]

		err := h.service.DeleteAccountById(id)
		if err != nil {
			log.Printf("error deleting account %v", err)
			return //TODO what to return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *AccountHandler) HandleWriteAccountsToFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accounts, err := h.service.GetAccounts() //TODO move this to the TransferAccountsToFileService
		if err != nil {
			log.Printf("error getting accounts %v", err)
			return
		}
		err = h.service.TransferAccountsToFile(accounts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *AccountHandler) HandleReadAccountsFromFile() http.HandlerFunc {
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
		h.service.TransferAccountsFromFile(buf)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
