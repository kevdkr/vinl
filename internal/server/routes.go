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
	//s.router.HandleFunc("/transactions/{id}/tofile", s.handleWriteTransactionToFile()).Methods(http.MethodGet)
	s.router.HandleFunc("/transactionstofile", transactionHandler.HandleWriteTransactionsToFile()).Methods(http.MethodGet)
	s.router.HandleFunc("/transactions/{id}", transactionHandler.HandleDeleteTransactionById()).Methods(http.MethodDelete)
	s.router.HandleFunc("/uploadfile", transactionHandler.HandleReadTransactionsFromFile()).Methods(http.MethodPost)
}
