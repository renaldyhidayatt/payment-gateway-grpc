import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
} from '@/components/ui/dialog';
import { Pencil } from 'lucide-react';

import useModalCard from '@/store/card/modal';
import { useState } from 'react';
import EditCardForm from '../form/UpdateForm';

export function EditCard() {
  const {
    editCardId,
    isModalVisibleEdit,
    showModalEdit,
    hideModalEdit,
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
    console.log('Editing Card ID:', editCardId);

   
    setFormData({ cardType: '', cardProvider: '', expireDate: '', cvv: '' });
    hideModalEdit(); 
  };

  return (
    <Dialog open={isModalVisibleEdit} onOpenChange={(open) => (open ? showModalEdit(editCardId!) : hideModalEdit())}>
      <DialogTrigger asChild>
        <Button variant="outline" size="sm" onClick={() => showModalEdit(editCardId!)}>
          <Pencil className="mr-2 h-4 w-4" />
          Edit Card
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl w-full">
        <DialogHeader>
          <DialogTitle>Edit Card</DialogTitle>
        </DialogHeader>
        <EditCardForm
          formData={formData}
          onFormChange={handleFormChange}
          formErrors={formErrors}
        />
        <DialogFooter>
          <Button variant="outline" onClick={hideModalEdit}>
            Cancel
          </Button>
          <Button variant="default" onClick={handleSubmit}>
            Save Changes
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}