import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';

import AddIcon from '@mui/icons-material/Add';
import Fab from '@mui/material/Fab';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import OutlinedInput from '@mui/material/OutlinedInput';
import React, { ChangeEvent } from 'react';
import Account from '../models/Account';
import { AccountFormValues } from '../services/Accounts';

type Props = {
    saveAccount: (formData: Account | any) => void
}

const fabStyle = {
    position: 'fixed',
    bottom: 80,
    right: 30,
};

const FormAccountDialog: React.FC<Props> = ({ saveAccount }) => {
    const [open, setOpen] = React.useState(false);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const [name, setName] = React.useState("");

    const onChange = (event: ChangeEvent<HTMLInputElement>) => {
        let value: typeof name[keyof typeof name] = event.target.value;

        setName(value);
    };

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const formData: AccountFormValues = {
            name: name
        }
        saveAccount(formData)
        handleClose();
    }

    return (
        <div>
            <Fab sx={fabStyle} color="primary" aria-label="add" onClick={handleClickOpen}>
                <AddIcon />
            </Fab>
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle>Add Account</DialogTitle>
                <DialogContent>

                    <form id="accountform" onSubmit={handleSubmit}>

                        <div>
                            <FormControl fullWidth sx={{ m: 1 }}>
                                <InputLabel htmlFor="outlined-adornment-date">Name</InputLabel>
                                <OutlinedInput
                                    id="outlined-adornment-date"
                                    label="Name"
                                    onChange={onChange}
                                />
                            </FormControl>
                        </div>
                    </form>
                </DialogContent>
                <DialogActions>
                    <Button type="submit" form="accountform">Submit</Button>
                    <Button onClick={handleClose}>Cancel</Button>
                </DialogActions>
            </Dialog>
        </div>
    );
}

export default FormAccountDialog
