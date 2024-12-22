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
import CreateTransactionForm from '../form/CreateForm';

export function AddTransaction() {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    card_number: '',
    amount: '',
    payment_method: '',
    merchant_id: '',
    transaction_time: '',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' })); // Clear error for field
  };

  const handleSubmit = () => {
    const errors: Record<string, string> = {};
    if (!formData.card_number) errors.card_number = 'Card number is required.';
    if (!formData.amount || isNaN(Number(formData.amount))) {
      errors.amount = 'Amount must be a valid number.';
    }
    if (!formData.payment_method)
      errors.payment_method = 'Payment method is required.';
    if (!formData.merchant_id) errors.merchant_id = 'Merchant ID is required.';
    if (!formData.transaction_time)
      errors.transaction_time = 'Transaction time is required.';

    if (Object.keys(errors).length > 0) {
      setFormErrors(errors);
      return;
    }

    console.log('Submitted Data:', formData);

    // Reset form
    setFormData({
      card_number: '',
      amount: '',
      payment_method: '',
      merchant_id: '',
      transaction_time: '',
    });
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          <Plus className="mr-2 h-4 w-4" />
          Add Transaction
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Add New Transaction</DialogTitle>
        </DialogHeader>
        <CreateTransactionForm
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
