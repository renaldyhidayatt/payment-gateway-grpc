import { create } from "zustand";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";
import { UserStore } from "@/types/state/user";
import { CreateUser, UpdateUser } from "@/types/domain/request";
import { getAccessToken } from "../auth";

const useUserStore = create<UserStore>((set, get) => ({
  users: null,
  user: null,

  pagination: {
    currentPage: 1,
    pageSize: 20,
    totalItems: 0,
    totalPages: 0,
  },

  loadingGetUsers: false,
  loadingGetUser: false,
  loadingGetActiveUsers: false,
  loadingGetTrashedUsers: false,
  loadingCreateUser: false,
  loadingUpdateUser: false,
  loadingTrashedUser: false,
  loadingRestoreUser: false,
  loadingDeleteUser: false,

  errorGetUsers: null,
  errorGetUser: null,
  errorGetActiveUsers: null,
  errorGetTrashedUsers: null,
  errorCreateUser: null,
  errorUpdateUser: null,
  errorTrashedUser: null,
  errorRestoreUser: null,
  errorDeleteUser: null,

  setErrorGetUsers: (value: string | null) => set({ errorGetUsers: value }),
  setErrorGetUser: (value: string | null) => set({ errorGetUser: value }),
  setErrorGetActiveUsers: (value: string | null) =>
    set({ errorGetActiveUsers: value }),
  setErrorGetTrashedUsers: (value: string | null) =>
    set({ errorGetTrashedUsers: value }),
  setErrorCreateUser: (value: string | null) => set({ errorCreateUser: value }),
  setErrorUpdateUser: (value: string | null) => set({ errorUpdateUser: value }),
  setErrorTrashedUser: (value: string | null) =>
    set({ errorTrashedUser: value }),
  setErrorRestoreUser: (value: string | null) =>
    set({ errorRestoreUser: value }),
  setErrorDeleteUser: (value: string | null) => set({ errorDeleteUser: value }),

  setLoadingGetUsers: (value: boolean) => set({ loadingGetUsers: value }),
  setLoadingGetUser: (value: boolean) => set({ loadingGetUser: value }),
  setLoadingGetActiveUsers: (value: boolean) =>
    set({ loadingGetActiveUsers: value }),
  setLoadingGetTrashedUsers: (value: boolean) =>
    set({ loadingGetTrashedUsers: value }),
  setLoadingCreateUser: (value: boolean) => set({ loadingCreateUser: value }),
  setLoadingUpdateUser: (value: boolean) => set({ loadingUpdateUser: value }),
  setLoadingTrashedUser: (value: boolean) => set({ loadingTrashedUser: value }),
  setLoadingRestoreUser: (value: boolean) => set({ loadingRestoreUser: value }),
  setLoadingDeleteUser: (value: boolean) => set({ loadingDeleteUser: value }),

  findAllUsers: async (search: string, page: number, pageSize: number) => {
    set({ loadingGetUsers: true, errorGetUsers: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get(`/user`, {
        params: { page, page_size: pageSize, search },
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      console.log("response", response);

      set({
        users: response.data.data,
        pagination: {
          currentPage: response.data.pagination.current_page,
          pageSize: response.data.pagination.page_size,
          totalItems: response.data.pagination.total_records,
          totalPages: response.data.pagination.total_pages,
        },
        loadingGetUsers: false,
        errorGetUsers: null,
      });
    } catch (err) {
      console.log("error: ", err);

      handleApiError(
        err,
        () => set({ loadingGetUsers: false }),
        (message: any) => set({ errorGetUsers: message }),
      );
    }
  },

  findById: async (id: number) => {
    set({ loadingGetUser: true, errorGetUser: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get(`/users/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      set({ user: response.data, loadingGetUser: false, errorGetUser: null });
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingGetUser: false }),
        (message: any) => set({ errorGetUser: message }),
      );
    }
  },

  findByActive: async (search: string, page: number, pageSize: number) => {
    set({ loadingGetActiveUsers: true, errorGetActiveUsers: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get(`/users/active`, {
        params: { page, pageSize, search },
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      set({
        users: response.data.items,
        pagination: {
          currentPage: response.data.pagination.current_page,
          pageSize: response.data.pagination.page_size,
          totalItems: response.data.pagination.total_records,
          totalPages: response.data.pagination.total_pages,
        },
        loadingGetActiveUsers: false,
        errorGetActiveUsers: null,
      });
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingGetActiveUsers: false }),
        (message: any) => set({ errorGetActiveUsers: message }),
      );
    }
  },

  findByTrashed: async (search: string, page: number, pageSize: number) => {
    set({ loadingGetTrashedUsers: true, errorGetTrashedUsers: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get(`/users/trashed`, {
        params: { page, pageSize, search },
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      set({
        users: response.data.items,
        pagination: {
          currentPage: response.data.pagination.current_page,
          pageSize: response.data.pagination.page_size,
          totalItems: response.data.pagination.total_records,
          totalPages: response.data.pagination.total_pages,
        },
        loadingGetTrashedUsers: false,
        errorGetTrashedUsers: null,
      });
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingGetTrashedUsers: false }),
        (message: any) => set({ errorGetTrashedUsers: message }),
      );
    }
  },

  createUser: async (req: CreateUser) => {
    set({ loadingCreateUser: true, errorCreateUser: null });
    try {
      const token = getAccessToken();
      const response = await myApi.post(`/users`, req, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.status == 201) {
        set({
          loadingCreateUser: false,
          errorCreateUser: null,
          user: response.data.data,
        });
        return true;
      } else {
        throw new Error("Create User gagal. Silahkan coba create lagi");
      }
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingCreateUser: false }),
        (message: any) => set({ errorCreateUser: message }),
      );

      return false;
    }
  },

  updateUser: async (id: number, req: UpdateUser) => {
    set({ loadingUpdateUser: true, errorUpdateUser: null });
    try {
      const token = getAccessToken();
      const response = await myApi.put(`/users/${id}`, req, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.status == 200) {
        set({ loadingUpdateUser: false, errorUpdateUser: null });

        return true;
      } else {
        throw new Error("Update User gagal. Silahkan update lagi");
      }
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingUpdateUser: false }),
        (message: any) => set({ errorUpdateUser: message }),
      );

      return false;
    }
  },

  trashedUser: async (id: number) => {
    set({ loadingTrashedUser: true, errorTrashedUser: null });

    try {
      const token = getAccessToken();
      const response = await myApi.delete(`/users/trash/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      if (response.status === 200) {
        set({ loadingTrashedUser: false, errorTrashedUser: null });
        return true;
      } else {
        throw new Error("Trashed User Gagal. Silahkan coba lagi");
      }
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingTrashedUser: false }),
        (message: any) => set({ errorTrashedUser: message }),
      );
      return false;
    }
  },

  restoreUser: async (id: number) => {
    set({ loadingRestoreUser: true, errorRestoreUser: null });

    try {
      const token = getAccessToken();
      const response = await myApi.patch(`/users/restore/${id}`, null, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.status === 200) {
        set({ loadingRestoreUser: false, errorRestoreUser: null });
        return true;
      } else {
        throw new Error("Restore User Gagal. Silahkan coba lagi");
      }
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingRestoreUser: false }),
        (message: any) => set({ errorRestoreUser: message }),
      );
      return false;
    }
  },

  deleteUserPermanent: async (id: number) => {
    set({ loadingDeleteUser: true, errorDeleteUser: null });

    try {
      const token = getAccessToken();
      const response = await myApi.delete(`/users/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      if (response.status === -200) {
        set({ loadingDeleteUser: false, errorDeleteUser: null });
        return true;
      } else {
        throw new Error("Delete Permanent User Gagal. Silahkan coba lagi");
      }
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingDeleteUser: false }),
        (message: any) => set({ errorDeleteUser: message }),
      );
      return false;
    }
  },
}));

export default useUserStore;
