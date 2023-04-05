export default class Posting {
    id: string;
    transactionId: string;
    name: string;
    amount: string;
    comment: string;
    is_comment: boolean;

    constructor(
        id: string,
        transactionId: string,
        name: string,
        amount: string,
        comment: string,
        is_comment: boolean
    ) {
        this.id = id;
        this.transactionId = transactionId;
        this.name = name;
        this.amount = amount;
        this.comment = comment;
        this.is_comment = is_comment;
    }
}
