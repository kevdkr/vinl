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
import { CardActionArea, Divider, Fab } from '@mui/material';
import Container from '@mui/material/Container';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

type Props = {
  transaction: Transaction
  deleteTransaction: (id: string | any) => void
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

const TransactionItem: React.FC<Props> = ({ transaction, deleteTransaction }) => {

    const handleClickDelete = (id: string) => {
        deleteTransaction(id);
    }

    return (
    <Card sx={{ minWidth: 275 }}>
        <CardActionArea>
      <CardContent>
        <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
          {transaction.date}
        </Typography>
        <Typography variant="h5" component="div">
          {transaction.payee}
        </Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          {transaction.comment}
        </Typography>
        <Typography>
            <List>
            {transaction.accounts ? Object.entries(transaction.accounts).map(account => {
                if (account[1]["is_comment"] === false) {
                    return  (
                        <div key={account[1].id}>{account[1].name}      {account[1].amount}</div>
                    )
                }
            }): null}
            </List>
        </Typography>

        <Fab
            sx={editButtonStyle}
            size="small"
            color="primary"
            aria-label="edit"
            onMouseDown={event => event.stopPropagation()}
            onClick={event => {
              event.stopPropagation();
              event.preventDefault();
            }}
        >
            <EditIcon />
        </Fab>

        <Fab
            sx={deleteButtonStyle}
            size="small"
            color="primary"
            aria-label="delete"
            onMouseDown={event => event.stopPropagation()}
            onClick={event => {
                event.stopPropagation();
                event.preventDefault();
                handleClickDelete(transaction.id);
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
export default TransactionItem;
