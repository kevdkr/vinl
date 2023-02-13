import Navbar from './Navbar'
import TransactionItem from './components/Transaction'
import Transaction from './models/Transaction'
import Transactions from './components/Transactions'
import Account from './models/Account'

import TransactionService from './services/transaction.service'

import Fab from '@mui/material/Fab';
import AddIcon from '@mui/icons-material/Add';
import Paper from '@mui/material/Paper';
import { createTheme, ThemeProvider, styled } from '@mui/material/styles';
import BottomNav from "./components/BottomNav";
import FormDialog2 from "./components/AddTransactionFormDialog2"

import List from '@mui/material/List';
import React, { useEffect, useState } from 'react'
const darkTheme = createTheme({ palette: { mode: 'dark' } });
const lightTheme = createTheme({ palette: { mode: 'light' } });


// const accountItem = new Account("1", "1", "name", "amount", "", false);
// const accountList = new Array<Account>();
// accountList.push(accountItem);
// const transactionItem = new Transaction("1", "10-22-2022", "payee", "", "", accountList, false);

      // <TransactionItem transaction = {transactionItem}/>
const fabStyle = {
  position: 'fixed',
  bottom: 80,
  right: 30,
};

type FormValues = {
  date: string;
  payee: string;
  payeeComment: string;
  comment: string;
  accounts: {
    name: string;
    amount: string;
    comment: string;
    //is_comment: boolean;
  }[];
  isComment: boolean;
};

export default function App() {
  const [transactions, setTransactions] = useState<Transaction[]>([])

  useEffect(() => {
    getTransactions()
  }, [])

  const getTransactions = (): void => {
    fetch("transactions")
    //.then(({ data: { transactions } }: Transaction[] | any) => setTransactions(transactions))
      .then(async (response) => {
        if (response.ok) {
          setTransactions(await response.json());
        }
      })
    .catch((err: Error) => console.log(err))
  }

  const handleSaveTransaction = (formData: FormValues): void => {
    //e.preventDefault()

    fetch('transactions', {
      method: 'POST',
      body: JSON.stringify(formData)
    })
      .then(({ status }) => {
        // if (status !== 201) {
        //   throw new Error('Error! Transaction not saved')
        // }
        //setTransactions(data.transactions)
        getTransactions()
      })
      .catch((err) => console.log(err))
  }

  return (
    //<Router>
      //
    <ThemeProvider theme={darkTheme}>
      <Navbar />
      //<Transactions />
      <div>
      <List sx={{ flexGrow: 1, height: '100%', width: '100%', position: 'fixed', bgcolor: 'background.paper', overflow: 'auto' }}>
          {transactions.map(transaction => <TransactionItem transaction={transaction} />)}
      </List>
      <FormDialog2 saveTransaction={handleSaveTransaction}/>
      </div>
      <BottomNav />
    </ThemeProvider>
    //</Router>
  );
}
