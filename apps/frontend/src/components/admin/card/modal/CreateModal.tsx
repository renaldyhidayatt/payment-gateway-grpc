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
import { Plus } from 'lucide-react';
import CreateCardForm from '../form/CreateForm';

export function AddCard() {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    cardType: '', // Default value
    cardProvider: '', // Default value
    expireDate: '', // Default value
    cvv: '', // Default value
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' })); // Clear error for field
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

    // Reset form
    setFormData({ cardType: '', cardProvider: '', expireDate: '', cvv: '' });
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
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
