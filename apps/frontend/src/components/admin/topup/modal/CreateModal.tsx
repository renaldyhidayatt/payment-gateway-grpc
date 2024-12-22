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
import CreateTopupForm from '../form/CreateForm';

export function AddTopup() {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    card_number: '',
    topup_no: '',
    topup_amount: '',
    topup_method: '',
    topup_time: '',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' })); // Clear error for field
  };

  const handleSubmit = () => {
    const errors: Record<string, string> = {};
    if (!formData.card_number) errors.card_number = 'Card number is required.';
    if (!formData.topup_no) errors.topup_no = 'Top-up number is required.';
    if (!formData.topup_amount || isNaN(Number(formData.topup_amount))) {
      errors.topup_amount = 'Top-up amount must be a valid number.';
    }
    if (!formData.topup_method)
      errors.topup_method = 'Top-up method is required.';
    if (!formData.topup_time) errors.topup_time = 'Top-up time is required.';

    if (Object.keys(errors).length > 0) {
      setFormErrors(errors);
      return;
    }

    console.log('Submitted Data:', formData);

    // Reset form
    setFormData({
      card_number: '',
      topup_no: '',
      topup_amount: '',
      topup_method: '',
      topup_time: '',
    });
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          <Plus className="mr-2 h-4 w-4" />
          Add Top-up
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Add New Top-up</DialogTitle>
        </DialogHeader>
        <CreateTopupForm
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
