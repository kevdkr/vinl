package transfer

import (
	"bytes"
	"vinl/internal/models"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseFile(t *testing.T) {
	cases := []struct {
		name	 string
		file     []byte
		expected *models.Transactions
	}{
		{
			name: "basic",
			file: []byte(
`2022/10/26 * Transaction 1 At Store
    Expenses:Groceries:Store    $21.16
    Liabilities:CreditCard`),
			expected: &models.Transactions{
				models.Transaction{
					Date:         "2022/10/26",
					Payee:        "Transaction 1 At Store",
					PayeeComment: "",
					Comment:      "",
					Accounts: []models.Account{
						{
							Name:      "Expenses:Groceries:Store",
							Amount:    "$21.16",
							Comment:   "",
							IsComment: false,
						},
						{
							Name:		"Liabilities:CreditCard",
							Amount:		"",
							Comment: 	"",
							IsComment:	false,
						},
					},
					IsComment: false,
				},
			},
		},
		{
			name: "virtual account",
			file: []byte(
`2022/10/27 * Loan
    Liabilities:Loan    $112.21
    (Expenses:Liabilities:LoanItem)        $112.21
    Assets:Checking`),
			expected: &models.Transactions{
				models.Transaction{
					Date:		"2022/10/27",
					Payee:		"Loan",
					PayeeComment: "",
					Comment:	"",
					Accounts: []models.Account{
						{
							Name:	"Liabilities:Loan",
							Amount:	"$112.21",
							Comment: "",
							IsComment: false,
						},
						{
							Name:   "(Expenses:Liabilities:LoanItem)",
							Amount: "$112.21",
							Comment: "",
							IsComment: false,
						},
						{
							Name: "Assets:Checking",
							Amount: "",
							Comment: "",
							IsComment: false,
						},
					},
				},
			},
		},
		{
			name: "account comment in line",
			file: []byte(
`2022/10/28 * Check
    Assets:Bank:Checking                    $1.00        ; put in taxes / 401k / etc
    Income:Salary`),
			expected: &models.Transactions{
				models.Transaction{
					Date:	"2022/10/28",
					Payee:  "Check",
					PayeeComment: "",
					Comment: "",
					Accounts: []models.Account{
						{
							Name: "Assets:Bank:Checking",
							Amount: "$1.00",
							Comment: "; put in taxes / 401k / etc",
							IsComment: false,
						},
						{
							Name: "Income:Salary",
							Amount: "",
							Comment: "",
							IsComment: false,
						},
					},
				},
			},
		},
		{
			name: "account amount expression",
			file: []byte(
`2022/10/29 * Loan
    Liabilities:Loan:Principal           $1.99
    Expenses:Liabilities:Loan:Interest   $5.99
    Assets:UnappliedLoan               $-10.67
    (Expenses:Liabilities:Loan)        ($1.99 + $5.99 - $10.67)
    Assets:Checking                   $-10.68`),
			expected: &models.Transactions{
				models.Transaction{
					Date: "2022/10/29",
					Payee: "Loan",
					PayeeComment: "",
					Comment: "",
					Accounts: []models.Account{
						{
							Name: "Liabilities:Loan:Principal",
							Amount: "$1.99",
							Comment: "",
							IsComment: false,
						},
						{
							Name: "Expenses:Liabilities:Loan:Interest",
							Amount: "$5.99",
							Comment: "",
							IsComment: false,
						},
						{
							Name: "Assets:UnappliedLoan",
							Amount: "$-10.67",
							Comment: "",
							IsComment: false,
						},
						{
							Name: "(Expenses:Liabilities:Loan)",
							Amount: "($1.99 + $5.99 - $10.67)",
							Comment: "",
							IsComment: false,
						},
						{
							Name: "Assets:Checking",
							Amount: "$-10.68",
							Comment: "",
							IsComment: false,
						},
					},
				},
			},
		},
		{
			name: "transactions comments",
			file: []byte(
`; test comment at top
2022/08/16 * Electric Bill  ; test comment in line
    Expenses:Utilities:Electricity    $1.59
	Liabilities:CreditCard
; test ;test2sameline
; test3newline`),
			expected: &models.Transactions{
				models.Transaction{
					Date: "2022/08/16",
					Payee: "Electric Bill  ",
					PayeeComment: "; test comment in line",
					Comment: "; test comment at top\n; test ;test2sameline\n; test3newline\n",
					Accounts: []models.Account{
						{
							Name: "Expenses:Utilities:Electricity",
							Amount: "$1.59",
							Comment: "",
							IsComment: false,
						},
						{
							Name: "Liabilities:CreditCard",
							Amount: "",
							Comment: "",
							IsComment: false,
						},
					},
					IsComment: false,
				},
				models.Transaction{
					Date: "",
					Payee: "",
					PayeeComment: "",
					Comment: 	"; test 4 unrelated\n; test 5 unrelated\n",
					IsComment: true,
				},
			},
		},
	}

	for _, tc := range cases {
		//t.Log("Running test:\n", string(tc.file[:]))
		//file, err := os.Open(tc.file)
		//file := bytes.NewBufferString(tc.file)
		//if err != nil {
		//	t.Log("failed opening test ledger file: ", tc.file)
		//	t.Fail()
		//}
		//defer file.Close()
		//scanner := bufio.NewScanner(file)
		//scanner.Split(bufio.ScanLines)
		reader := bytes.NewReader(tc.file)
		ts, err := parseFile(reader)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		//if ts != tc.expected {
		if !cmp.Equal(ts, tc.expected) {
			t.Logf("Transactions don't match:\nactual: %+v\n\n expected: %+v\n", ts, tc.expected)
			t.Fail()
		}
	}
}
