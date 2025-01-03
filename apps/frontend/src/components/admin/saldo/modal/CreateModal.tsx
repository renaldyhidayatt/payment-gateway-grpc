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
import CreateSaldoForm from '../form/CreateForm';
import useModalSaldo from '@/store/saldo/modal';

export function AddSaldo() {
  const {
    isModalVisible,
    showModal,
    hideModal,
  } = useModalSaldo();

  const [formData, setFormData] = useState({
    card_number: '',
    total_balance: '',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' }));
  };

  const handleSubmit = () => {
    const errors: Record<string, string> = {};
    if (!formData.card_number) errors.card_number = 'Card number is required.';
    if (!formData.total_balance || isNaN(Number(formData.total_balance))) {
      errors.total_balance = 'Total balance must be a valid number.';
    }

    if (Object.keys(errors).length > 0) {
      setFormErrors(errors);
      return;
    }

    console.log('Submitted Data:', formData);

  
    setFormData({ card_number: '', total_balance: '' });
    hideModal();
  };

  return (
    <Dialog open={isModalVisible} onOpenChange={(open) => (open ? showModal() : hideModal())}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          Add Card
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Add New Card</DialogTitle>
        </DialogHeader>
        <CreateSaldoForm
          formData={formData}
          onFormChange={handleFormChange}
          formErrors={formErrors}
        />
        <DialogFooter>
          <Button variant="outline" onClick={hideModal}>
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
