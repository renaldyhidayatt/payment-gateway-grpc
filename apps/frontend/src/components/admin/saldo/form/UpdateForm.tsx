import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

export default function UpdateSaldoForm({
  formData = { card_number: '', total_balance: '' },
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
        <Label htmlFor="total_balance" className="text-sm font-medium">
          Total Balance
        </Label>
        <Input
          id="total_balance"
          type="number"
          placeholder="Enter total balance"
          className="mt-1"
          value={formData?.total_balance || ''}
          onChange={(e) => onFormChange('total_balance', e.target.value)}
        />
        {formErrors.total_balance && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.total_balance}
          </p>
        )}
      </div>
    </div>
  );
}
