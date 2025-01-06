import { NotificationState } from "@/types/state/notification";
import { create } from "zustand";
import { persist, PersistOptions } from "zustand/middleware";

const useNotificationStore = create(
  persist<NotificationState>(
    (set) => ({
      notifications: [],
      addNotification: (message) => {
        set((state) => ({
          notifications: [...state.notifications, { id: Date.now(), message }],
        }));
      },
      removeNotification: (id) => {
        set((state) => ({
          notifications: state.notifications.filter(
            (notification) => notification.id !== id,
          ),
        }));
      },
    }),
    {
      name: "notification-storage",
      getStorage: () => localStorage,
    } as PersistOptions<NotificationState>,
  ),
);

export default useNotificationStore;
