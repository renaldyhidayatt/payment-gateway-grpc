import React, { useState } from 'react';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
  DropdownMenuItem,
  DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu';
import { Bell, MessageSquare, CheckCircle, AlertCircle } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { NotificationCommand } from './command-notification';

export function NotificationMenu() {
  const [openCommand, setOpenCommand] = useState(false);

  React.useEffect(() => {
    const handleShortcut = (e: KeyboardEvent) => {
      if (e.key === 'd' && (e.metaKey || e.ctrlKey)) {
        e.preventDefault();
        setOpenCommand(true);
      }
    };

    document.addEventListener('keydown', handleShortcut);
    return () => document.removeEventListener('keydown', handleShortcut);
  }, []);

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" size="icon" className="relative">
            <Bell className="h-5 w-5" />
            <span className="absolute right-1 top-1 inline-flex h-2 w-2 rounded-full bg-red-500"></span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end" className="w-64">
          <DropdownMenuItem>
            <MessageSquare className="mr-2 h-4 w-4 text-blue-500" />
            <span>New comment on your post</span>
          </DropdownMenuItem>
          <DropdownMenuItem>
            <CheckCircle className="mr-2 h-4 w-4 text-green-500" />
            <span>Task completed successfully</span>
          </DropdownMenuItem>
          <DropdownMenuItem>
            <AlertCircle className="mr-2 h-4 w-4 text-yellow-500" />
            <span>Warning: Server usage is high</span>
          </DropdownMenuItem>
          <DropdownMenuSeparator />
          <DropdownMenuItem
            className="text-blue-600"
            onClick={() => setOpenCommand(true)}
          >
            <span>See all notifications</span>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>

      <NotificationCommand open={openCommand} setOpen={setOpenCommand} />
    </>
  );
}
