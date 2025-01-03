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
import useTransactionStore from '@/store/transaction/transaction';
import useModalTransaction from '@/store/transaction/modal';

export function EditTransaction() {
  const {
    editTransactionId,
    isModalVisibleEdit,
    showModalEdit,
    hideModalEdit,
  } = useModalTransaction();

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
    setFormErrors((prev) => ({ ...prev, [field]: '' })); 
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

    setFormData({
      card_number: '',
      amount: '',
      payment_method: '',
      merchant_id: '',
      transaction_time: '',
    });
    hideModalEdit()
  };

  return (
    <Dialog open={isModalVisibleEdit} onOpenChange={(open) => (open ? showModalEdit(editTransactionId!) : hideModalEdit())}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm" onClick={() => showModalEdit(editTransactionId!)}>
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
          <Button variant="outline" onClick={hideModalEdit}>
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
