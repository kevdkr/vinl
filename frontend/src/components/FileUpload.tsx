import FileUploadIcon from '@mui/icons-material/FileUpload';
import { DialogContent } from '@mui/material';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import React, { useState } from 'react';
import FileUploadService from "../services/FileUploadService";

export default function FileUpload() {
    const [open, setOpen] = React.useState(false);
    const handleClickOpen = () => {
        setOpen(true);
    };
    const handleClose = () => {
        setOpen(false);
    };

    const [files, setFiles] = useState<File>()

    const handleFileInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { files } = event.target;
        const selectedFiles = files as FileList;
        setFiles(selectedFiles?.[0]);
    };

    const handleSubmit = () => {
        if (!files) return;
        FileUploadService.saveTransactionsFromFile(files);
        handleClose();
    };

    return (
        <div>
            <Button onClick={handleClickOpen}>
                <FileUploadIcon />
            </Button>

            <Dialog open={open} onClose={handleClose}>
                <DialogTitle>Upload a Ledger file</DialogTitle>
                <DialogContent>

                    <input type="file" onChange={handleFileInput} />
                    <Button onClick={handleSubmit}>Submit</Button>
                </DialogContent>
            </Dialog>
        </div>
    );
}
