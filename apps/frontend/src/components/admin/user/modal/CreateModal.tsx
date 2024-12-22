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
import CreateUserForm from '../form/CreateForm';

export function AddUser() {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    firstname: '',
    lastname: '',
    email: '',
    password: '',
    confirm_password: '',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const handleFormChange = (field: string, value: any) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    setFormErrors((prev) => ({ ...prev, [field]: '' }));
  };

  const handleSubmit = () => {
    const errors: Record<string, string> = {};
    if (!formData.firstname) errors.firstname = 'First name is required.';
    if (!formData.lastname) errors.lastname = 'Last name is required.';
    if (!formData.email) errors.email = 'Email is required.';
    if (!formData.password) errors.password = 'Password is required.';
    if (formData.password !== formData.confirm_password)
      errors.confirm_password = 'Password and confirm password must match.';

    if (Object.keys(errors).length > 0) {
      setFormErrors(errors);
      return;
    }

    console.log('Submitted Data:', formData);

    // Reset form
    setFormData({
      firstname: '',
      lastname: '',
      email: '',
      password: '',
      confirm_password: '',
    });
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          <Plus className="mr-2 h-4 w-4" />
          Add User
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Add New User</DialogTitle>
        </DialogHeader>
        <CreateUserForm
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
