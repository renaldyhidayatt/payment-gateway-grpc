import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

const cardTypeOptions = [
  { value: 'credit', label: 'Credit Card' },
  { value: 'debit', label: 'Debit Card' },
];

const cardProviderOptions = [
  { value: 'visa', label: 'Visa' },
  { value: 'mastercard', label: 'MasterCard' },
  { value: 'amex', label: 'American Express' },
];

export default function CreateCardForm({
  formData = { cardType: '', cardProvider: '', expireDate: '', cvv: '' },
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
        <Label htmlFor="cardType" className="text-sm font-medium">
          Card Type
        </Label>
        <select
          id="cardType"
          className="w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:border-blue-500"
          value={formData?.cardType || ''} 
          onChange={(e) => onFormChange('cardType', e.target.value)}
        >
          <option value="">Select card type</option>
          {cardTypeOptions.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
        {formErrors.cardType && (
          <p className="text-red-500 text-sm mt-1">{formErrors.cardType}</p>
        )}
      </div>
      <div>
        <Label htmlFor="cardProvider" className="text-sm font-medium">
          Card Provider
        </Label>
        <select
          id="cardProvider"
          className="w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:border-blue-500"
          value={formData?.cardProvider || ''}
          onChange={(e) => onFormChange('cardProvider', e.target.value)}
        >
          <option value="">Select card provider</option>
          {cardProviderOptions.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
        {formErrors.cardProvider && (
          <p className="text-red-500 text-sm mt-1">{formErrors.cardProvider}</p>
        )}
      </div>
      <div>
        <Label htmlFor="expireDate" className="text-sm font-medium">
          Expire Date
        </Label>
        <Input
          id="expireDate"
          type="text"
          placeholder="MM/YY"
          className="mt-1"
          value={formData?.expireDate || ''} 
          onChange={(e) => onFormChange('expireDate', e.target.value)}
        />
        {formErrors.expireDate && (
          <p className="text-red-500 text-sm mt-1">{formErrors.expireDate}</p>
        )}
      </div>

    
      <div>
        <Label htmlFor="cvv" className="text-sm font-medium">
          CVV
        </Label>
        <Input
          id="cvv"
          type="password"
          placeholder="Enter CVV"
          className="mt-1"
          value={formData?.cvv || ''} 
          onChange={(e) => onFormChange('cvv', e.target.value)}
        />
        {formErrors.cvv && (
          <p className="text-red-500 text-sm mt-1">{formErrors.cvv}</p>
        )}
      </div>
    </div>
  );
}
