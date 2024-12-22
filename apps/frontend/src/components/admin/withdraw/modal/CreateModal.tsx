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
import CreateWithdrawForm from '../form/CreateForm';

export function AddWithdraw() {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    card_number: '',
    withdraw_amount: 0,
    withdraw_time: '',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' }));
  };

  const handleSubmit = () => {
    const errors: Record<string, string> = {};
    if (!formData.card_number) errors.card_number = 'Card number is required.';
    if (!formData.withdraw_amount)
      errors.withdraw_amount = 'Withdraw amount is required.';
    if (!formData.withdraw_time)
      errors.withdraw_time = 'Withdraw time is required.';

    if (Object.keys(errors).length > 0) {
      setFormErrors(errors);
      return;
    }

    console.log('Submitted Data:', formData);

    // Reset form
    setFormData({ card_number: '', withdraw_amount: 0, withdraw_time: '' });
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          <Plus className="mr-2 h-4 w-4" />
          Add Withdraw
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Add New Withdraw</DialogTitle>
        </DialogHeader>
        <CreateWithdrawForm
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
