import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

const paymentMethodOptions = [
  { value: 'credit_card', label: 'Credit Card' },
  { value: 'bank_transfer', label: 'Bank Transfer' },
  { value: 'e_wallet', label: 'E-Wallet' },
  { value: 'cash', label: 'Cash' },
];

export default function CreateTransactionForm({
  formData = {
    card_number: '',
    amount: '',
    payment_method: '',
    merchant_id: '',
    transaction_time: '',
  },
  onFormChange,
  formErrors = {},
}: {
  formData: any;
  onFormChange: (field: string, value: any) => void;
  formErrors: Record<string, string>;
}) {
  return (
    <div className="space-y-4">
      {/* Card Number */}
      <div>
        <Label htmlFor="card_number" className="text-sm font-medium">
          Card Number
        </Label>
        <Input
          id="card_number"
          type="text"
          placeholder="Enter card number"
          className="mt-1"
          value={formData?.card_number || ''}
          onChange={(e) => onFormChange('card_number', e.target.value)}
        />
        {formErrors.card_number && (
          <p className="text-red-500 text-sm mt-1">{formErrors.card_number}</p>
        )}
      </div>

      {/* Amount */}
      <div>
        <Label htmlFor="amount" className="text-sm font-medium">
          Amount
        </Label>
        <Input
          id="amount"
          type="number"
          placeholder="Enter amount"
          className="mt-1"
          value={formData?.amount || ''}
          onChange={(e) => onFormChange('amount', e.target.value)}
        />
        {formErrors.amount && (
          <p className="text-red-500 text-sm mt-1">{formErrors.amount}</p>
        )}
      </div>

      {/* Payment Method */}
      <div>
        <Label htmlFor="payment_method" className="text-sm font-medium">
          Payment Method
        </Label>
        <select
          id="payment_method"
          className="w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:border-blue-500"
          value={formData?.payment_method || ''}
          onChange={(e) => onFormChange('payment_method', e.target.value)}
        >
          <option value="">Select payment method</option>
          {paymentMethodOptions.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
        {formErrors.payment_method && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.payment_method}
          </p>
        )}
      </div>

      {/* Merchant ID */}
      <div>
        <Label htmlFor="merchant_id" className="text-sm font-medium">
          Merchant ID
        </Label>
        <Input
          id="merchant_id"
          type="number"
          placeholder="Enter merchant ID"
          className="mt-1"
          value={formData?.merchant_id || ''}
          onChange={(e) => onFormChange('merchant_id', e.target.value)}
        />
        {formErrors.merchant_id && (
          <p className="text-red-500 text-sm mt-1">{formErrors.merchant_id}</p>
        )}
      </div>

      {/* Transaction Time */}
      <div>
        <Label htmlFor="transaction_time" className="text-sm font-medium">
          Transaction Time
        </Label>
        <Input
          id="transaction_time"
          type="datetime-local"
          className="mt-1"
          value={formData?.transaction_time || ''}
          onChange={(e) => onFormChange('transaction_time', e.target.value)}
        />
        {formErrors.transaction_time && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.transaction_time}
          </p>
        )}
      </div>
    </div>
  );
}
