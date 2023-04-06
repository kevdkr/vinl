import Api from './Api'

//const api:string = 'http://localhost:3000/api/' // TODO extract this from being hard-coded
const saveTransactionsFromFile = (file: File): void => {
    const formData = new FormData();
    formData.append("file", file);
    fetch(Api.url + 'uploadfile', {
      method: 'POST',
      body: formData
    })
      .then((response) => {
        return response
      })
        .catch((err) => console.log(err))
  }

const FileUploadService = {
    saveTransactionsFromFile
};

export default FileUploadService;
