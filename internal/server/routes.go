package server

import (
	"net/http"
	"vinl/internal/service"
	"vinl/internal/handler"
)

func (s *Server) routes() {
	// s.router.HandleFunc("/transactions", s.handleAddTransaction()).Methods(http.MethodPost)
	// s.router.HandleFunc("/transactions", s.handleGetTransactions()).Methods(http.MethodGet)
	// s.router.HandleFunc("/transactions/{id}", s.handleGetTransactionById()).Methods(http.MethodGet)
	// //s.router.HandleFunc("/transactions/{id}/tofile", s.handleWriteTransactionToFile()).Methods(http.MethodGet)
	// s.router.HandleFunc("/transactionstofile", s.handleWriteTransactionsToFile()).Methods(http.MethodGet)
	// s.router.HandleFunc("/transactions/{id}", s.handleDeleteTransactionById()).Methods(http.MethodDelete)
	// s.router.HandleFunc("/uploadfile", s.handleReadTransactionsFromFile()).Methods(http.MethodPost)
}

func (s *Server) registerTransactionRoutes(transactionService *service.TransactionService) {
	transactionHandler := handler.NewTransactionHandler(transactionService)

	s.router.HandleFunc("/transactions", transactionHandler.HandleAddTransaction()).Methods(http.MethodPost)
	s.router.HandleFunc("/transactions", transactionHandler.HandleGetTransactions()).Methods(http.MethodGet)
	s.router.HandleFunc("/transactions/{id}", transactionHandler.HandleGetTransactionById()).Methods(http.MethodGet)
	s.router.HandleFunc("/transactions/{id}", transactionHandler.HandleDeleteTransactionById()).Methods(http.MethodDelete)
	s.router.HandleFunc("/transactionstofile", transactionHandler.HandleWriteTransactionsToFile()).Methods(http.MethodGet)
	s.router.HandleFunc("/uploadfile", transactionHandler.HandleReadTransactionsFromFile()).Methods(http.MethodPost)
}

func (s *Server) registerAccountRoutes(accountService *service.AccountService) {
	accountHandler := handler.NewAccountHandler(accountService)

	s.router.HandleFunc("/accounts", accountHandler.HandleAddAccount()).Methods(http.MethodPost)
	s.router.HandleFunc("/accounts", accountHandler.HandleGetAccounts()).Methods(http.MethodGet)
	s.router.HandleFunc("/accounts/{id}", accountHandler.HandleGetAccountById()).Methods(http.MethodGet)
	s.router.HandleFunc("/accounts/{id}", accountHandler.HandleDeleteAccountById()).Methods(http.MethodDelete)
	s.router.HandleFunc("/accounts/tofile", accountHandler.HandleWriteAccountsToFile()).Methods(http.MethodGet)
	s.router.HandleFunc("/accounts/uploadfile", accountHandler.HandleReadAccountsFromFile()).Methods(http.MethodPost)
}

func (s *Server) registerPostingRoutes(postingService *service.PostingService) {
	postingHandler := handler.NewPostingHandler(postingService)

	s.router.HandleFunc("/postings", postingHandler.HandleGetPostings()).Methods(http.MethodGet)
	s.router.HandleFunc("/postings/transaction/{id}", postingHandler.HandleGetPostingsByTransactionId()).Methods(http.MethodGet)
	s.router.HandleFunc("/postings/account/{id}", postingHandler.HandleGetPostingsByAccountId()).Methods(http.MethodGet)
}

func (s *Server) registerBalanceRoutes(balanceService *service.BalanceService) {
	balanceHandler := handler.NewBalanceHandler(balanceService)

	s.router.HandleFunc("/balance/{id}", balanceHandler.HandleGetBalanceOfAccount()).Methods(http.MethodGet)
	s.router.HandleFunc("/balances", balanceHandler.HandleGetBalanceOfAccounts()).Methods(http.MethodGet)
}
