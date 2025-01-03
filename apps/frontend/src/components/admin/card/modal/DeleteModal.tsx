import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
} from '@/components/ui/dialog';
import { Trash } from 'lucide-react';
import useModalCard from '@/store/card/modal';

export function DeleteCard() {
  const {
    deleteCardId,
    isModalVisibleDelete,
    showModalDelete,
    hideModalDelete,
  } = useModalCard();

  const handleDelete = () => {
    
    hideModalDelete();
  };

  return (
    <Dialog open={isModalVisibleDelete} onOpenChange={(open) => (open ? showModalDelete(deleteCardId!) : hideModalDelete())}>
      <DialogTrigger asChild>
        <Button variant="destructive" size="sm" onClick={() => showModalDelete(deleteCardId!)}>
          <Trash className="mr-2 h-4 w-4" />
          Delete Card
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md w-full">
        <DialogHeader>
          <DialogTitle>Delete Card</DialogTitle>
        </DialogHeader>
        <div className="text-sm text-gray-600">
          Are you sure you want to delete this card? This action cannot be undone.
        </div>
        <DialogFooter>
          <Button variant="outline" onClick={hideModalDelete}>
            Cancel
          </Button>
          <Button variant="destructive" onClick={handleDelete}>
            Delete
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}