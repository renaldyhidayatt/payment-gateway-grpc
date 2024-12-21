export type Transaction = {
  transaction_id: number;
  card_number: string;
  amount: number;
  payment_method: 'credit_card' | 'bank_transfer' | 'e_wallet' | 'cash';
  merchant_id: number;
  transaction_time: string;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
};
