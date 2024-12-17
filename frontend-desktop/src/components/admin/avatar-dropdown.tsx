import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import Luffy from '@/assets/Luffy.jpg';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { ChevronDown } from 'lucide-react';

export function AvatarDropdown() {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="ghost"
          className="flex items-center space-x-2 text-gray-900 dark:text-white"
        >
          <Avatar className="h-8 w-8">
            <AvatarImage src={Luffy} alt="User" />
            <AvatarFallback>R</AvatarFallback>
          </Avatar>
          <ChevronDown className="h-4 w-4" />
        </Button>
      </DropdownMenuTrigger>

      <DropdownMenuContent align="end" className="bg-white dark:bg-gray-700">
        <DropdownMenuLabel className="text-gray-900 dark:text-gray-200">
          My Account
        </DropdownMenuLabel>
        <DropdownMenuSeparator className="border-gray-200 dark:border-gray-600" />
        <DropdownMenuItem className="text-gray-900 dark:text-gray-200">
          Profile
        </DropdownMenuItem>
        <DropdownMenuItem className="text-gray-900 dark:text-gray-200">
          Settings
        </DropdownMenuItem>
        <DropdownMenuSeparator className="border-gray-200 dark:border-gray-600" />
        <DropdownMenuItem className="text-gray-900 dark:text-gray-200">
          Log out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
