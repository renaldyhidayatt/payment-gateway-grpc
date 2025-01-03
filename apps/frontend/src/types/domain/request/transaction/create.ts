export interface CreateTransaction{
    card_number: string;
    amount: number;
    payment_method: string;
    merchant_id: number;
    transaction_time: Date;
}