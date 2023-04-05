import Navbar from './Navbar'
import Transactions from './components/Transactions'
import { createTheme, ThemeProvider, styled } from '@mui/material/styles';
import BottomNav from "./components/BottomNav";

const darkTheme = createTheme({ palette: { mode: 'dark' } });
const lightTheme = createTheme({ palette: { mode: 'light' } });

export default function App() {
  return (
    <ThemeProvider theme={darkTheme}>
      <Navbar />
      <Transactions />
      <BottomNav />
    </ThemeProvider>
  );
}
