import React from 'react';
import { MessageSquare, CheckCircle, AlertCircle } from 'lucide-react';
import {
  CommandItem,
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandList,
} from '@/components/ui/command';

type NotificationCommandProps = {
  open: boolean;
  setOpen: (open: boolean) => void;
};

export function NotificationCommand({
  open,
  setOpen,
}: NotificationCommandProps) {
  React.useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === 'd' && (e.metaKey || e.ctrlKey)) {
        e.preventDefault();
        setOpen(!open);
      }
    };
    document.addEventListener('keydown', down);
    return () => document.removeEventListener('keydown', down);
  }, [setOpen]);

  return (
    <>
      <CommandDialog open={open} onOpenChange={setOpen}>
        <CommandInput placeholder="Type a command or search..." />
        <CommandList>
          <CommandEmpty>No results found.</CommandEmpty>
          <CommandGroup heading="Notifications">
            <CommandItem>
              <MessageSquare className="mr-2 h-4 w-4 text-blue-500" />
              <span>New comment on your post</span>
            </CommandItem>
            <CommandItem>
              <CheckCircle className="mr-2 h-4 w-4 text-green-500" />
              <span>Task completed successfully</span>
            </CommandItem>
            <CommandItem>
              <AlertCircle className="mr-2 h-4 w-4 text-yellow-500" />
              <span>Warning: Server usage is high</span>
            </CommandItem>
          </CommandGroup>
        </CommandList>
      </CommandDialog>
    </>
  );
}
