import { useState } from 'react';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
  DialogClose,
} from '@/components/ui/dialog';
import CreateMerchantForm from '../form/CreateForm';

export function AddMerchant() {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    api_key: '',
    user_id: '',
    status: 'active', // Default value
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' })); // Clear error for field
  };

  const handleSubmit = () => {
    const errors: Record<string, string> = {};
    if (!formData.name) errors.name = 'Name is required.';
    if (!formData.api_key) errors.api_key = 'API Key is required.';
    if (!formData.user_id) errors.user_id = 'User ID is required.';

    if (Object.keys(errors).length > 0) {
      setFormErrors(errors);
      return;
    }

    console.log('Submitted Data:', formData);

    // Reset form
    setFormData({ name: '', api_key: '', user_id: '', status: 'active' });
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          Add Merchant
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl w-full">
        <DialogHeader>
          <DialogTitle>Add New Merchant</DialogTitle>
        </DialogHeader>
        <CreateMerchantForm
          formData={formData}
          onFormChange={handleFormChange}
          formErrors={formErrors}
        />
        <DialogFooter>
          <Button variant="outline" onClick={() => setIsOpen(false)}>
            Cancel
          </Button>
          <Button variant="default" onClick={handleSubmit}>
            Submit
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
