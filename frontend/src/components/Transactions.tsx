import Transaction from '../models/Transaction'
import TransactionItem from './Transaction'
import React, { useState, useEffect } from 'react';

import List from '@mui/material/List';
import { getTransactions, createTransaction, deleteTransaction } from '../services/Transactions'
import FormDialog from "./AddTransactionFormDialog"

const Transactions: React.FC = () => {

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

    type FormValues = {
        date: string;
        payee: string;
        payeeComment: string;
        comment: string;
        postings: {
            name: string;
            amount: string;
            comment: string;
            //is_comment: boolean;
        }[];
        isComment: boolean;
    };

    const api:string = 'http://localhost:3000/api/'
    const handleSaveTransaction = (formData: FormValues): void => { // TODO move transactions state variable into TransactionList component (parent) and pass down as props to Update/Delete, etc components
        fetch(api + 'transactions', {
            method: 'POST',
            body: JSON.stringify(formData)
        })
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

export default Transactions;
