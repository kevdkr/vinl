import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import ReceiptIcon from '@mui/icons-material/Receipt';
import { Paper } from '@mui/material';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import { Link, useNavigate } from 'react-router-dom';
import * as React from 'react';

export default function BottomNav() {
    const [value, setValue] = React.useState("");
    const navigate=useNavigate();

    return (

        <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
            <BottomNavigation
                showLabels
                value={value}
                onChange={(event, newValue) => {
                    setValue(newValue);
                }}
            >
                <BottomNavigationAction label="Transactions" value={value} icon={<ReceiptIcon />} onClick={() => navigate("/transactions")} />
                <BottomNavigationAction label="Accounts" value={value} icon={<AccountBalanceIcon />} onClick={() => navigate("/accounts")} />
            </BottomNavigation>
        </Paper>
    );
}
