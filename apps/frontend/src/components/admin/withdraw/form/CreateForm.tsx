import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

export default function CreateWithdrawForm({
  formData = { card_number: '', withdraw_amount: 0, withdraw_time: '' },
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

      {/* Withdraw Amount */}
      <div>
        <Label htmlFor="withdraw_amount" className="text-sm font-medium">
          Withdraw Amount
        </Label>
        <Input
          id="withdraw_amount"
          type="number"
          placeholder="Enter withdraw amount"
          className="mt-1"
          value={formData?.withdraw_amount || 0}
          onChange={(e) =>
            onFormChange('withdraw_amount', parseFloat(e.target.value))
          }
        />
        {formErrors.withdraw_amount && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.withdraw_amount}
          </p>
        )}
      </div>

      {/* Withdraw Time */}
      <div>
        <Label htmlFor="withdraw_time" className="text-sm font-medium">
          Withdraw Time
        </Label>
        <Input
          id="withdraw_time"
          type="datetime-local"
          className="mt-1"
          value={formData?.withdraw_time || ''}
          onChange={(e) => onFormChange('withdraw_time', e.target.value)}
        />
        {formErrors.withdraw_time && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.withdraw_time}
          </p>
        )}
      </div>
    </div>
  );
}
