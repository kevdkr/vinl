package transfer

import (
	"bufio"
	"bytes"
	"database/sql"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"vinl/models"
)

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

func WriteTransactionsToFile(ts models.Transactions) {
	f, err := os.Create("ledger.dat")
	if err != nil {
		log.Printf("%s", err)
	}

	defer f.Close()

	for _, t := range ts {
		if t.IsComment == false {

			f.WriteString(t.Date + " * ")
			f.WriteString(t.Payee + t.PayeeComment + "\n")
			for _, a := range t.Accounts {
				f.WriteString("    " + a.Name + "    ")
				f.WriteString(" " + a.Amount + "  " + a.Comment + "\n")
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
				log.Fatal(err)
			}
		}
	}
}

func TransferTransactionFromFile(buf *bytes.Buffer, db *sql.DB) {
	//file, err := os.Open(path)
	//if err != nil {
	//	log.Printf("%v", err)
	//}
	//defer file.Close()

	content := buf.Bytes()
	reader := bytes.NewReader(content)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	ts, err := parseFile(reader)
	if err != nil {
		log.Printf("%v", err)
	}

	ts.SaveTransactions(db)
}

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
	var as []models.Account
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
			as = []models.Account{}
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
				t.Accounts = as
				//t.Comment = comment
				t.PayeeComment = commentInLine
			}
		} else if (len(line) > 0) && (line[0:4] == "    ") {
			if strings.HasPrefix(strings.TrimSpace(line), ";") {
				name := strings.TrimSpace(line)
				a := models.Account {
					Name: name,
					IsComment: true,
				}
				as = append(as, a)
			} else if !strings.HasPrefix(strings.TrimSpace(line), ";") {
				name := accountnameregex.FindString(line)
				amount := amountregex.FindString(line)
				amount = strings.Trim(amount, " ")
				comment := commentregex.FindString(line)
				a := models.Account {
					Name: name,
					Amount: amount,
					Comment: comment,
				}
				as = append(as, a)
			}
		}
		t.Accounts = as
	}
	if t.Date == "" && t.Payee == "" { // TODO is this needed here and at the top?
		t.IsComment = true
	}
	ts = append(ts, t)
	return &ts, nil
}
