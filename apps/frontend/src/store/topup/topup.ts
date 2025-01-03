import { TopupStore } from "@/types/state/topup/topup";
import { create } from "zustand";
import { getAccessToken } from "../auth";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";

const useTopupStore = create<TopupStore>((set, get) => ({
    topups: null,
    topup: null,

    pagination: {
        currentPage: 1,
        pageSize: 10,
        totalItems: 0,
        totalPages: 0,
    },

   
    loadingGetTopups: false,
    loadingGetTopup: false,
    loadingGetActiveTopup: false,
    loadingGetTrashedTopup: false,
    loadingGetCardNumberTopup: false,

    loadingCreateTopup: false,
    loadingUpdateTopup: false,
    loadingTrashedTopup: false,
    loadingRestoreTopup: false,
    loadingPermanentTopup: false,

    
    errorGetTopups: null,
    errorGetTopup: null,
    errorGetActiveTopup: null,
    errorGetTrashedTopup: null,
    errorGetCardNumberTopup: null,

    errorCreateTopup: null,
    errorUpdateTopup: null,
    errorTrashedTopup: null,
    errorRestoreTopup: null,
    errorPermanentTopup: null,

    
    setLoadingGetTopups: (value) => set({ loadingGetTopups: value }),
    setLoadingGetTopup: (value) => set({ loadingGetTopup: value }),
    setLoadingGetActiveTopup: (value) => set({ loadingGetActiveTopup: value }),
    setLoadingGetTrashedTopup: (value) => set({ loadingGetTrashedTopup: value }),
    setLoadingGetCardNumberTopup: (value) => set({ loadingGetCardNumberTopup: value }),

    setLoadingCreateTopup: (value) => set({ loadingCreateTopup: value }),
    setLoadingUpdateTopup: (value) => set({ loadingUpdateTopup: value }),
    setLoadingTrashedTopup: (value) => set({ loadingTrashedTopup: value }),
    setLoadingRestoreTopup: (value) => set({ loadingRestoreTopup: value }),
    setLoadingPermanentTopup: (value) => set({ loadingPermanentTopup: value }),

    
    setErrorGetTopups: (value) => set({ errorGetTopups: value }),
    setErrorGetTopup: (value) => set({ errorGetTopup: value }),
    setErrorGetActiveTopup: (value) => set({ errorGetActiveTopup: value }),
    setErrorGetTrashedTopup: (value) => set({ errorGetTrashedTopup: value }),
    setErrorGetCardNumberTopup: (value) => set({ errorGetCardNumberTopup: value }),

    setErrorCreateTopup: (value) => set({ errorCreateTopup: value }),
    setErrorUpdateTopup: (value) => set({ errorUpdateTopup: value }),
    setErrorTrashedTopup: (value) => set({ errorTrashedTopup: value }),
    setErrorRestoreTopup: (value) => set({ errorRestoreTopup: value }),
    setErrorPermanentTopup: (value) => set({ errorPermanentTopup: value }),

    
    findAllTopups: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTopups: true, errorGetTopups: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/topups", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                topups: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTopups: false,
                errorGetTopups: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTopups: false }),
                (message: any) => set({ errorGetTopups: message })
            );
        }
    },

    findByIdTopup: async (id: number) => {
        set({ loadingGetTopup: true, errorGetTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/topups/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ topup: response.data, loadingGetTopup: false, errorGetTopup: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTopup: false }),
                (message: any) => set({ errorGetTopup: message })
            );
        }
    },

    findByActiveTopup: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetActiveTopup: true, errorGetActiveTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/topups/active", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                topups: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetActiveTopup: false,
                errorGetActiveTopup: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetActiveTopup: false }),
                (message: any) => set({ errorGetActiveTopup: message })
            );
        }
    },

    findByTrashedTopup: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTrashedTopup: true, errorGetTrashedTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/topups/trashed", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                topups: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTrashedTopup: false,
                errorGetTrashedTopup: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTrashedTopup: false }),
                (message: any) => set({ errorGetTrashedTopup: message })
            );
        }
    },

    findByCardNumberTopup: async (cardNumber: string) => {
        set({ loadingGetCardNumberTopup: true, errorGetCardNumberTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/topups/card-number/${cardNumber}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ topup: response.data, loadingGetCardNumberTopup: false, errorGetCardNumberTopup: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetCardNumberTopup: false }),
                (message: any) => set({ errorGetCardNumberTopup: message })
            );
        }
    },

    createTopup: async (req: any) => {
        set({ loadingCreateTopup: true, errorCreateTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.post("/topups", req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingCreateTopup: false, errorCreateTopup: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingCreateTopup: false }),
                (message: any) => set({ errorCreateTopup: message })
            );
        }
    },

    updateTopup: async (id: number, req: any) => {
        set({ loadingUpdateTopup: true, errorUpdateTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.put(`/topups/${id}`, req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingUpdateTopup: false, errorUpdateTopup: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingUpdateTopup: false }),
                (message: any) => set({ errorUpdateTopup: message })
            );
        }
    },

    restoreTopup: async (id: number) => {
        set({ loadingRestoreTopup: true, errorRestoreTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/topups/restore/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingRestoreTopup: false, errorRestoreTopup: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingRestoreTopup: false }),
                (message: any) => set({ errorRestoreTopup: message })
            );
        }
    },

    trashedTopup: async (id: number) => {
        set({ loadingTrashedTopup: true, errorTrashedTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/topups/trashed/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingTrashedTopup: false, errorTrashedTopup: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingTrashedTopup: false }),
                (message: any) => set({ errorTrashedTopup: message })
            );
        }
    },

    deleteTopupPermanent: async (id: number) => {
        set({ loadingPermanentTopup: true, errorPermanentTopup: null });
        try {
            const token = getAccessToken();
            const response = await myApi.delete(`/topups/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingPermanentTopup: false, errorPermanentTopup: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingPermanentTopup: false }),
                (message: any) => set({ errorPermanentTopup: message })
            );
        }
    },
}));

export default useTopupStore;