import { SaldoStore } from "@/types/state/saldo";
import { create } from "zustand";
import { getAccessToken } from "../auth";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";
import { CreateSaldo, UpdateSaldo } from "@/types/domain/request";

const useSaldoStore = create<SaldoStore>((set, get) => ({
    saldos: null,
    saldo: null,

    pagination: {
        currentPage: 1,
        pageSize: 10,
        totalItems: 0,
        totalPages: 0,
    },

    loadingGetSaldos: false,
    loadingGetSaldo: false,
    loadingGetActiveSaldo: false,
    loadingGetTrashedSaldo: false,
    loadingGetCardNumberSaldo: false,

    loadingCreateSaldo: false,
    loadingUpdateSaldo: false,
    loadingTrashedSaldo: false,
    loadingRestoreSaldo: false,
    loadingDeletePermanent: false,

    
    errorGetSaldos: null,
    errorGetSaldo: null,
    errorGetActiveSaldo: null,
    errorGetTrashedSaldo: null,
    errorGetCardNumberSaldo: null,

    errorCreateSaldo: null,
    errorUpdateSaldo: null,
    errorTrashedSaldo: null,
    errorRestoreSaldo: null,
    errorDeletePermanent: null,

    setLoadingGetSaldos: (value) => set({ loadingGetSaldos: value }),
    setLoadingGetSaldo: (value) => set({ loadingGetSaldo: value }),
    setLoadingGetActiveSaldo: (value) => set({ loadingGetActiveSaldo: value }),
    setLoadingGetTrashedSaldo: (value) => set({ loadingGetTrashedSaldo: value }),
    setLoadingGetCardNumberSaldo: (value) => set({ loadingGetCardNumberSaldo: value }),

    setLoadingCreateSaldo: (value) => set({ loadingCreateSaldo: value }),
    setLoadingUpdateSaldo: (value) => set({ loadingUpdateSaldo: value }),
    setLoadingTrashedSaldo: (value) => set({ loadingTrashedSaldo: value }),
    setLoadingRestoreSaldo: (value) => set({ loadingRestoreSaldo: value }),
    setLoadingDeletePermanent: (value) => set({ loadingDeletePermanent: value }),

    setErrorGetSaldos: (value) => set({ errorGetSaldos: value }),
    setErrorGetSaldo: (value) => set({ errorGetSaldo: value }),
    setErrorGetActiveSaldo: (value) => set({ errorGetActiveSaldo: value }),
    setErrorGetTrashedSaldo: (value) => set({ errorGetTrashedSaldo: value }),
    setErrorGetCardNumberSaldo: (value) => set({ errorGetCardNumberSaldo: value }),

    setErrorCreateSaldo: (value) => set({ errorCreateSaldo: value }),
    setErrorUpdateSaldo: (value) => set({ errorUpdateSaldo: value }),
    setErrorTrashedSaldo: (value) => set({ errorTrashedSaldo: value }),
    setErrorRestoreSaldo: (value) => set({ errorRestoreSaldo: value }),
    setErrorDeletePermanent: (value) => set({ errorDeletePermanent: value }),

    
    findAllSaldos: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetSaldos: true, errorGetSaldos: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/saldos", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                saldos: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetSaldos: false,
                errorGetSaldos: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetSaldos: false }),
                (message: any) => set({ errorGetSaldos: message })
            );
        }
    },

    findByIdSaldo: async (id: number) => {
        set({ loadingGetSaldo: true, errorGetSaldo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/saldos/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ saldo: response.data, loadingGetSaldo: false, errorGetSaldo: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetSaldo: false }),
                (message: any) => set({ errorGetSaldo: message })
            );
        }
    },

    findByActiveSaldo: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetActiveSaldo: true, errorGetActiveSaldo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/saldos/active", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                saldos: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetActiveSaldo: false,
                errorGetActiveSaldo: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetActiveSaldo: false }),
                (message: any) => set({ errorGetActiveSaldo: message })
            );
        }
    },

    findByTrashedSaldo: async (search: string, page: number, pageSize: number) => {
        set({ loadingGetTrashedSaldo: true, errorGetTrashedSaldo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get("/saldos/trashed", {
                params: { search, page, pageSize },
                headers: { Authorization: `Bearer ${token}` },
            });
            set({
                saldos: response.data.items,
                pagination: {
                    currentPage: response.data.currentPage,
                    pageSize: response.data.pageSize,
                    totalItems: response.data.totalItems,
                    totalPages: response.data.totalPages,
                },
                loadingGetTrashedSaldo: false,
                errorGetTrashedSaldo: null,
            });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetTrashedSaldo: false }),
                (message: any) => set({ errorGetTrashedSaldo: message })
            );
        }
    },

    findByCardNumberSaldo: async (cardNumber: string) => {
        set({ loadingGetCardNumberSaldo: true, errorGetCardNumberSaldo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.get(`/saldos/card-number/${cardNumber}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ saldo: response.data, loadingGetCardNumberSaldo: false, errorGetCardNumberSaldo: null });
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingGetCardNumberSaldo: false }),
                (message: any) => set({ errorGetCardNumberSaldo: message })
            );
        }
    },

    createSaldo: async (req: CreateSaldo) => {
        set({ loadingCreateSaldo: true, errorCreateSaldo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.post("/saldos", req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingCreateSaldo: false, errorCreateSaldo: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingCreateSaldo: false }),
                (message: any) => set({ errorCreateSaldo: message })
            );
        }
    },

    updateSaldo: async (id: number, req: UpdateSaldo) => {
        set({ loadingUpdateSaldo: true, errorUpdateSaldo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.put(`/saldos/${id}`, req, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingUpdateSaldo: false, errorUpdateSaldo: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingUpdateSaldo: false }),
                (message: any) => set({ errorUpdateSaldo: message })
            );
        }
    },

    restoreSaldo: async (id: number) => {
        set({ loadingRestoreSaldo: true, errorRestoreSaldo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/saldos/restore/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingRestoreSaldo: false, errorRestoreSaldo: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingRestoreSaldo: false }),
                (message: any) => set({ errorRestoreSaldo: message })
            );
        }
    },

    trashedSaldo: async (id: number) => {
        set({ loadingTrashedSaldo: true, errorTrashedSaldo: null });
        try {
            const token = getAccessToken();
            const response = await myApi.patch(`/saldos/trashed/${id}`, null, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingTrashedSaldo: false, errorTrashedSaldo: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingTrashedSaldo: false }),
                (message: any) => set({ errorTrashedSaldo: message })
            );
        }
    },

    deleteSaldoPermanent: async (id: number) => {
        set({ loadingDeletePermanent: true, errorDeletePermanent: null });
        try {
            const token = getAccessToken();
            const response = await myApi.delete(`/saldos/${id}`, {
                headers: { Authorization: `Bearer ${token}` },
            });
            set({ loadingDeletePermanent: false, errorDeletePermanent: null });
            return response.data;
        } catch (err) {
            handleApiError(
                err,
                () => set({ loadingDeletePermanent: false }),
                (message: any) => set({ errorDeletePermanent: message })
            );
        }
    },
}));

export default useSaldoStore;