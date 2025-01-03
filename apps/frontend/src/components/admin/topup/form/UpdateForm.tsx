import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

const topupMethodOptions = [
  { value: 'credit_card', label: 'Credit Card' },
  { value: 'bank_transfer', label: 'Bank Transfer' },
  { value: 'e_wallet', label: 'E-Wallet' },
  { value: 'cash', label: 'Cash' },
];

export default function UpdateTopupForm({
  formData = {
    card_number: '',
    topup_no: '',
    topup_amount: '',
    topup_method: '',
    topup_time: '',
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
      <div>
        <Label htmlFor="topup_no" className="text-sm font-medium">
          Top-up Number
        </Label>
        <Input
          id="topup_no"
          type="text"
          placeholder="Enter top-up number"
          className="mt-1"
          value={formData?.topup_no || ''}
          onChange={(e) => onFormChange('topup_no', e.target.value)}
        />
        {formErrors.topup_no && (
          <p className="text-red-500 text-sm mt-1">{formErrors.topup_no}</p>
        )}
      </div>
      <div>
        <Label htmlFor="topup_amount" className="text-sm font-medium">
          Top-up Amount
        </Label>
        <Input
          id="topup_amount"
          type="number"
          placeholder="Enter top-up amount"
          className="mt-1"
          value={formData?.topup_amount || ''}
          onChange={(e) => onFormChange('topup_amount', e.target.value)}
        />
        {formErrors.topup_amount && (
          <p className="text-red-500 text-sm mt-1">{formErrors.topup_amount}</p>
        )}
      </div>
      <div>
        <Label htmlFor="topup_method" className="text-sm font-medium">
          Top-up Method
        </Label>
        <select
          id="topup_method"
          className="w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:border-blue-500"
          value={formData?.topup_method || ''}
          onChange={(e) => onFormChange('topup_method', e.target.value)}
        >
          <option value="">Select top-up method</option>
          {topupMethodOptions.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
        {formErrors.topup_method && (
          <p className="text-red-500 text-sm mt-1">{formErrors.topup_method}</p>
        )}
      </div>
      <div>
        <Label htmlFor="topup_time" className="text-sm font-medium">
          Top-up Time
        </Label>
        <Input
          id="topup_time"
          type="datetime-local"
          className="mt-1"
          value={formData?.topup_time || ''}
          onChange={(e) => onFormChange('topup_time', e.target.value)}
        />
        {formErrors.topup_time && (
          <p className="text-red-500 text-sm mt-1">{formErrors.topup_time}</p>
        )}
      </div>
    </div>
  );
}
