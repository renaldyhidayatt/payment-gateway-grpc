import { WithdrawStore } from "@/types/state/withdraw/withdraw";
import { create } from "zustand";
import { getAccessToken } from "../auth";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";

const useWithdrawStore = create<WithdrawStore>((set, get) => ({
    withdraws: null,
    withdraw: null,

    pagination: {
        currentPage: 1,
        pageSize: 10,
        totalItems: 0,
        totalPages: 0,
    },

    
    loadingGetWithdraws: false,
    loadingGetWithdraw: false,
    loadingGetCardNumberWithdraw: false,
    loadingGetActiveWithdraw: false,
    loadingGetTrashedWithdraw: false,

    loadingCreateWithdraw: false,
    loadingUpdateWithdraw: false,
    loadingTrashedWithdraw: false,
    loadingRestoreWithdraw: false,
    loadingPermanentWithdraw: false,

    
    errorGetWithdraws: null,
    errorGetWithdraw: null,
    errorGetCardNumberWithdraw: null,
    errorGetActiveWithdraw: null,
    errorGetTrashedWithdraw: null,

    errorCreateWithdraw: null,
    errorUpdateWithdraw: null,
    errorTrashedWithdraw: null,
    errorRestoreWithdraw: null,
    errorPermanentWithdraw: null,

   
    setLoadingGetWithdraws: (value) => set({ loadingGetWithdraws: value }),
    setLoadingGetWithdraw: (value) => set({ loadingGetWithdraw: value }),
    setLoadingGetCardNumberWithdraw: (value) => set({ loadingGetCardNumberWithdraw: value }),
    setLoadingGetActiveWithdraw: (value) => set({ loadingGetActiveWithdraw: value }),
    setLoadingGetTrashedWithdraw: (value) => set({ loadingGetTrashedWithdraw: value }),

    setLoadingCreateWithdraw: (value) => set({ loadingCreateWithdraw: value }),
    setLoadingUpdateWithdraw: (value) => set({ loadingUpdateWithdraw: value }),
    setLoadingTrashedWithdraw: (value) => set({ loadingTrashedWithdraw: value }),
    setLoadingRestoreWithdraw: (value) => set({ loadingRestoreWithdraw: value }),
    setLoadingPermanentWithdraw: (value) => set({ loadingPermanentWithdraw: value }),

    setErrorGetWithdraws: (value) => set({ errorGetWithdraws: value }),
    setErrorGetWithdraw: (value) => set({ errorGetWithdraw: value }),
    setErrorGetCardNumberWithdraw: (value) => set({ errorGetCardNumberWithdraw: value }),
    setErrorGetActiveWithdraw: (value) => set({ errorGetActiveWithdraw: value }),
    setErrorGetTrashedWithdraw: (value) => set({ errorGetTrashedWithdraw: value }),

    setErrorCreateWithdraw: (value) => set({ errorCreateWithdraw: value }),
    setErrorUpdateWithdraw: (value) => set({ errorUpdateWithdraw: value }),
    setErrorTrashedWithdraw: (value) => set({ errorTrashedWithdraw: value }),
    setErrorRestoreWithdraw: (value) => set({ errorRestoreWithdraw: value }),
    setErrorPermanentWithdraw: (value) => set({ errorPermanentWithdraw: value }),

    
    findAllWithdraws: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetWithdraws: true, errorGetWithdraws: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/withdraws", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                withdraws: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetWithdraws: false,
                errorGetWithdraws: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetWithdraws: false }),
                (message: any) => set({ errorGetWithdraws: message })
            );
        }
    },

    findByIdWithdraw: async (id: number) => {
        set({ loadingGetWithdraw: true, errorGetWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/withdraws/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ withdraw: response.data, loadingGetWithdraw: false, errorGetWithdraw: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetWithdraw: false }),
                (message: any) => set({ errorGetWithdraw: message })
            );
        }
    },

    findByCardNumberWithdraw: async (cardNumber: string) => {
        set({ loadingGetCardNumberWithdraw: true, errorGetCardNumberWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/withdraws/card-number/${cardNumber}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ withdraw: response.data, loadingGetCardNumberWithdraw: false, errorGetCardNumberWithdraw: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetCardNumberWithdraw: false }),
                (message: any) => set({ errorGetCardNumberWithdraw: message })
            );
        }
    },

    findByActiveWithdraw: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetActiveWithdraw: true, errorGetActiveWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/withdraws/active", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                withdraws: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetActiveWithdraw: false,
                errorGetActiveWithdraw: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetActiveWithdraw: false }),
                (message: any) => set({ errorGetActiveWithdraw: message })
            );
        }
    },

    findByTrashedWithdraw: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTrashedWithdraw: true, errorGetTrashedWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/withdraws/trashed", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                withdraws: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTrashedWithdraw: false,
                errorGetTrashedWithdraw: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTrashedWithdraw: false }),
                (message: any) => set({ errorGetTrashedWithdraw: message })
            );
        }
    },

    createWithdraw: async (req: any) => {
        set({ loadingCreateWithdraw: true, errorCreateWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.post("/withdraws", req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingCreateWithdraw: false, errorCreateWithdraw: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingCreateWithdraw: false }),
                (message: any) => set({ errorCreateWithdraw: message })
            );
        }
    },

    updateWithdraw: async (id: number, req: any) => {
        set({ loadingUpdateWithdraw: true, errorUpdateWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.put(`/withdraws/${id}`, req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingUpdateWithdraw: false, errorUpdateWithdraw: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingUpdateWithdraw: false }),
                (message: any) => set({ errorUpdateWithdraw: message })
            );
        }
    },

    restoreWithdraw: async (id: number) => {
        set({ loadingRestoreWithdraw: true, errorRestoreWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/withdraws/restore/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingRestoreWithdraw: false, errorRestoreWithdraw: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingRestoreWithdraw: false }),
                (message: any) => set({ errorRestoreWithdraw: message })
            );
        }
    },

    trashedWithdraw: async (id: number) => {
        set({ loadingTrashedWithdraw: true, errorTrashedWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/withdraws/trashed/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingTrashedWithdraw: false, errorTrashedWithdraw: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingTrashedWithdraw: false }),
                (message: any) => set({ errorTrashedWithdraw: message })
            );
        }
    },

    deleteWithdrawPermanent: async (id: number) => {
        set({ loadingPermanentWithdraw: true, errorPermanentWithdraw: null });
        try {
            const token = getAccessToken();
            const response = await myApi.delete(`/withdraws/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingPermanentWithdraw: false, errorPermanentWithdraw: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingPermanentWithdraw: false }),
                (message: any) => set({ errorPermanentWithdraw: message })
            );
        }
    },
}));

export default useWithdrawStore;