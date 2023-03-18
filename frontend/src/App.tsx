import Navbar from './Navbar'
import TransactionItem from './components/Transaction'
import Transaction from './models/Transaction'
import Transactions from './components/Transactions'
import Account from './models/Account'

import TransactionService from './services/TransactionService'

import Fab from '@mui/material/Fab';
import AddIcon from '@mui/icons-material/Add';
import Paper from '@mui/material/Paper';
import { createTheme, ThemeProvider, styled } from '@mui/material/styles';
import BottomNav from "./components/BottomNav";
import FormDialog from "./components/AddTransactionFormDialog"
import Box from '@mui/material/Box';
import { FixedSizeList, ListChildComponentProps } from 'react-window';
import List from '@mui/material/List';
import React, { useEffect, useState } from 'react'
const darkTheme = createTheme({ palette: { mode: 'dark' } });
const lightTheme = createTheme({ palette: { mode: 'light' } });

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

const api:string = 'http://localhost:3000/api/' // TODO extract this from being hard-coded

export default function App() {
  const [transactions, setTransactions] = useState<Transaction[]>([])

  useEffect(() => {
    getTransactions()
  }, [])

  // const getTransactions = (): void => {
  //   fetch("transactions")
  //   //.then(({ data: { transactions } }: Transaction[] | any) => setTransactions(transactions))
  //     .then(async (response) => {
  //       if (response.ok) {
  //         setTransactions(await response.json());
  //       }
  //     })
  //   .catch((err: Error) => console.log(err))
  // }
  async function getTransactions() {
    const response = await fetch(api + "transactions");
    setTransactions(await response.json());

  }

  const handleSaveTransaction = (formData: FormValues): void => {
    //e.preventDefault()

    fetch(api + 'transactions', {
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

  const handleDeleteTransaction = (id: string): void => {
    fetch(api + 'transactions/' + id, {
      method: 'DELETE',
    })
      .then((response) => {
        getTransactions()
        return response;
      })
      // .then(({ status }) => {
      //   getTransactions()
      // })
      .catch((err) => console.log(err))

  }

  return (
    <ThemeProvider theme={darkTheme}>
      <Navbar />
      <List sx={{ height: '87%', width: '100%', position: 'fixed', bgcolor: 'background.paper', overflow: 'auto' }}>
          {transactions.map(transaction =>
            <TransactionItem
                            key={transaction.id}
                            transaction={transaction}
                            deleteTransaction={handleDeleteTransaction}/>)}
      </List>
      <FormDialog saveTransaction={handleSaveTransaction}/>
      <BottomNav />
    </ThemeProvider>
  );
}
