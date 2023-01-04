import Account from './Account'

export default class Transaction {
    id: string;
    date: string;
    payee: string;
    payeeComment: string;
    comment: string;
    accounts: Account[];
    isComment: boolean;

    constructor(
        id: string,
        date: string,
        payee: string,
        payeeComment: string,
        comment: string,
        accounts: Account[],
        isComment: boolean
    ) {
        this.id = id;
        this.date = date;
        this.payee = payee;
        this.payeeComment = payeeComment;
        this.comment = comment;
        this.accounts = accounts;
        this.isComment = isComment;
    }
}

