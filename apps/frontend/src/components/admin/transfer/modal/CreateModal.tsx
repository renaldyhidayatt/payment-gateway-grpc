import { useState } from 'react';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
} from '@/components/ui/dialog';
import { Plus } from 'lucide-react';
import CreateTransferForm from '../form/CreateForm';

export function AddTransfer() {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    transfer_from: '',
    transfer_to: '',
    transfer_amount: '',
    transfer_time: '',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' })); // Clear error for field
  };

  const handleSubmit = () => {
    const errors: Record<string, string> = {};
    if (!formData.transfer_from)
      errors.transfer_from = 'Transfer from is required.';
    if (!formData.transfer_to) errors.transfer_to = 'Transfer to is required.';
    if (!formData.transfer_amount || isNaN(Number(formData.transfer_amount))) {
      errors.transfer_amount = 'Amount must be a valid number.';
    }
    if (!formData.transfer_time)
      errors.transfer_time = 'Transfer time is required.';

    if (Object.keys(errors).length > 0) {
      setFormErrors(errors);
      return;
    }

    console.log('Submitted Data:', formData);

    // Reset form
    setFormData({
      transfer_from: '',
      transfer_to: '',
      transfer_amount: '',
      transfer_time: '',
    });
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          <Plus className="mr-2 h-4 w-4" />
          Add Transfer
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Add New Transfer</DialogTitle>
        </DialogHeader>
        <CreateTransferForm
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
