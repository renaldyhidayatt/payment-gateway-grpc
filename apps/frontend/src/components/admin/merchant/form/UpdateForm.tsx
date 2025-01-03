import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

const statusOptions = [
  { value: 'active', label: 'Active' },
  { value: 'inactive', label: 'Inactive' },
  { value: 'pending', label: 'Pending' },
];

export default function UpdateMerchantForm({
  formData = { name: '', api_key: '', user_id: '', status: 'active' },
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
        <Label htmlFor="name" className="text-sm font-medium">
          Name
        </Label>
        <Input
          id="name"
          type="text"
          placeholder="Enter merchant name"
          className="mt-1"
          value={formData?.name || ''}
          onChange={(e) => onFormChange('name', e.target.value)}
        />
        {formErrors.name && (
          <p className="text-red-500 text-sm mt-1">{formErrors.name}</p>
        )}
      </div>
      <div>
        <Label htmlFor="api_key" className="text-sm font-medium">
          API Key
        </Label>
        <Input
          id="api_key"
          type="text"
          placeholder="Enter API key"
          className="mt-1"
          value={formData?.api_key || ''}
          onChange={(e) => onFormChange('api_key', e.target.value)}
        />
        {formErrors.api_key && (
          <p className="text-red-500 text-sm mt-1">{formErrors.api_key}</p>
        )}
      </div>
      <div>
        <Label htmlFor="user_id" className="text-sm font-medium">
          User ID
        </Label>
        <Input
          id="user_id"
          type="number"
          placeholder="Enter user ID"
          className="mt-1"
          value={formData?.user_id || ''}
          onChange={(e) => onFormChange('user_id', e.target.value)}
        />
        {formErrors.user_id && (
          <p className="text-red-500 text-sm mt-1">{formErrors.user_id}</p>
        )}
      </div>

      {/* Status */}
      <div>
        <Label htmlFor="status" className="text-sm font-medium">
          Status
        </Label>
        <select
          id="status"
          className="w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:border-blue-500"
          value={formData?.status || 'active'}
          onChange={(e) => onFormChange('status', e.target.value)}
        >
          {statusOptions.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
      </div>
    </div>
  );
}
