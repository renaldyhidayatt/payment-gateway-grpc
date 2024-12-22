export type Merchant = {
  merchant_id: number;
  name: string;
  api_key: string;
  user_id: number;
  status: 'active' | 'inactive' | 'pending';
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
};
