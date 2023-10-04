package service

import (
	"vinl/internal/models"
	"vinl/internal/storage"
)

type PostingService struct {
	storage storage.PostingStorage
	//accountService AccountService
}

func NewPostingService(postingStorage storage.PostingStorage) *PostingService {
	return &PostingService{postingStorage}
}

func (s *PostingService) GetPostings() (*[]models.Posting, error) {
	postings, err := s.storage.GetPostings()
	if err != nil {
		return nil, err
	}
	return postings, nil
}

func (s *PostingService) GetPostingsByTransactionId(id string) (*[]models.Posting, error) {
	postings, err := s.storage.GetPostingsByTransactionId(id)
	if err != nil {
		return nil, err
	}

	return postings, nil
}

func (s *PostingService) GetPostingsByAccountId(id string) (*[]models.Posting, error) {
	postings, err := s.storage.GetPostingsByAccountId(id)
	if err != nil {
		return nil, err
	}

	return postings, nil
}
