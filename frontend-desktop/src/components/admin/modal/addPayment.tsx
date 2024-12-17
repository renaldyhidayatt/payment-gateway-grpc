import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogFooter, DialogTitle, DialogClose } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Plus } from "lucide-react";

export function AddPayment({ onSubmit }: any) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          <Plus className="mr-2 h-4 w-4" />
          Add Payment
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>Add New Payment</DialogTitle>
        </DialogHeader>
        <div className="space-y-4 max-h-[400px] overflow-y-auto"> {/* Added max height and overflow */}
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Payment Amount</label>
            <Input 
              type="number" 
              placeholder="Enter amount" 
              className="mt-1" 
             
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Payment Date</label>
            <Input 
              type="date" 
              className="mt-1" 
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Description</label>
            <Input 
              type="text" 
              placeholder="Description" 
              className="mt-1" 
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Payment Amount</label>
            <Input 
              type="number" 
              placeholder="Enter amount" 
              className="mt-1" 
             
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Payment Date</label>
            <Input 
              type="date" 
              className="mt-1" 
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Description</label>
            <Input 
              type="text" 
              placeholder="Description" 
              className="mt-1" 
            />
          </div>
        </div>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline" onClick={() => setIsOpen(false)}>
              Cancel
            </Button>
          </DialogClose>
          <Button variant="default" onClick={onSubmit}>
            Submit
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
