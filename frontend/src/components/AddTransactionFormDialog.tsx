import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';

import Fab from '@mui/material/Fab';
import AddIcon from '@mui/icons-material/Add';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import InputAdornment from '@mui/material/InputAdornment';
import OutlinedInput from '@mui/material/OutlinedInput';
import Input from '@mui/material/Input';
import FilledInput from '@mui/material/FilledInput';

const fabStyle = {
  position: 'fixed',
  bottom: 80,
  right: 30,
};

interface Transaction {
  date: string;
  payee: string;
  payeeComment: string;
  comment: string;
  accounts: [
      {
          name: string;
          amount: string;
          comment: string;
          is_comment: boolean;
      }
  ]
  isComment: boolean;
}

export default function FormDialog() {
  const [open, setOpen] = React.useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };


  // const [values, setValues] = React.useState({
  //   date: '',
  //   payee: '',
  //   payeeComment: '',
  //   comment: '',
  //   accounts: [
  //       {
  //           name: '',
  //           amount: '',
  //           comment: '',
  //           is_comment: false,
  //       },
  //   ],
  //   isComment: false,
  // });

  const [transactionValues, setTransactionValues] = React.useState({
    date: '',
    payee: '',
    payeeComment: '',
    comment: '',
    isComment: false,
  });

  const [accountsValues, setAccountsValues] = React.useState({
    accounts: [
      {
        name: '',
        amount: '',
        comment: '',
        is_comment: false,
      },
    ],
  });

  // const handleTransactionValuesChange = (event: React.ChangeEvent<HTMLInputElement>) => {
  //   setTransactionValues({...transactionValues,[event.target.name] : event.target.value});
  // }

  //const handleTransactionValuesChange = (e: any) => setTransactionValues(e.target.value);
  //const handleSubmit = () => {
  //  console.log(transactionValues)
 // }

  const handleTransactionValuesChange = (event: any) => {
    const date = event.target.date;
    //setTransactionValues(values => ({...transactionValues, [name]: value}))
  }
  return (
    <div>
      <Fab sx={fabStyle} color="primary" aria-label="add" onClick={handleClickOpen}>
        <AddIcon />
      </Fab>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Add Transaction</DialogTitle>
        <DialogContent>
          //<FormControl fullWidth sx={{ m: 1 }}>
            <InputLabel htmlFor="outlined-adornment-amount">Date</InputLabel>
            <OutlinedInput
                id="outlined-adornment-amount"
                value={transactionValues.date}
                onChange={handleTransactionValuesChange}
                label="Date"
            />
          //</FormControl>

          //<FormControl fullWidth sx={{ m: 1 }}>
            <InputLabel htmlFor="outlined-adornment-amount">Payee</InputLabel>
            <OutlinedInput
                id="outlined-adornment-amount"
                value={transactionValues.payee}
                onChange={handleTransactionValuesChange}
                label="Payee"
            />
          //</FormControl>
          //<FormControl fullWidth sx={{ m: 1 }}>
            <InputLabel htmlFor="outlined-adornment-amount">Amount</InputLabel>
            <OutlinedInput
                id="outlined-adornment-amount"
                //value={values.amount}
                //onChange={handleChange('amount')}
                startAdornment={<InputAdornment position="start">$</InputAdornment>}
                label="Amount"
          />
        //</FormControl>
        //<FormControl fullWidth sx={{ m: 1 }} variant="filled">
          <InputLabel htmlFor="filled-adornment-amount">Amount</InputLabel>
          <FilledInput
            id="filled-adornment-amount"
            //value={values.amount}
            //onChange={handleChange('amount')}
            startAdornment={<InputAdornment position="start">$</InputAdornment>}
          />
        //</FormControl>
        //<FormControl fullWidth sx={{ m: 1 }} variant="standard">
          <InputLabel htmlFor="standard-adornment-amount">Amount</InputLabel>
          <Input
            id="standard-adornment-amount"
            //value={values.amount}
            //onChange={handleChange('amount')}
            startAdornment={<InputAdornment position="start">$</InputAdornment>}
          />
        //</FormControl>
        </DialogContent>
        <DialogActions>
          //<Button onClick={handleClose}>Cancel</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}
