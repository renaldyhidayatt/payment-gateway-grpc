import { TransferStore } from "@/types/state/transfer/transfer";
import { create } from "zustand";
import { getAccessToken } from "../auth";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";

const useTransferStore = create<TransferStore>((set, get) => ({
    transfers: null,
    transfer: null,

    pagination: {
        currentPage: 1,
        pageSize: 10,
        totalItems: 0,
        totalPages: 0,
    },

    
    loadingGetTransfers: false,
    loadingGetTransfer: false,
    loadingGetTransferFrom: false,
    loadingGetTransferTo: false,
    loadingGetActiveTransfer: false,
    loadingGetTrashedTransfer: false,

    loadingCreateTransfer: false,
    loadingUpdateTransfer: false,
    loadingTrashedTransfer: false,
    loadingRestoreTransfer: false,
    loadingPermanentTransfer: false,

    
    errorGetTransfers: null,
    errorGetTransfer: null,
    errorGetTransferFrom: null,
    errorGetTransferTo: null,
    errorGetActiveTransfer: null,
    errorGetTrashedTransfer: null,

    errorCreateTransfer: null,
    errorUpdateTransfer: null,
    errorTrashedTransfer: null,
    errorRestoreTransfer: null,
    errorPermanentTransfer: null,

    
    setLoadingGetTransfers: (value) => set({ loadingGetTransfers: value }),
    setLoadingGetTransfer: (value) => set({ loadingGetTransfer: value }),
    setLoadingGetTransferFrom: (value) => set({ loadingGetTransferFrom: value }),
    setLoadingGetTransferTo: (value) => set({ loadingGetTransferTo: value }),
    setLoadingGetActiveTransfer: (value) => set({ loadingGetActiveTransfer: value }),
    setLoadingGetTrashedTransfer: (value) => set({ loadingGetTrashedTransfer: value }),

    setLoadingCreateTransfer: (value) => set({ loadingCreateTransfer: value }),
    setLoadingUpdateTransfer: (value) => set({ loadingUpdateTransfer: value }),
    setLoadingTrashedTransfer: (value) => set({ loadingTrashedTransfer: value }),
    setLoadingRestoreTransfer: (value) => set({ loadingRestoreTransfer: value }),
    setLoadingPermanentTransfer: (value) => set({ loadingPermanentTransfer: value }),

    
    setErrorGetTransfers: (value) => set({ errorGetTransfers: value }),
    setErrorGetTransfer: (value) => set({ errorGetTransfer: value }),
    setErrorGetTransferFrom: (value) => set({ errorGetTransferFrom: value }),
    setErrorGetTransferTo: (value) => set({ errorGetTransferTo: value }),
    setErrorGetActiveTransfer: (value) => set({ errorGetActiveTransfer: value }),
    setErrorGetTrashedTransfer: (value) => set({ errorGetTrashedTransfer: value }),

    setErrorCreateTransfer: (value) => set({ errorCreateTransfer: value }),
    setErrorUpdateTransfer: (value) => set({ errorUpdateTransfer: value }),
    setErrorTrashedTransfer: (value) => set({ errorTrashedTransfer: value }),
    setErrorRestoreTransfer: (value) => set({ errorRestoreTransfer: value }),
    setErrorPermanentTransfer: (value) => set({ errorPermanentTransfer: value }),

   
    findAllTransfers: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTransfers: true, errorGetTransfers: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/transfers", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                transfers: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTransfers: false,
                errorGetTransfers: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTransfers: false }),
                (message: any) => set({ errorGetTransfers: message })
            );
        }
    },

    findByIdTransfer: async (id: number) => {
        set({ loadingGetTransfer: true, errorGetTransfer: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/transfers/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ transfer: response.data, loadingGetTransfer: false, errorGetTransfer: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTransfer: false }),
                (message: any) => set({ errorGetTransfer: message })
            );
        }
    },

    findByTransferFrom: async (fromAccountId: number) => {
        set({ loadingGetTransferFrom: true, errorGetTransferFrom: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/transfers/from/${fromAccountId}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ transfers: response.data, loadingGetTransferFrom: false, errorGetTransferFrom: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTransferFrom: false }),
                (message: any) => set({ errorGetTransferFrom: message })
            );
        }
    },

    findByTransferTo: async (toAccountId: number) => {
        set({ loadingGetTransferTo: true, errorGetTransferTo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/transfers/to/${toAccountId}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ transfers: response.data, loadingGetTransferTo: false, errorGetTransferTo: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTransferTo: false }),
                (message: any) => set({ errorGetTransferTo: message })
            );
        }
    },

    findByActiveTransfer: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetActiveTransfer: true, errorGetActiveTransfer: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/transfers/active", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                transfers: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetActiveTransfer: false,
                errorGetActiveTransfer: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetActiveTransfer: false }),
                (message: any) => set({ errorGetActiveTransfer: message })
            );
        }
    },

    findByTrashedTransfer: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTrashedTransfer: true, errorGetTrashedTransfer: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/transfers/trashed", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                transfers: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTrashedTransfer: false,
                errorGetTrashedTransfer: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTrashedTransfer: false }),
                (message: any) => set({ errorGetTrashedTransfer: message })
            );
        }
    },

    createTransfer: async (req: any) => {
        set({ loadingCreateTransfer: true, errorCreateTransfer: null });
        try {
            const token = getAccessToken();
            const response = await myApi.post("/transfers", req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingCreateTransfer: false, errorCreateTransfer: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingCreateTransfer: false }),
                (message: any) => set({ errorCreateTransfer: message })
            );
        }
    },

    updateTransfer: async (id: number, req: any) => {
        set({ loadingUpdateTransfer: true, errorUpdateTransfer: null });
        try {
            const token = getAccessToken();
            const response = await myApi.put(`/transfers/${id}`, req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingUpdateTransfer: false, errorUpdateTransfer: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingUpdateTransfer: false }),
                (message: any) => set({ errorUpdateTransfer: message })
            );
        }
    },

    restoreTransfer: async (id: number) => {
        set({ loadingRestoreTransfer: true, errorRestoreTransfer: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/transfers/restore/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingRestoreTransfer: false, errorRestoreTransfer: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingRestoreTransfer: false }),
                (message: any) => set({ errorRestoreTransfer: message })
            );
        }
    },

    trashedTransfer: async (id: number) => {
        set({ loadingTrashedTransfer: true, errorTrashedTransfer: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/transfers/trashed/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingTrashedTransfer: false, errorTrashedTransfer: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingTrashedTransfer: false }),
                (message: any) => set({ errorTrashedTransfer: message })
            );
        }
    },

    deleteTransferPermanent: async (id: number) => {
        set({ loadingPermanentTransfer: true, errorPermanentTransfer: null });
        try {
            const token = getAccessToken();
            const response = await myApi.delete(`/transfers/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingPermanentTransfer: false, errorPermanentTransfer: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingPermanentTransfer: false }),
                (message: any) => set({ errorPermanentTransfer: message })
            );
        }
    },
}));

export default useTransferStore;