import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import * as React from 'react';

import DeleteIcon from '@mui/icons-material/Delete';
import { CardActionArea, Divider, Fab } from '@mui/material';
import Account from '../models/Account';

type Props = {
    account: Account
    deleteAccount: (id: string | any) => void
}

const deleteButtonStyle = {
    position: 'absolute',
    bottom: 16,
    right: 72,
}

const AccountItem: React.FC<Props> = ({ account, deleteAccount }) => {

    const handleClickDelete = (id: string) => {
        deleteAccount(id);
    }

    return (
        <Card sx={{ minWidth: 275 }}>
            <CardActionArea>
                <CardContent>
                    <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                        {account.name}
                    </Typography>

                    <Fab
                        sx={deleteButtonStyle}
                        size="small"
                        color="primary"
                        aria-label="delete"
                        onMouseDown={event => event.stopPropagation()}
                        onClick={event => {
                            event.stopPropagation();
                            event.preventDefault();
                            handleClickDelete(account.id);
                        }}
                    >
                        <DeleteIcon />
                    </Fab>
                </CardContent>

            </CardActionArea>

            <Divider />
        </Card>
    );
}

export default AccountItem;
