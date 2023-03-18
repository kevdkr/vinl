import Transaction from '../models/Transaction'
import TransactionItem from './Transaction'
import React, { useState, useEffect } from 'react';

import List from '@mui/material/List';
import { Paper } from '@mui/material';
import FormDialog2 from "./AddTransactionFormDialog";

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
        </div>
    )
};

export default Transactions;
