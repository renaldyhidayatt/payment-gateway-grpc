import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Minus, Plus, Trash2 } from 'lucide-react';

export default function PurchaseHistory({
  cart,
  updateQuantity,
  subtotal,
  discount,
  total,
  setIsCheckoutModalOpen,
  setSelectedItemId,
  setIsDeleteModalOpen,
}: any) {
  return (
    <>
      <h2 className="text-xl font-semibold mb-4 text-gray-900 dark:text-gray-100">
        Riwayat Pembelian
      </h2>
      <hr className="border-gray-300 dark:border-gray-700" />
      <br />
      <ScrollArea className="flex-1 mb-4">
        {cart.map((item: any) => (
          <div key={item.id} className="flex justify-between items-center mb-4">
            <div>
              <p className="font-medium text-gray-900 dark:text-gray-100">
                {item.name}
              </p>
              <p className="text-sm text-gray-500 dark:text-gray-400">
                Rp.{item.price.toLocaleString()} x {item.quantity}
              </p>
            </div>
            <div className="flex items-center space-x-2">
              <Button
                variant="outline"
                size="icon"
                className="text-gray-900 dark:text-gray-100 border-gray-300 dark:border-gray-700"
                onClick={() => updateQuantity(item.id, item.quantity - 1)}
              >
                <Minus className="h-4 w-4" />
              </Button>
              <span className="text-gray-900 dark:text-gray-100">
                {item.quantity}
              </span>
              <Button
                variant="outline"
                size="icon"
                className="text-gray-900 dark:text-gray-100 border-gray-300 dark:border-gray-700"
                onClick={() => updateQuantity(item.id, item.quantity + 1)}
              >
                <Plus className="h-4 w-4" />
              </Button>
              <Button
                variant="outline"
                size="icon"
                className="text-red-600 dark:text-red-500 border-gray-300 dark:border-gray-700"
                onClick={() => {
                  setSelectedItemId(item.id);
                  setIsDeleteModalOpen(true);
                }}
              >
                <Trash2 className="h-4 w-4" />
              </Button>
            </div>
          </div>
        ))}
      </ScrollArea>
      <div className="border-t pt-4 mt-4 border-gray-300 dark:border-gray-700">
        <div className="flex justify-between mb-2">
          <span className="text-gray-900 dark:text-gray-100">Subtotal:</span>
          <span className="text-gray-900 dark:text-gray-100">
            Rp.{subtotal.toLocaleString()}
          </span>
        </div>
        <div className="flex justify-between mb-2 text-red-500 dark:text-red-400">
          <span>Discount (10%):</span>
          <span>-Rp.{discount.toLocaleString()}</span>
        </div>
        <div className="flex justify-between font-semibold">
          <span className="text-gray-900 dark:text-gray-100">Total:</span>
          <span className="text-gray-900 dark:text-gray-100">
            Rp.{total.toLocaleString()}
          </span>
        </div>
        <Button
          className="w-full mt-4 text-white dark:text-white"
          variant="default"
          onClick={() => setIsCheckoutModalOpen(true)}
        >
          Checkout
        </Button>
      </div>
    </>
  );
}
