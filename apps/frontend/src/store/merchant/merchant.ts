import { create } from "zustand";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";
import { MerchantStore } from "@/types/state/merchant/merchant";
import { CreateMerchant, UpdateMerchant } from "@/types/domain/request";
import { getAccessToken } from "../auth";

const useMerchantStore = create<MerchantStore>((set, get) => ({
    merchants: null,
    merchant: null,

    pagination: {
        currentPage: 1,
        pageSize: 10,
        totalItems: 0,
        totalPages: 0,
    },

    loadingGetMerchants: false,
    loadingGetMerchant: false,
    loadingGetApiKey: false,
    loadingGetMerchantUser: false,
    loadingGetActiveMerchant: false,
    loadingGetTrashedMerchant: false,
    loadingCreateMerchant: false,
    loadingUpdateMerchant: false,
    loadingTrashedMerchant: false,
    loadingRestoreMerchant: false,
    loadingDeletePermanentMerchant: false,

    errorGetMerchants: null,
    errorGetMerchant: null,
    errorGetApiKey: null,
    errorGetMerchantUser: null,
    errorGetActiveMerchant: null,
    errorGetTrashedMerchant: null,
    errorCreateMerchant: null,
    errorUpdateMerchant: null,
    errorTrashedMerchant: null,
    errorRestoreMerchant: null,
    errorDeletePermanentMerchant: null,

    
    setLoadingGetMerchants: (value) => set({ loadingGetMerchants: value }),
    setLoadingGetMerchant: (value) => set({ loadingGetMerchant: value }),
    setLoadingGetApiKey: (value) => set({ loadingGetApiKey: value }),
    setLoadingGetMerchantUser: (value) => set({ loadingGetMerchantUser: value }),
    setLoadingGetActiveMerchant: (value) => set({ loadingGetActiveMerchant: value }),
    setLoadingGetTrashedMerchant: (value) => set({ loadingGetTrashedMerchant: value }),
    setLoadingCreateMerchant: (value) => set({ loadingCreateMerchant: value }),
    setLoadingUpdateMerchant: (value) => set({ loadingUpdateMerchant: value }),
    setLoadingTrashedMerchant: (value) => set({ loadingTrashedMerchant: value }),
    setLoadingRestoreMerchant: (value) => set({ loadingRestoreMerchant: value }),
    setLoadingDeletePermanentMerchant: (value) => set({ loadingDeletePermanentMerchant: value }),

   
    setErrorGetMerchants: (value) => set({ errorGetMerchants: value }),
    setErrorGetMerchant: (value) => set({ errorGetMerchant: value }),
    setErrorGetApiKey: (value) => set({ errorGetApiKey: value }),
    setErrorGetMerchantUser: (value) => set({ errorGetMerchantUser: value }),
    setErrorGetActiveMerchant: (value) => set({ errorGetActiveMerchant: value }),
    setErrorGetTrashedMerchant: (value) => set({ errorGetTrashedMerchant: value }),
    setErrorCreateMerchant: (value) => set({ errorCreateMerchant: value }),
    setErrorUpdateMerchant: (value) => set({ errorUpdateMerchant: value }),
    setErrorTrashedMerchant: (value) => set({ errorTrashedMerchant: value }),
    setErrorRestoreMerchant: (value) => set({ errorRestoreMerchant: value }),
    setErrorDeletePermanentMerchant: (value: string | null) => set({ errorDeletePermanentMerchant: value }),

    // Methods
    findAllMerchants: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetMerchants: true, errorGetMerchants: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/merchants", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                merchants: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetMerchants: false,
                errorGetMerchants: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetMerchants: false }),
                (message: any) => set({ errorGetMerchants: message })
            );
        }
    },

    findById: async (id: number) => {
        set({ loadingGetMerchant: true, errorGetMerchant: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/merchants/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ merchant: response.data, loadingGetMerchant: false, errorGetMerchant: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetMerchant: false }),
                (message: any) => set({ errorGetMerchant: message })
            );
        }
    },

    findByApiKey: async (api_key: string) => {
        set({ loadingGetApiKey: true, errorGetApiKey: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/merchants/api-key/${api_key}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ merchant: response.data, loadingGetApiKey: false, errorGetApiKey: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetApiKey: false }),
                (message: any) => set({ errorGetApiKey: message })
            );
        }
    },

    findByMerchantUser: async (user_id: number) => {
        set({ loadingGetMerchantUser: true, errorGetMerchantUser: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/merchants/user/${user_id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ merchants: response.data, loadingGetMerchantUser: false, errorGetMerchantUser: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetMerchantUser: false }),
                (message: any) => set({ errorGetMerchantUser: message })
            );
        }
    },

    findByActive: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetActiveMerchant: true, errorGetActiveMerchant: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/merchants/active", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                merchants: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetActiveMerchant: false,
                errorGetActiveMerchant: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetActiveMerchant: false }),
                (message: any) => set({ errorGetActiveMerchant: null })
            );
        }
    },

    findByTrashed: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTrashedMerchant: true, errorGetTrashedMerchant: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/merchants/trashed", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                merchants: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTrashedMerchant: false,
                errorGetTrashedMerchant: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTrashedMerchant: false }),
                (message: any) => set({ errorGetTrashedMerchant: message })
            );
        }
    },

    createMerchant: async (req: CreateMerchant) => {
        set({ loadingCreateMerchant: true, errorCreateMerchant: null });
        try {
            const token = getAccessToken();
            const response = await myApi.post("/merchants", req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingCreateMerchant: false, errorCreateMerchant: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingCreateMerchant: false }),
                (message: any) => set({ errorCreateMerchant: message })
            );
        }
    },

    updateMerchant: async (id: number, req: UpdateMerchant) => {
        set({ loadingUpdateMerchant: true, errorUpdateMerchant: null });
        try {
            const token = getAccessToken();
            const response = await myApi.put(`/merchants/${id}`, req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingUpdateMerchant: false, errorUpdateMerchant: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingUpdateMerchant: false }),
                (message: any) => set({ errorUpdateMerchant: message })
            );
        }
    },

    restoreMerchant: async (id: number) => {
        set({ loadingRestoreMerchant: true, errorRestoreMerchant: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/merchants/restore/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingRestoreMerchant: false, errorRestoreMerchant: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingRestoreMerchant: false }),
                (message: any) => set({ errorRestoreMerchant: null })
            );
        }
    },

    trashedMerchant: async (id: number) => {
        set({ loadingTrashedMerchant: true, errorTrashedMerchant: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/merchants/trashed/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingTrashedMerchant: false, errorTrashedMerchant: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingTrashedMerchant: false }),
                (message: any) => set({ errorTrashedMerchant: message })
            );
        }
    },

    deleteMerchantPermanent: async (id: number) => {
        set({ loadingDeletePermanentMerchant: true, errorDeletePermanentMerchant: null });
        try {
            const token = getAccessToken();
            const response = await myApi.delete(`/merchants/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingDeletePermanentMerchant: false, errorDeletePermanentMerchant: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingDeletePermanentMerchant: false }),
                (message: any) => set({ errorDeletePermanentMerchant: message })
            );
        }
    },
}));

export default useMerchantStore;