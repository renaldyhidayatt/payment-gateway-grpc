export type Topup = {
  topup_id: number;
  card_number: string;
  topup_no: string;
  topup_amount: number;
  topup_method: 'credit_card' | 'bank_transfer' | 'e_wallet' | 'cash';
  topup_time: string;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
};
