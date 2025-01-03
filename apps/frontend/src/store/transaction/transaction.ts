import { TransactionStore } from "@/types/state/transaction/transaction";
import { create } from "zustand";
import { getAccessToken } from "../auth";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";

const useTransactionStore = create<TransactionStore>((set, get) => ({
    transactions: null,
    transaction: null,

    pagination: {
        currentPage: 1,
        pageSize: 10,
        totalItems: 0,
        totalPages: 0,
    },

   
    loadingGetTransactions: false,
    loadingGetTransaction: false,
    loadingGetCardNumberTransaction: false,
    loadingGetMerchantTransaction: false,
    loadingGetActiveTransaction: false,
    loadingGetTrashedTransaction: false,

    loadingCreateTransaction: false,
    loadingUpdateTransaction: false,
    loadingRestoreTransaction: false,
    loadingTrashedTransaction: false,
    loadingDeletePermanentTransaction: false,

    errorGetTransactions: null,
    errorGetTransaction: null,
    errorGetCardNumberTransaction: null,
    errorGetMerchantTransaction: null,
    errorGetActiveTransaction: null,
    errorGetTrashedTransaction: null,

    errorCreateTransaction: null,
    errorUpdateTransaction: null,
    errorRestoreTransaction: null,
    errorTrashedTransaction: null,
    errorDeletePermanentTransaction: null,

    
    setLoadingGetTransactions: (value) => set({ loadingGetTransactions: value }),
    setLoadingGetTransaction: (value) => set({ loadingGetTransaction: value }),
    setLoadingGetCardNumberTransaction: (value) => set({ loadingGetCardNumberTransaction: value }),
    setLoadingGetMerchantTransaction: (value) => set({ loadingGetMerchantTransaction: value }),
    setLoadingGetActiveTransaction: (value) => set({ loadingGetActiveTransaction: value }),
    setLoadingGetTrashedTransaction: (value) => set({ loadingGetTrashedTransaction: value }),

    setLoadingCreateTransaction: (value) => set({ loadingCreateTransaction: value }),
    setLoadingUpdateTransaction: (value) => set({ loadingUpdateTransaction: value }),
    setLoadingRestoreTransaction: (value) => set({ loadingRestoreTransaction: value }),
    setLoadingTrashedTransaction: (value) => set({ loadingTrashedTransaction: value }),
    setLoadingDeletePermanentTransaction: (value) => set({ loadingDeletePermanentTransaction: value }),

    setErrorGetTransactions: (value) => set({ errorGetTransactions: value }),
    setErrorGetTransaction: (value) => set({ errorGetTransaction: value }),
    setErrorGetCardNumberTransaction: (value) => set({ errorGetCardNumberTransaction: value }),
    setErrorGetMerchantTransaction: (value) => set({ errorGetMerchantTransaction: value }),
    setErrorGetActiveTransaction: (value) => set({ errorGetActiveTransaction: value }),
    setErrorGetTrashedTransaction: (value) => set({ errorGetTrashedTransaction: value }),

    setErrorCreateTransaction: (value) => set({ errorCreateTransaction: value }),
    setErrorUpdateTransaction: (value) => set({ errorUpdateTransaction: value }),
    setErrorRestoreTransaction: (value) => set({ errorRestoreTransaction: value }),
    setErrorTrashedTransaction: (value) => set({ errorTrashedTransaction: value }),
    setErrorDeletePermanentTransaction: (value) => set({ errorDeletePermanentTransaction: value }),

    findAllTransactions: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTransactions: true, errorGetTransactions: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/transactions", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                transactions: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTransactions: false,
                errorGetTransactions: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTransactions: false }),
                (message: any) => set({ errorGetTransactions: message })
            );
        }
    },

    findByIdTransaction: async (id: number) => {
        set({ loadingGetTransaction: true, errorGetTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/transactions/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ transaction: response.data, loadingGetTransaction: false, errorGetTransaction: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTransaction: false }),
                (message: any) => set({ errorGetTransaction: message })
            );
        }
    },

    findByCardNumberTransaction: async (cardNumber: string) => {
        set({ loadingGetCardNumberTransaction: true, errorGetCardNumberTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/transactions/card-number/${cardNumber}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ transaction: response.data, loadingGetCardNumberTransaction: false, errorGetCardNumberTransaction: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetCardNumberTransaction: false }),
                (message: any) => set({ errorGetCardNumberTransaction: message })
            );
        }
    },

    findByMerchantTransaction: async (merchantId: number) => {
        set({ loadingGetMerchantTransaction: true, errorGetMerchantTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/transactions/merchant/${merchantId}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ transactions: response.data, loadingGetMerchantTransaction: false, errorGetMerchantTransaction: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetMerchantTransaction: false }),
                (message: any) => set({ errorGetMerchantTransaction: message })
            );
        }
    },

    findByActiveTransaction: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetActiveTransaction: true, errorGetActiveTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/transactions/active", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                transactions: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetActiveTransaction: false,
                errorGetActiveTransaction: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetActiveTransaction: false }),
                (message: any) => set({ errorGetActiveTransaction: message })
            );
        }
    },

    findByTrashedTransaction: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTrashedTransaction: true, errorGetTrashedTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/transactions/trashed", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                transactions: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTrashedTransaction: false,
                errorGetTrashedTransaction: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTrashedTransaction: false }),
                (message: any) => set({ errorGetTrashedTransaction: message })
            );
        }
    },

    createTransaction: async (req: any) => {
        set({ loadingCreateTransaction: true, errorCreateTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.post("/transactions", req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingCreateTransaction: false, errorCreateTransaction: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingCreateTransaction: false }),
                (message: any) => set({ errorCreateTransaction: message })
            );
        }
    },

    updateTransaction: async (id: number, req: any) => {
        set({ loadingUpdateTransaction: true, errorUpdateTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.put(`/transactions/${id}`, req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingUpdateTransaction: false, errorUpdateTransaction: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingUpdateTransaction: false }),
                (message: any) => set({ errorUpdateTransaction: message })
            );
        }
    },

    restoreTransaction: async (id: number) => {
        set({ loadingRestoreTransaction: true, errorRestoreTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/transactions/restore/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingRestoreTransaction: false, errorRestoreTransaction: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingRestoreTransaction: false }),
                (message: any) => set({ errorRestoreTransaction: message })
            );
        }
    },

    trashedTransaction: async (id: number) => {
        set({ loadingTrashedTransaction: true, errorTrashedTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/transactions/trashed/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingTrashedTransaction: false, errorTrashedTransaction: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingTrashedTransaction: false }),
                (message: any) => set({ errorTrashedTransaction: message })
            );
        }
    },

    deleteTransactionPermanent: async (id: number) => {
        set({ loadingDeletePermanentTransaction: true, errorDeletePermanentTransaction: null });
        try {
            const token = getAccessToken();
            const response = await myApi.delete(`/transactions/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingDeletePermanentTransaction: false, errorDeletePermanentTransaction: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingDeletePermanentTransaction: false }),
                (message: any) => set({ errorDeletePermanentTransaction: message })
            );
        }
    },
}));

export default useTransactionStore;