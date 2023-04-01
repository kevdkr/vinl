package postgres

import (
	"database/sql"
	"log"
	"vinl/internal/models"

	"github.com/google/uuid"
)

type PostgresTransactionStorage struct {
	db *sql.DB
}

func NewPostgresTransactionStorage(db *sql.DB) *PostgresTransactionStorage {
	return &PostgresTransactionStorage{db: db}
}
// type PostgresAccountStorage struct {
// 	db *sql.DB
// }

func (storage *PostgresTransactionStorage) CreateTransaction(t *models.Transaction) (error) {
	transactionQuery := "INSERT INTO transactions (date, payee, comment, payee_comment, is_comment) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var transactionId uuid.UUID
	err := storage.db.QueryRow(transactionQuery, t.Date, t.Payee, t.Comment, t.PayeeComment, t.IsComment).Scan(&transactionId)
	//log.Printf("%d\n", transactionId)
	//log.Printf("%T\n", transactionId)
	//transactionId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting LastInsertId from transaction query result")
	}
	if err != nil {
		log.Printf("Error %s when inserting row into transactions table", err)
		//return err
	}
	for _, posting := range t.Postings {
		//log.Printf("%d: %v", index, account)
		postingQuery := "INSERT INTO postings (transactionid, name, amount, comment, is_comment) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		var postingId uuid.UUID
		//transactionIdQuery := "SELECT id FROM transactions" // should transaction id's get generated by the go code so I can select the transaction that was just added?
		//_, err = db.Exec(accountQuery, transactionId, account.Name, account.Amount)
		err := storage.db.QueryRow(postingQuery, transactionId, posting.Name, posting.Amount, posting.Comment, posting.IsComment).Scan(&postingId)
		if err != nil {
			log.Printf("Error %s when inserting row into postings table", err)
			//return err
		}

		// transactionsAccountsQuery := "INSERT INTO transactions_accounts (transaction_id, account_id) VALUES ($1, $2)"
		// _, err = storage.db.Exec(transactionsAccountsQuery, transactionId, accountId)
		// if err != nil {
		// 	log.Printf("Error %s when inserting row into transactions_accounts table", err)
		// }
	}

	//log.Printf("%d transactions created ", rows)
	return nil
}

func (storage *PostgresTransactionStorage) GetTransactions() (*models.Transactions, error) {
	transactionQuery := "SELECT id, date, payee, comment, payee_comment, is_comment FROM transactions"
	var transactions models.Transactions
	transactionRows, err := storage.db.Query(transactionQuery)
	checkError(err)
	defer transactionRows.Close()

	for transactionRows.Next() {
		var id uuid.UUID
		var date string
		var payee string
		var comment string
		var payeeComment string
		var isComment bool
		err = transactionRows.Scan(&id, &date, &payee, &comment, &payeeComment, &isComment)
		checkError(err)

		var postings []models.Posting
		postingsQuery := "SELECT id, transactionid, name, amount, comment, is_comment FROM postings WHERE transactionid = $1"
		postingRows, err := storage.db.Query(postingsQuery, id)
		checkError(err)
		defer postingRows.Close()
		for postingRows.Next() {
			var postingId uuid.UUID
			var transactionId uuid.UUID
			var name string
			var amount string
			var comment string
			var isComment bool
			err = postingRows.Scan(&postingId, &transactionId, &name, &amount, &comment, &isComment)
			checkError(err)
			p := models.Posting{
				Id:            postingId,
				TransactionId: transactionId,
				Name:          name,
				Amount:        amount,
				Comment:       comment,
				IsComment:     isComment,
			}
			postings = append(postings, p)
		}
		t := models.Transaction{
			Id:       id,
			Date:     date,
			Payee:    payee,
			Comment:  comment,
			PayeeComment: payeeComment,
			Postings: postings,
			IsComment: isComment,
		}
		transactions = append(transactions, t)
	}
	if transactions == nil {
		return &models.Transactions{}, nil
	}
	return &transactions, nil
}

func (storage *PostgresTransactionStorage) GetTransactionById(id string) (*models.Transaction, error) {
	transactionQuery := "SELECT id, date, payee, comment, payee_comment, is_comment FROM transactions WHERE id = $1"
	transactionRow := storage.db.QueryRow(transactionQuery, id)
	var t *models.Transaction
	var transactionId uuid.UUID
	var date string
	var payee string
	var comment string
	var payeeComment string
	var isComment bool
	err := transactionRow.Scan(&transactionId, &date, &payee, &comment, &payeeComment, &isComment)
	checkError(err)

	var postings []models.Posting
	postingsQuery := "SELECT id, transactionid, name, amount, comment, is_comment FROM postings WHERE transactionid = $1"
	postingRows, err := storage.db.Query(postingsQuery, id)
	checkError(err)
	defer postingRows.Close()
	for postingRows.Next() {
		var postingId uuid.UUID
		var transactionId uuid.UUID
		var name string
		var amount string
		var comment string
		var isComment bool
		err = postingRows.Scan(&postingId, &transactionId, &name, &amount, &comment, &isComment)
		checkError(err)
		a := models.Posting{
			Id:            postingId,
			TransactionId: transactionId,
			Name:          name,
			Amount:        amount,
			Comment:       comment,
			IsComment:     isComment,
		}
		postings = append(postings, a)
	}
	t = &models.Transaction{
		Id:       transactionId,
		Date:     date,
		Payee:    payee,
		Comment:  comment,
		PayeeComment: payeeComment,
		Postings: postings,
		IsComment: isComment,
	}
	return t, nil
}

func (storage *PostgresTransactionStorage) DeleteTransactionById(id string) error {

	// transactionsAccountsDeleteQuery := "DELETE FROM transactions_accounts WHERE transaction_id = $1"
	// res, err := storage.db.Exec(transactionsAccountsDeleteQuery, id)
	// checkError(err)
	// count, err := res.RowsAffected()
	// checkError(err)
	// log.Printf("Deleted %d transactions_accounts", count)

	postingDeleteQuery := "DELETE FROM postings WHERE transactionid = $1"
	res, err := storage.db.Exec(postingDeleteQuery, id)
	checkError(err)
	count, err := res.RowsAffected()
	checkError(err)
	log.Printf("Deleted %d postings", count)


	transactionDeleteQuery := "DELETE FROM transactions WHERE id = $1"
	res, err = storage.db.Exec(transactionDeleteQuery, id)
	checkError(err)
	count, err = res.RowsAffected()
	checkError(err)
	if count != 1 {
		log.Printf("Error: 1 row was supposed to be deleted, but %d was deleted", count)
	} else {
		log.Printf("Deleted %d transactions", count)
	}
	return nil //TODO return errors above
}

func checkError(err error) {
	if err != nil {
		log.Printf("%s", err)
	}
}