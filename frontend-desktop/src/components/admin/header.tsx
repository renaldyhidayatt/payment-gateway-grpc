import { useLocation } from 'react-router-dom'; // Import useLocation
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Search } from 'lucide-react';
import { ModeToggle } from './mode-toggle';
import { SidebarTrigger } from '../ui/sidebar';
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage,
} from '@/components/ui/breadcrumb';
import { Separator } from '../ui/separator';
import { CommandMenu } from './command';
import { NotificationDropdown } from './notification-dropdown';
import { AvatarDropdown } from './avatar-dropdown';
import { useState } from 'react';

export function Header() {
  const location = useLocation();
  const pathname = location.pathname;
  const [isSearchOpen, setIsSearchOpen] = useState(false);

  const pageTitles: Record<string, string> = {
    '/': 'Dashboard',
    '/point-of-sale': 'Point Of Sale',
    '/transactions': 'Transactions',
    '/settings': 'Settings',
  };

  let currentTitle = 'Page';
  if (pageTitles[pathname]) {
    currentTitle = pageTitles[pathname];
  } else if (pathname.startsWith('/user/detail')) {
    currentTitle = 'User Details';
  } else if (pathname.startsWith('/transactions/detail')) {
    currentTitle = 'Transaction Details';
  }

  return (
    <header className="flex h-16 shrink-0 items-center justify-between gap-2 border-b px-4">
      <div className="flex items-center gap-2">
        <SidebarTrigger />
        <Separator orientation="vertical" className="h-6" />
        <Breadcrumb>
          <BreadcrumbList>
            <BreadcrumbItem>
              <BreadcrumbPage>{currentTitle}</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
      </div>

      <div className="flex items-center gap-4">
        <div className="relative">
          <Input
            type="text"
            placeholder="Search..."
            className="w-48 cursor-pointer"
            readOnly
          />
          <Search className="absolute right-3 top-1/2 h-4 w-4 -translate-y-1/2 text-gray-400" />
        </div>

        <CommandMenu open={isSearchOpen} setOpen={setIsSearchOpen} />
        <ModeToggle />
        <NotificationDropdown />
        <AvatarDropdown />
      </div>
    </header>
  );
}
