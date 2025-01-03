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
import useModalMerchant from '@/store/merchant/modal';
import UpdateMerchantForm from '../form/UpdateForm';

export function EditMerchant() {
  const {
    isModalVisibleEdit,
    showModalEdit,
    hideModalEdit,
    editMerchantId
  } = useModalMerchant();

  const [formData, setFormData] = useState({
    name: '',
    api_key: '',
    user_id: '',
    status: 'active',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' }));
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

    
    setFormData({ name: '', api_key: '', user_id: '', status: 'active' });
    hideModalEdit();
    
  };

  return (
    <Dialog open={isModalVisibleEdit} onOpenChange={(open) => (open ? showModalEdit(editMerchantId!) : hideModalEdit())}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          Update Merchant
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl w-full">
        <DialogHeader>
          <DialogTitle>Update Merchant</DialogTitle>
        </DialogHeader>
        <UpdateMerchantForm
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
