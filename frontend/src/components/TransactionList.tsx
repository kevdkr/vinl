import Transaction from '../models/Transaction'
import TransactionItem from './TransactionItem'
import React, { useState, useEffect } from 'react';

import List from '@mui/material/List';
import { getTransactions, createTransaction, deleteTransaction, TransactionFormValues } from '../services/Transactions'
import FormDialog from "./AddTransactionFormDialog"

const TransactionList: React.FC = () => {

    const [transactions, setTransactions] = useState<Transaction[]>([])

    useEffect(() => {
        getTransactions().then((response) => {
            setTransactions(response);
        })
    }, [])

    const handleDeleteTransaction = (id: string): void => {
        deleteTransaction(id)
            .then((response) => {
                getTransactions().then((response) => {
                    setTransactions(response);
                })
                return response;
            })
            .catch((err) => console.log(err))
    }

    //const api:string = 'http://localhost:3000/api/'
    /* const handleSaveTransaction = (formData: FormValues): void => { // TODO move transactions state variable into TransactionList component (parent) and pass down as props to Update/Delete, etc components
     *     fetch(api + 'transactions', {
     *         method: 'POST',
     *         body: JSON.stringify(formData)
     *     })
     *     .then((response) => {
     *         getTransactions().then((response) => {
     *             setTransactions(response);
     *         })
     *         return response;
     *     })
     *     .catch((err) => console.log(err))
     * } */
    const handleSaveTransaction = (formData: TransactionFormValues): void => {
        createTransaction(formData)
            .then((response) => {
                getTransactions().then((response) => {
                    setTransactions(response);
                })
                return response;
            })
            .catch((err) => console.log(err))
    }

    return (
        <div>
            <List sx={{ height: '87%', width: '100%', position: 'fixed', bgcolor: 'background.paper', overflow: 'auto' }}>
                {transactions.map(transaction =>
                    <TransactionItem
                        key={transaction.id}
                        transaction={transaction}
                        deleteTransaction={handleDeleteTransaction}/>)}
            </List>
            <FormDialog saveTransaction={handleSaveTransaction}/>
        </div>
    )
};

export default TransactionList;
