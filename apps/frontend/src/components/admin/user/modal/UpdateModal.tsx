import { useRef } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
} from "@/components/ui/dialog";
import useModalUser from "@/store/user/modal";
import { UpdateUserFormValues } from "@/schemas";
import UpdateUserForm from "../form/UpdateForm";

export function EditUser() {
  const { editUserId, isModalVisibleEdit, showModalEdit, hideModalEdit } =
    useModalUser();
  const formRef = useRef<HTMLFormElement>(null);

  const handleSubmit = (data: UpdateUserFormValues) => {
    alert(`Submitted Data: ${JSON.stringify(data, null, 2)}`);
    hideModalEdit();
  };

  const handleButtonSubmit = () => {
    formRef.current?.requestSubmit();
  };

  return (
    <Dialog
      open={isModalVisibleEdit}
      onOpenChange={(open) =>
        open ? showModalEdit(editUserId!) : hideModalEdit()
      }
    >
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Edit New User</DialogTitle>
        </DialogHeader>
        <UpdateUserForm onSubmit={handleSubmit} ref={formRef} />
        <DialogFooter>
          <Button variant="outline" onClick={hideModalEdit}>
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
