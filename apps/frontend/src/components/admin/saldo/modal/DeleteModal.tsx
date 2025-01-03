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
import useModalSaldo from '@/store/saldo/modal';

export function DeleteSaldo() {
  const {
    deleteSaldoId,
    isModalVisibleDelete,
    showModalDelete,
    hideModalDelete,
  } = useModalSaldo();

  const handleDelete = () => {
    
    hideModalDelete();
  };

  return (
    <Dialog open={isModalVisibleDelete} onOpenChange={(open) => (open ? showModalDelete(deleteSaldoId!) : hideModalDelete())}>
      <DialogTrigger asChild>
        <Button variant="destructive" size="sm" onClick={() => showModalDelete(deleteSaldoId!)}>
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