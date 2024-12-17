import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogFooter,
  } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
  

export default function CheckoutModal({
  isCheckoutModalOpen,
  setIsCheckoutModalOpen,
  setCart,
  total,
  discount,
}: any) {
  if (!isCheckoutModalOpen) return null;

  return (
    <Dialog open={isCheckoutModalOpen} onOpenChange={setIsCheckoutModalOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Confirm Checkout</DialogTitle>
          <DialogDescription>
            You are about to pay a total of Rp.{total.toLocaleString()} after a
            discount of Rp.
            {discount.toLocaleString()}.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button
            variant="destructive"
            onClick={() => setIsCheckoutModalOpen(false)}
          >
            Cancel
          </Button>
          <Button
            variant="default"
            className="shadow-sm"
            onClick={() => {
              
              setIsCheckoutModalOpen(false);
              setCart([]); 
            }}
          >
            Confirm and Pay
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
