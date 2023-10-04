package postgres

import (
	"database/sql"
	"log"
	"vinl/internal/models"

	"github.com/google/uuid"
)

type PostgresPostingStorage struct {
	db *sql.DB
}

func NewPostgresPostingStorage(db *sql.DB) *PostgresPostingStorage {
	return &PostgresPostingStorage{db: db}
}

func (storage *PostgresPostingStorage) GetPostings() (*[]models.Posting, error) {
	var postings []models.Posting
	postingsQuery := `SELECT p.id, p.transactionid, p.accountid, a.name, p.amount, p.comment, p.is_comment
						  FROM postings p
						  LEFT JOIN  accounts a ON a.id = p.accountid`

	postingRows, err := storage.db.Query(postingsQuery)
	if err != nil {
		log.Printf("%s", err)
	}

	defer postingRows.Close()
	for postingRows.Next() {
		var postingId uuid.UUID
		var transactionId uuid.UUID
		var accountId uuid.UUID
		var name string
		var amount string
		var comment string
		var isComment bool
		err = postingRows.Scan(&postingId, &transactionId, &accountId, &name, &amount, &comment, &isComment)
		if err != nil {
			log.Printf("%s", err)
		}
		p := models.Posting{
			Id:            postingId,
			TransactionId: transactionId,
			Account: models.Account{
				Id:   accountId,
				Name: name,
			},
			Amount:    amount,
			Comment:   comment,
			IsComment: isComment,
		}
		postings = append(postings, p)
	}
	return &postings, nil
}

func (storage *PostgresPostingStorage) GetPostingsByTransactionId(id string) (*[]models.Posting, error) {

	var postings []models.Posting
	postingsQuery := `SELECT p.id, p.transactionid, p.accountid, a.name, p.amount, p.comment, p.is_comment
						  FROM postings p
						  LEFT JOIN  accounts a ON a.id = p.accountid
						  WHERE p.transactionid = $1`
	postingRows, err := storage.db.Query(postingsQuery, id)

	checkError(err)
	defer postingRows.Close()
	for postingRows.Next() {
		var postingId uuid.UUID
		var transactionId uuid.UUID
		var accountId uuid.UUID
		var name string
		var amount string
		var comment string
		var isComment bool
		err = postingRows.Scan(&postingId, &transactionId, &accountId, &name, &amount, &comment, &isComment)
		checkError(err)
		a := models.Posting{
			Id:            postingId,
			TransactionId: transactionId,
			Account: models.Account{
				Id:   accountId,
				Name: name,
			},
			Amount:    amount,
			Comment:   comment,
			IsComment: isComment,
		}
		postings = append(postings, a)
	}
	return &postings, nil
}

func (storage *PostgresPostingStorage) GetPostingsByAccountId(id string) (*[]models.Posting, error) {

	var postings []models.Posting
	postingsQuery := `SELECT p.id, p.transactionid, p.accountid, a.name, p.amount, p.comment, p.is_comment
						  FROM postings p
						  LEFT JOIN  accounts a ON a.id = p.accountid
						  WHERE p.accountid = $1`
	postingRows, err := storage.db.Query(postingsQuery, id)

	checkError(err)
	defer postingRows.Close()
	for postingRows.Next() {
		var postingId uuid.UUID
		var transactionId uuid.UUID
		var accountId uuid.UUID
		var name string
		var amount string
		var comment string
		var isComment bool
		err = postingRows.Scan(&postingId, &transactionId, &accountId, &name, &amount, &comment, &isComment)
		checkError(err)
		a := models.Posting{
			Id:            postingId,
			TransactionId: transactionId,
			Account: models.Account{
				Id:   accountId,
				Name: name,
			},
			Amount:    amount,
			Comment:   comment,
			IsComment: isComment,
		}
		postings = append(postings, a)
	}
	return &postings, nil
}
