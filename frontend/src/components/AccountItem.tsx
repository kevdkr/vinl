import * as React from 'react';
import Box from '@mui/material/Box';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import { FixedSizeList, ListChildComponentProps } from 'react-window';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import List from '@mui/material/List';

import Transaction from '../models/Transaction'
import Account from '../models/Account'
import Posting from '../models/Posting'
import { CardActionArea, Divider, Fab } from '@mui/material';
import Container from '@mui/material/Container';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

type Props = {
  account: Account
  deleteAccount: (id: string | any) => void
}

const editButtonStyle = {
    position: 'absolute',
    bottom: 16,
    right: 16,
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
