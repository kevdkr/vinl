import React, { useEffect, useState } from 'react';
import Account from '../models/Account';
import AccountItem from './AccountItem';

import List from '@mui/material/List';
import { AccountFormValues, createAccount, deleteAccount, getAccounts } from '../services/Accounts';
import FormAccountDialog from "./AddAccountFormDialog";

const AccountList: React.FC = () => {

    const [accounts, setAccounts] = useState<Account[]>([])

    useEffect(() => {
        getAccounts().then((response) => {
            setAccounts(response);
        })
    }, [])

    const handleDeleteAccount = (id: string): void => {
        deleteAccount(id)
            .then((response) => {
                getAccounts().then((response) => {
                    setAccounts(response);
                })
                return response;
            })
            .catch((err) => console.log(err))
    }

    const handleSaveAccount = (formData: AccountFormValues): void => {
        createAccount(formData)
            .then((response) => {
                getAccounts().then((response) => {
                    setAccounts(response);
                })
                return response;
            })
            .catch((err) => console.log(err))
    }

    return (
        <div>
            <List sx={{ height: '87%', width: '100%', position: 'fixed', bgcolor: 'background.paper', overflow: 'auto' }}>
                {accounts.map(account =>
                    <AccountItem
                        key={account.id}
                        account={account}
                        deleteAccount={handleDeleteAccount} />)}
            </List>
            <FormAccountDialog saveAccount={handleSaveAccount} />
        </div>
    )
};

export default AccountList;
