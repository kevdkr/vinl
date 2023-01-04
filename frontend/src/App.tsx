//import logo from './logo.svg';
//import './App.css';

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

//export default App;

import React from "react";
//import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import Navbar from './Navbar'
import TransactionItem from './components/Transaction'
import Transaction from './models/Transaction'
import Transactions from './components/Transactions'
import Account from './models/Account'

import TransactionService from './services/transaction.service'

import Fab from '@mui/material/Fab';
import AddIcon from '@mui/icons-material/Add';
import Paper from '@mui/material/Paper';
import { createTheme, ThemeProvider, styled } from '@mui/material/styles';
import BottomNav from "./components/BottomNav";
import FormDialog2 from "./components/AddTransactionFormDialog2"

const darkTheme = createTheme({ palette: { mode: 'dark' } });
const lightTheme = createTheme({ palette: { mode: 'light' } });


// const accountItem = new Account("1", "1", "name", "amount", "", false);
// const accountList = new Array<Account>();
// accountList.push(accountItem);
// const transactionItem = new Transaction("1", "10-22-2022", "payee", "", "", accountList, false);

      // <TransactionItem transaction = {transactionItem}/>
const fabStyle = {
  position: 'fixed',
  bottom: 80,
  right: 30,
};

export default function App() {


  return (
    //<Router>
      //
    <ThemeProvider theme={darkTheme}>
      <Navbar />
      <Transactions />
      // <Fab sx={fabStyle} color="primary" aria-label="add">
      //   <AddIcon />
      // </Fab>
      <FormDialog2 />
      <BottomNav />
    </ThemeProvider>
    //</Router>
  );
}
