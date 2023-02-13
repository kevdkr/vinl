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

import Transaction from '../models/Transaction'
import Account from '../models/Account'
import { CardActionArea, Divider, Fab } from '@mui/material';
import Container from '@mui/material/Container';
import EditIcon from '@mui/icons-material/Edit';

type Props = {
  transaction: Transaction;
}

const editButtonStyle = {
    position: 'absolute',
    bottom: 16,
    right: 16,
}

const TransactionItem: React.FC<Props> = ({ transaction }) => (

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
            <li>
            {transaction.accounts ? Object.entries(transaction.accounts).map(account => {
                if (account[1]["is_comment"] === false) {
                    return  (
                        <div>{account[1].name}      {account[1].amount}</div>
                    )
                }
            }): null}
            </li>
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
              console.log("Button clicked");
            }}
        >
            <EditIcon />
        </Fab>
      </CardContent>

        </CardActionArea>

    <Divider />
    </Card>
    // <ListItemButton>
    //     <ListItemText sx={{ fontSize: 14 }}
    //         color="text.secondary"
    //         primary={transaction.date}
    //         secondary={transaction.payee}
    //     >
    //     </ListItemText>
    //     <ListItemText color="text.secondary">
    //       {transaction.payee}
    //     </ListItemText>
    //     <ListItemText sx={{ mb: 1.5 }} color="white">
    //       {transaction.comment}
    //     </ListItemText>
    //     <ListItemText>
    //       account
    //     </ListItemText>
    // </ListItemButton>
);

export default TransactionItem;
