import Transaction from '../models/Transaction'
import TransactionService from '../services/transaction.service'
import TransactionItem from './Transaction'
import React, { useState, useEffect } from 'react';

import List from '@mui/material/List';
import { Paper } from '@mui/material';
import FormDialog2 from "./AddTransactionFormDialog2";

type Props = {
    transactions: Array<Transaction>;
}

const Transactions: React.FC = () => {

    const [transactions, setTransactions] = useState([]);

    useEffect(() => {
        fetch("transactions").then(async (response) => {
            if (response.ok) {
                setTransactions(await response.json());
            }
        });
    }, []);
    //console.log(transactions);
    return (
        <div>
        <List sx={{ flexGrow: 1, height: '100%', width: '100%', position: 'fixed', bgcolor: 'background.paper', overflow: 'auto' }}>
            {transactions.map(transaction => <TransactionItem transaction={transaction} />)}
        </List>
        </div>
    )
};

export default Transactions;
