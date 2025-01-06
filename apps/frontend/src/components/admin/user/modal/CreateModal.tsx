import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import CreateUserForm from "../form/CreateForm";
import useModalUser from "@/store/user/modal";
import { CreateUserFormValues } from "@/schemas";
import { useRef } from "react";

export function AddUser() {
  const { isModalVisible, showModal, hideModal } = useModalUser();
  const formRef = useRef<HTMLFormElement>(null);

  const handleSubmit = (data: CreateUserFormValues) => {
    alert(`Submitted Data: ${JSON.stringify(data, null, 2)}`);
    hideModal();
  };

  const handleButtonSubmit = () => {
    formRef.current?.requestSubmit();
  };

  return (
    <Dialog
      open={isModalVisible}
      onOpenChange={(open) => (open ? showModal() : hideModal())}
    >
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Add User</DialogTitle>
        </DialogHeader>
        <CreateUserForm onSubmit={handleSubmit} ref={formRef} />
        <DialogFooter>
          <Button variant="outline" onClick={hideModal}>
            Cancel
          </Button>
          <Button variant="default" onClick={handleButtonSubmit}>
            Submit
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
