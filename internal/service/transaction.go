package service

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"vinl/internal/models"
	"vinl/internal/storage"
)

type TransactionService struct {
	storage storage.TransactionStorage
}

func NewTransactionService(transactionStorage storage.TransactionStorage) *TransactionService {
	return &TransactionService{transactionStorage}
}

func (s *TransactionService) GetTransactionById(id string) (*models.Transaction, error) {
	transaction, err := s.storage.GetTransactionById(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) GetTransactions() (*models.Transactions, error) {

	transactions, err := s.storage.GetTransactions()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *TransactionService) CreateTransaction(t *models.Transaction) error {
	return s.storage.CreateTransaction(t)
}

func (s *TransactionService) CreateTransactions(transactions *models.Transactions) error {
	for _, t := range *transactions {
		err := s.CreateTransaction(&t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *TransactionService) DeleteTransactionById(id string) error {
	return s.storage.DeleteTransactionById(id)
}

func (s *TransactionService) TransferTransactionsToFile(ts *models.Transactions) error {
	f, err := os.Create("ledger.dat")
	if err != nil {
		log.Printf("%s", err)
	}

	defer f.Close()

	for _, t := range *ts {
		if t.IsComment == false {

			f.WriteString(t.Date + " * ")
			f.WriteString(t.Payee + t.PayeeComment + "\n")
			for _, p := range t.Postings {
				f.WriteString("    " + p.Account.Name + "    ")
				f.WriteString(" " + p.Amount + "  " + p.Comment + "\n")
			}
			//for _, comment := range t.Comment {
			f.WriteString(t.Comment)
			//}
			f.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			//f.WriteString(t.Payee + "\n")
			//for _, comment := range t.Comment {
			f.WriteString(t.Comment + "\n")
			//}
			if err != nil {
				log.Fatal(err) //TODO return errors
			}
		}
	}

	return nil
}

func (s *TransactionService) TransferTransactionFromFile(buf *bytes.Buffer) error {

	content := buf.Bytes()
	reader := bytes.NewReader(content)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	ts, err := parseFile(reader)
	if err != nil {
		log.Printf("%v", err)
	}

	return s.CreateTransactions(ts)
}

const (
	ledgerdate = "\\d{4}\\/(0[1-9]|1[0-2])\\/\\d{2}"
	ledgerpayeeclearedregex = "^[0-9]+[-/][-/.=0-9]+\\s-\\*\\s-+\\(([^)]+)\\s-+\\)?\\([^*].+?\\)\\s-*\\(;\\|$\\)"
	ledgeraccountnameregex = "\\(?1:[^][(); \t\r\n]+\\(?: [^][(); \t\r\n]+\\)*\\)"
	ledgeraccountamountregex = "-?[0-9][0-9,]*[.]?[0-9]*"
	myaccountnameregex = "[a-zA-Z0-9:]+[ ]?[a-zA-Z0-9:]+"
	myaccountnameregexincludeparenthesis = "[(]?[a-zA-Z0-9:]+[ ]?[a-zA-Z0-9:]+[)]?"
	myaccountamountregex = " {2,}\\$?-?[0-9][0-9,]*[.]?[0-9]*"
	myaccountamountregexwithexpressions = " {2,}[(]?\\$?-?[0-9][0-9,]*[.]?[0-9]* ?[+-/*]? ?\\$?-?[0-9][0-9,]*[.]?[0-9]*[)]?"
	myaccountamountregexwithmultiexpressions = " {2,}[(]?\\$?-?[0-9][0-9,]*[.]?[0-9]*( ?[\\+\\-\\/\\*]? ?\\$?-?[0-9][0-9,]*[.]?[0-9]*){0,}[)]?"
	myaccountamountregexwithmultiexpressionsandhandlestockvalues = " {2,}[(]?\\$?-?[0-9][0-9,]*[.]?[0-9]*( ?[\\+\\-\\/\\*]? ?\\$?-?[0-9][0-9,]*[.]?[0-9]*){0,}[)]?[^;]*"
	mycommentregex = ";{1,}.*"
)

func parseFile(reader io.Reader) (*models.Transactions, error){

	accountnameregex, err := regexp.Compile(myaccountnameregexincludeparenthesis)
	if err != nil {
		//return nil, errors.Errorf("Error compiling regex for account names: ", err)
		return nil, err
	}
	amountregex, err := regexp.Compile(myaccountamountregexwithmultiexpressionsandhandlestockvalues)
	if err != nil {
		log.Printf("%v", err)
		//return nil, errors.Errorf("Error compiling regex for account amount: ", err)
		return nil, err
	}

	commentregex, err := regexp.Compile(mycommentregex)
	if err != nil {
		//return nil, errors.Errorf("Error compiling regex for comment: ", err)
		return nil, err
	}

	var ts models.Transactions
	var ps []models.Posting
	var t models.Transaction

	var line string

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line = scanner.Text()
		commentCheck := strings.TrimSpace(line)
		if len(line) == 0 {

			if t.Date == "" && t.Payee == "" {  //TODO is this needed at the top and bottom?
				t.IsComment = true
			}
			ts = append(ts, t)
			ps = []models.Posting{}
			t = models.Transaction{}
			continue
		} else if (len(line) > 0) && (line[0:4] != "    ") {
			if strings.HasPrefix(strings.TrimSpace(line), ";") {
				t.Comment += commentCheck + "\n"
				//comment := strings.TrimSpace(line)
				//c := models.Transaction {
				// 	Payee: comment,
				// 	IsComment: true,
				// }
				// ts = append(ts, c)
			} else {
				transactionDef := strings.Split(line, " * ")
				date := transactionDef[0]
				//payee := strings.SplitAfter(transactionDef[1], ";")[0]
				var payee, commentInLine string
				if i := strings.Index(transactionDef[1], ";"); i >= 0 {
					payee, commentInLine = transactionDef[1][:i], transactionDef[1][i:]
				} else {
					payee = transactionDef[1]
				}
				//var commentInLine string
				//if len(strings.SplitAfter(transactionDef[1], ";")) > 1 {
				//	commentInLine = strings.SplitAfter(transactionDef[1], ";")[1]
				//}
				//comment := commentregex.FindString(line)
				t.Date = date
				t.Payee = payee
				t.Postings = ps
				//t.Comment = comment
				t.PayeeComment = commentInLine
			}
		} else if (len(line) > 0) && (line[0:4] == "    ") {
			if strings.HasPrefix(strings.TrimSpace(line), ";") {
				name := strings.TrimSpace(line)
				p := models.Posting {
					Account: models.Account{
						Name: name,
					},
					IsComment: true,
				}
				ps = append(ps, p)
			} else if !strings.HasPrefix(strings.TrimSpace(line), ";") {
				name := accountnameregex.FindString(line)
				amount := amountregex.FindString(line)
				amount = strings.Trim(amount, " ")
				comment := commentregex.FindString(line)
				p := models.Posting {
					Account: models.Account{
						Name: name,
					},
					Amount: amount,
					Comment: comment,
				}
				ps = append(ps, p)
			}
		}
		t.Postings = ps
	}
	if t.Date == "" && t.Payee == "" { // TODO is this needed here and at the top?
		t.IsComment = true
	}
	ts = append(ts, t)
	return &ts, nil
}
