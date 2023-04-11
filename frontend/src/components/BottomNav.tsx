import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import ReceiptIcon from '@mui/icons-material/Receipt';
import { Paper } from '@mui/material';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import * as React from 'react';

export default function BottomNav() {
    const [value, setValue] = React.useState("");

    return (

        <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
            <BottomNavigation
                showLabels
                value={value}
                onChange={(event, newValue) => {
                    setValue(newValue);
                }}
            >
                <BottomNavigationAction label="Transactions" icon={<ReceiptIcon />} />
                <BottomNavigationAction label="Accounts" icon={<AccountBalanceIcon />} />
            </BottomNavigation>
        </Paper>
    );
}
