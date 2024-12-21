import { useEffect, useState } from 'react';
import { Sheet, SheetContent, SheetTrigger } from '@/components/ui/sheet';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { ShoppingCart } from 'lucide-react';
import DeleteCartModal from '@/components/admin/modal/deleteCart';
import CheckoutModal from '@/components/admin/modal/checkout';
import PurchaseHistory from '@/components/admin/point-of-sale/purchaseHistory';
import { products } from '@/helpers/product_data';
import { Product } from '@/types/product';

export default function PointOfSalePage() {
  const [cart, setCart] = useState<
    { id: number; name: string; price: number; quantity: number }[]
  >([]);
  const [isMobile, setIsMobile] = useState(false);
  const [isCheckoutModalOpen, setIsCheckoutModalOpen] = useState(false);
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState(false);
  const [selectedItemId, setSelectedItemId] = useState<number | null>(null);

  useEffect(() => {
    const checkMobile = () => {
      setIsMobile(window.innerWidth < 768);
    };
    checkMobile();
    window.addEventListener('resize', checkMobile);
    return () => window.removeEventListener('resize', checkMobile);
  }, []);

  const addToCart = (product: { id: number; name: string; price: number }) => {
    setCart((prevCart) => {
      const existingItem = prevCart.find((item) => item.id === product.id);
      if (existingItem) {
        return prevCart.map((item) =>
          item.id === product.id
            ? { ...item, quantity: item.quantity + 1 }
            : item
        );
      }
      return [...prevCart, { ...product, quantity: 1 }];
    });
  };

  const removeFromCart = (productId: number) => {
    setCart((prevCart) => prevCart.filter((item) => item.id !== productId));
    setIsDeleteModalOpen(false);
  };

  const updateQuantity = (productId: number, newQuantity: number) => {
    if (newQuantity === 0) {
      setSelectedItemId(productId);
      setIsDeleteModalOpen(true);
    } else {
      setCart((prevCart) =>
        prevCart.map((item) =>
          item.id === productId ? { ...item, quantity: newQuantity } : item
        )
      );
    }
  };

  const subtotal = cart.reduce(
    (sum, item) => sum + item.price * item.quantity,
    0
  );
  const discount = subtotal * 0.1;
  const total = subtotal - discount;

  return (
    <div className="flex h-full overflow-hidden">
      <main className="flex-1 overflow-y-auto">
        <div className="p-6">
          <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {products.map((product: Product) => (
              <Card key={product.id} className="flex flex-col justify-between">
                <CardHeader className="p-0">
                  <img
                    src={product.image}
                    alt={product.name}
                    className="w-full h-32 object-cover"
                  />
                </CardHeader>
                <CardContent className="p-4">
                  <CardTitle className="text-lg">{product.name}</CardTitle>
                  <p className="text-sm text-gray-500">
                    Rp.{product.price.toLocaleString()}
                  </p>
                </CardContent>
                <CardFooter className="p-4 pt-0">
                  <Button
                    className="w-full text-white dark:text-white"
                    onClick={() => addToCart(product)}
                  >
                    Add to Cart
                  </Button>
                </CardFooter>
              </Card>
            ))}
          </div>
        </div>
      </main>

      {isMobile ? (
        <Sheet>
          <SheetTrigger asChild>
            <Button className="fixed bottom-16 right-4 z-50" size="icon">
              <ShoppingCart />
            </Button>
          </SheetTrigger>
          <SheetContent className="w-full sm:w-[400px] sm:max-w-full">
            <div className="h-full overflow-y-auto">
              <PurchaseHistory
                cart={cart}
                updateQuantity={updateQuantity}
                removeFromCart={removeFromCart}
                subtotal={subtotal}
                discount={discount}
                total={total}
                setIsCheckoutModalOpen={setIsCheckoutModalOpen}
                setSelectedItemId={setSelectedItemId}
                setIsDeleteModalOpen={setIsDeleteModalOpen}
              />
            </div>
          </SheetContent>
        </Sheet>
      ) : (
        <aside className="w-80 bg-white dark:bg-[#0A0A0A] overflow-y-auto flex-shrink-0 border-l border-gray-200 dark:border-gray-700">
          <div className="p-6">
            <PurchaseHistory
              cart={cart}
              updateQuantity={updateQuantity}
              removeFromCart={removeFromCart}
              subtotal={subtotal}
              discount={discount}
              total={total}
              setIsCheckoutModalOpen={setIsCheckoutModalOpen}
              setSelectedItemId={setSelectedItemId}
              setIsDeleteModalOpen={setIsDeleteModalOpen}
            />
          </div>
        </aside>
      )}

      <CheckoutModal
        isCheckoutModalOpen={isCheckoutModalOpen}
        setIsCheckoutModalOpen={setIsCheckoutModalOpen}
        total={total}
        discount={discount}
        setCart={setCart}
      />

      <DeleteCartModal
        isDeleteModalOpen={isDeleteModalOpen}
        setIsDeleteModalOpen={setIsDeleteModalOpen}
        removeFromCart={removeFromCart}
        selectedItemId={selectedItemId}
      />
    </div>
  );
}
