export interface UpdateTransaction{
    card_number: string;
    amount: number;
    payment_method: string;
    merchant_id: number;
    transaction_time: Date;
}