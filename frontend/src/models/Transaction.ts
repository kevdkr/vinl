import Posting from './Posting'

export default class Transaction {
    id: string;
    date: string;
    payee: string;
    payeeComment: string;
    comment: string;
    postings: Posting[];
    isComment: boolean;

    constructor(
        id: string,
        date: string,
        payee: string,
        payeeComment: string,
        comment: string,
        postings: Posting[],
        isComment: boolean
    ) {
        this.id = id;
        this.date = date;
        this.payee = payee;
        this.payeeComment = payeeComment;
        this.comment = comment;
        this.postings = postings;
        this.isComment = isComment;
    }
}

