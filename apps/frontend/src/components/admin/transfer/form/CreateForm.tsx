import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

export default function CreateTransferForm({
  formData = {
    transfer_from: '',
    transfer_to: '',
    transfer_amount: '',
    transfer_time: '',
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
      {/* Transfer From */}
      <div>
        <Label htmlFor="transfer_from" className="text-sm font-medium">
          Transfer From
        </Label>
        <Input
          id="transfer_from"
          type="text"
          placeholder="Enter account or source"
          className="mt-1"
          value={formData?.transfer_from || ''}
          onChange={(e) => onFormChange('transfer_from', e.target.value)}
        />
        {formErrors.transfer_from && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.transfer_from}
          </p>
        )}
      </div>

      {/* Transfer To */}
      <div>
        <Label htmlFor="transfer_to" className="text-sm font-medium">
          Transfer To
        </Label>
        <Input
          id="transfer_to"
          type="text"
          placeholder="Enter recipient account"
          className="mt-1"
          value={formData?.transfer_to || ''}
          onChange={(e) => onFormChange('transfer_to', e.target.value)}
        />
        {formErrors.transfer_to && (
          <p className="text-red-500 text-sm mt-1">{formErrors.transfer_to}</p>
        )}
      </div>

      {/* Transfer Amount */}
      <div>
        <Label htmlFor="transfer_amount" className="text-sm font-medium">
          Transfer Amount
        </Label>
        <Input
          id="transfer_amount"
          type="number"
          placeholder="Enter transfer amount"
          className="mt-1"
          value={formData?.transfer_amount || ''}
          onChange={(e) => onFormChange('transfer_amount', e.target.value)}
        />
        {formErrors.transfer_amount && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.transfer_amount}
          </p>
        )}
      </div>

      {/* Transfer Time */}
      <div>
        <Label htmlFor="transfer_time" className="text-sm font-medium">
          Transfer Time
        </Label>
        <Input
          id="transfer_time"
          type="datetime-local"
          className="mt-1"
          value={formData?.transfer_time || ''}
          onChange={(e) => onFormChange('transfer_time', e.target.value)}
        />
        {formErrors.transfer_time && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.transfer_time}
          </p>
        )}
      </div>
    </div>
  );
}
