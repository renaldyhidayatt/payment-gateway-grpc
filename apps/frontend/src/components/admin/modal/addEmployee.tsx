import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogFooter, DialogTitle, DialogClose } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Plus } from "lucide-react";
import Select from "react-select";
import { ScrollArea } from "@/components/ui/scroll-area"; // Import your ScrollArea component

// Example options for the select input
const roleOptions = [
  { value: 'admin', label: 'Admin' },
  { value: 'manager', label: 'Manager' },
  { value: 'employee', label: 'Employee' },
];

export function AddEmployee({ onSubmit }: any) {
  const [isOpen, setIsOpen] = useState(false);
  const [selectedRole, setSelectedRole] = useState(null);
  const [employeeName, setEmployeeName] = useState('');
  const [employeeEmail, setEmployeeEmail] = useState('');

  const handleSubmit = () => {
    const employeeData = {
      name: employeeName,
      email: employeeEmail,
      role: selectedRole ? selectedRole.value : null,
    };
    onSubmit(employeeData);
    setIsOpen(false); // Close dialog after submission
    // Reset input fields
    setEmployeeName('');
    setEmployeeEmail('');
    setSelectedRole(null);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="default" size="sm">
          <Plus className="mr-2 h-4 w-4" />
          Add Employee
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>Add New Employee</DialogTitle>
        </DialogHeader>
        <ScrollArea className="max-h-[400px]">
          <div className="space-y-4 p-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Employee Name</label>
              <Input 
                type="text" 
                placeholder="Enter name" 
                className="mt-1" 
                value={employeeName}
                onChange={(e) => setEmployeeName(e.target.value)}
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Employee Email</label>
              <Input 
                type="email" 
                placeholder="Enter email" 
                className="mt-1" 
                value={employeeEmail}
                onChange={(e) => setEmployeeEmail(e.target.value)}
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Role</label>
              <Select 
                options={roleOptions} 
                value={selectedRole} 
                onChange={setSelectedRole} 
                placeholder="Select a role" 
                isSearchable // Enable searching in the dropdown
              />
            </div>
          </div>
        </ScrollArea>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline" onClick={() => setIsOpen(false)}>
              Cancel
            </Button>
          </DialogClose>
          <Button variant="default" onClick={handleSubmit}>
            Submit
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
