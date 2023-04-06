import Account from '../models/Account'
import Api from './Api'

export async function getAccounts(): Promise<Account[]> {
    try {
        const response = await fetch(Api.url + "accounts");
        return await response.json();
    } catch(error) {
        return [];
    }
}

export function deleteAccount(id: string): Promise<Response> {
    return fetch(Api.url + 'accounts/' + id, {
        method: 'DELETE',
    })
}

export type AccountFormValues = {
  name: string;
};

export async function createAccount(formData: AccountFormValues): Promise<Response> {
    const response = await fetch(Api.url +'accounts', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(formData)
    })
    return response;
}
