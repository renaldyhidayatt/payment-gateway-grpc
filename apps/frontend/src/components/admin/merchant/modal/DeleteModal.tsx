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
import { Trash } from 'lucide-react';


export function DeleteMerchant() {
  const {
    deleteMerchantId,
    isModalVisibleDelete,
    showModalDelete,
    hideModalDelete,
  } = useModalMerchant();

  const handleDelete = () => {
    
    hideModalDelete();
  };

  return (
    <Dialog open={isModalVisibleDelete} onOpenChange={(open) => (open ? showModalDelete(deleteMerchantId!) : hideModalDelete())}>
      <DialogTrigger asChild>
        <Button variant="destructive" size="sm" onClick={() => showModalDelete(deleteMerchantId!)}>
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