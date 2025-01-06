import React, { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
  DropdownMenuItem,
  DropdownMenuSeparator,
} from "@/components/ui/dropdown-menu";
import { Bell, MessageSquare } from "lucide-react";
import { Button } from "@/components/ui/button";
import { NotificationCommand } from "./command-notification";
import { Notification } from "@/types/model";
import useNotificationStore from "@/store/notication";

export function NotificationMenu() {
  const [openCommand, setOpenCommand] = useState(false);
  const { notifications, removeNotification } = useNotificationStore();

  const displayedNotifications = notifications.slice(0, 3);

  React.useEffect(() => {
    const handleShortcut = (e: KeyboardEvent) => {
      if (e.key === "d" && (e.metaKey || e.ctrlKey)) {
        e.preventDefault();
        setOpenCommand(true);
      }
    };

    document.addEventListener("keydown", handleShortcut);
    return () => document.removeEventListener("keydown", handleShortcut);
  }, []);

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" size="icon" className="relative">
            <Bell className="h-5 w-5" />
            {notifications.length > 0 && (
              <span className="absolute right-1 top-1 inline-flex h-2 w-2 rounded-full bg-red-500"></span>
            )}
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end" className="w-64">
          {notifications.length === 0 ? (
            <DropdownMenuItem disabled className="text-gray-500">
              No notifications
            </DropdownMenuItem>
          ) : (
            <>
              {displayedNotifications.map((notification: Notification) => (
                <DropdownMenuItem
                  key={notification.id}
                  onSelect={() => removeNotification(notification.id)}
                >
                  <MessageSquare className="mr-2 h-4 w-4 text-blue-500" />
                  <span>{notification.message}</span>
                </DropdownMenuItem>
              ))}
              {notifications.length > 3 && (
                <>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem
                    className="text-blue-600"
                    onClick={() => setOpenCommand(true)}
                  >
                    <span>See all notifications</span>
                  </DropdownMenuItem>
                </>
              )}
            </>
          )}
        </DropdownMenuContent>
      </DropdownMenu>

      <NotificationCommand open={openCommand} setOpen={setOpenCommand} />
    </>
  );
}
