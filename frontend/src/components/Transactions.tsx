import Transaction from '../models/Transaction'
import TransactionService from '../services/transaction.service'
import TransactionItem from './Transaction'
import React, { useState, useEffect } from 'react';

import List from '@mui/material/List';
import { Paper } from '@mui/material';

type Props = {
    transactions: Array<Transaction>;
}

const Transactions: React.FC = () => {
    // return (
    //     <div>
    //     {props.transactions.length > 0 ? (
    //         props.transactions.map(i => (
    //             <tr key={i.id}>
    //             <td>{i["date"]}</td>
    //             <td>{i["payee"]}</td>
    //             <td>
    //             </td>
    //           </tr>
    //         ))
    //       ) : (
    //         <tr>
    //           <td colSpan={3}>no users</td>
    //         </tr>
    //       )}
    //     </div>
    // );


// function renderRow(props: ListChildComponentProps) {
//   const { index, style } = props;

//   return (
//     <ListItem style={style} key={index} component="div" disablePadding>
//       <ListItemButton>
//         <ListItemText primary={`Item ${index + 1}`} />
//       </ListItemButton>
//     </ListItem>
//   );
// }

// export default function VirtualizedList() {
//   return (
//     <Box
//       sx={{ width: '100%', height: '100%', bgcolor: 'background.paper' }}
//     >
//       <FixedSizeList
//         height={400}
//         width={360}
//         itemSize={46}
//         itemCount={200}
//         overscanCount={5}
//       >
//         {renderRow}
//       </FixedSizeList>
//     </Box>
//   );
// }

    const [transactions, setTransactions] = useState([]);

    useEffect(() => {
        fetch("transactions").then(async (response) => {
            if (response.ok) {
                setTransactions(await response.json());
            }
        });
    }, []);
    console.log(transactions);
    return (
        <List sx={{ flexGrow: 1, height: '100%', width: '100%', position: 'fixed', bgcolor: 'background.paper', overflow: 'auto' }}>
            {transactions.map(transaction => <TransactionItem transaction={transaction} />)}
        </List>
    )
};

export default Transactions;
