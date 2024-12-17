import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogFooter,
  } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";

export default function DeleteCartModal({
    isDeleteModalOpen,
    setIsDeleteModalOpen,
    removeFromCart,
    selectedItemId,
  }: any) {
    if (!isDeleteModalOpen) return null;
  
    return (
      <Dialog open={isDeleteModalOpen} onOpenChange={setIsDeleteModalOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Confirm Delete</DialogTitle>
            <DialogDescription>
              Are you sure you want to remove this item from your cart?
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button
              variant="destructive"
              onClick={() => setIsDeleteModalOpen(false)}
            >
              Cancel
            </Button>
            <Button
              variant="default"
              onClick={() => {
                if (selectedItemId !== null) {
                  removeFromCart(selectedItemId); // Remove the selected item
                }
              }}
            >
              Delete
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    );
  };