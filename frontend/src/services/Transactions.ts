import Transaction from '../models/Transaction'

const api:string = 'http://localhost:3000/api/' // TODO extract this from being hard-coded
export async function getTransactions(): Promise<Transaction[]> {
    try {

        const response = await fetch(api + "transactions");
        return await response.json();
    } catch(error) {
        return [];
    }
}

export function deleteTransaction(id: string): Promise<Response> {
    return fetch(api + 'transactions/' + id, {
        method: 'DELETE',
    })
}

type FormValues = {
  date: string;
  payee: string;
  payeeComment: string;
  comment: string;
  postings: {
    name: string;
    amount: string;
    comment: string;
    //is_comment: boolean;
  }[];
  isComment: boolean;
};

export async function createTransaction(formData: FormValues): Promise<Response> {
    return fetch(api + 'transactions', {
        method: 'POST',
        body: JSON.stringify(formData)
    })
}
