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
import CreateCardForm from '../form/CreateForm';
import useModalCard from '@/store/card/modal';
import { useState } from 'react';

export function AddCard() {
  const {
    isModalVisible,
    showModal,
    hideModal,
  } = useModalCard();

  const [formData, setFormData] = useState({
    cardType: '',
    cardProvider: '',
    expireDate: '',
    cvv: '',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' }));
  };

  const handleSubmit = () => {
    const errors: Record<string, string> = {};
    if (!formData.cardType) errors.cardType = 'Card type is required.';
    if (!formData.cardProvider)
      errors.cardProvider = 'Card provider is required.';
    if (!formData.expireDate) errors.expireDate = 'Expire date is required.';
    if (!formData.cvv) errors.cvv = 'CVV is required.';

    if (Object.keys(errors).length > 0) {
      setFormErrors(errors);
      return;
    }

    console.log('Submitted Data:', formData);

    setFormData({ cardType: '', cardProvider: '', expireDate: '', cvv: '' });
    hideModal();
  };

  return (
    <Dialog open={isModalVisible} onOpenChange={(open) => (open ? showModal() : hideModal())}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm" onClick={showModal}>
          <Plus className="mr-2 h-4 w-4" />
          Add Card
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl w-full">
        <DialogHeader>
          <DialogTitle>Add New Card</DialogTitle>
        </DialogHeader>
        <CreateCardForm
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