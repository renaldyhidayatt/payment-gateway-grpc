import { create } from "zustand";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";
import { CardStore } from "@/types/state/card/card";
import { CreateCard, UpdateCard } from "@/types/domain/request";
import { getAccessToken } from "../auth";

const useCardStore = create<CardStore>((set, get) => ({
  cards: null,
  card: null,

  pagination: {
    currentPage: 1,
    pageSize: 10,
    totalItems: 0,
    totalPages: 0,
  },

  loadingGetCards: false,
  loadingGetCard: false,
  loadingGetCardByUser: false,
  loadingGetActiveCards: false,
  loadingGetTrashedCards: false,
  loadingGetCardByCardNumber: false,
  loadingCreateCard: false,
  loadingUpdateCard: false,
  loadingTrashedCard: false,
  loadingRestoreCard: false,
  loadingDeleteCard: false,

  errorGetCards: null,
  errorGetCard: null,
  errorGetCardByUser: null,
  errorGetActiveCards: null,
  errorGetTrashedCards: null,
  errorGetCardByCardNumber: null,
  errorCreateCard: null,
  errorUpdateCard: null,
  errorTrashedCard: null,
  errorRestoreCard: null,
  errorDeleteCard: null,

  setErrorGetCards: (value: string | null) => set({ errorGetCards: value }),
  setErrorGetCard: (value: string | null) => set({ errorGetCard: value }),
  setErrorGetCardByUser: (value: string | null) =>
    set({ errorGetCardByUser: value }),
  setErrorGetActiveCards: (value: string | null) =>
    set({ errorGetActiveCards: value }),
  setErrorGetTrashedCards: (value: string | null) =>
    set({ errorGetTrashedCards: value }),
  setErrorGetCardByCardNumber: (value: string | null) =>
    set({ errorGetCardByCardNumber: value }),
  setErrorCreateCard: (value: string | null) => set({ errorCreateCard: value }),
  setErrorUpdateCard: (value: string | null) => set({ errorUpdateCard: value }),
  setErrorTrashedCard: (value: string | null) =>
    set({ errorTrashedCard: value }),
  setErrorRestoreCard: (value: string | null) =>
    set({ errorRestoreCard: value }),
  setErrorDeleteCard: (value: string | null) => set({ errorDeleteCard: value }),

  setLoadingGetCards: (value: boolean) => set({ loadingGetCards: value }),
  setLoadingGetCard: (value: boolean) => set({ loadingGetCard: value }),
  setLoadingGetCardByUser: (value: boolean) =>
    set({ loadingGetCardByUser: value }),
  setLoadingGetActiveCards: (value: boolean) =>
    set({ loadingGetActiveCards: value }),
  setLoadingGetTrashedCards: (value: boolean) =>
    set({ loadingGetTrashedCards: value }),
  setLoadingGetCardByCardNumber: (value: boolean) =>
    set({ loadingGetCardByCardNumber: value }),
  setLoadingCreateCard: (value: boolean) => set({ loadingCreateCard: value }),
  setLoadingUpdateCard: (value: boolean) => set({ loadingUpdateCard: value }),
  setLoadingTrashedCard: (value: boolean) => set({ loadingTrashedCard: value }),
  setLoadingRestoreCard: (value: boolean) => set({ loadingRestoreCard: value }),
  setLoadingDeleteCard: (value: boolean) => set({ loadingDeleteCard: value }),

  findAllCards: async (search: string, page: number, pageSize: number) => {
    set({ loadingGetCards: true, errorGetCards: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get("/card", {
        params: { page, pageSize, search },
        headers: { Authorization: `Bearer ${token}` },
      });
      set({
        cards: response.data.data,
        pagination: {
          currentPage: response.data.pagination.current_page,
          pageSize: response.data.pagination.page_size,
          totalItems: response.data.pagination.total_records,
          totalPages: response.data.pagination.total_pages,
        },
        loadingGetCards: false,
        errorGetCards: null,
      });
      return response.data;
    } catch (error) {
      handleApiError(
        error,
        () => set({ loadingGetCards: false }),
        (message: any) => set({ errorGetCards: message }),
      );
    }
  },

  findByIdCard: async (id: number) => {
    set({ loadingGetCard: true, errorGetCard: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get(`/cards/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      set({ card: response.data, loadingGetCard: false, errorGetCard: null });

      return response.data;
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingGetCard: false }),
        (message: any) => set({ errorGetCard: message }),
      );
    }
  },

  findByUser: async (id: number) => {
    set({ loadingGetCardByUser: true, errorGetCardByUser: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get(`/cards/user/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      set({
        cards: response.data,
        loadingGetCardByUser: false,
        errorGetCardByUser: null,
      });

      return response.data;
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingGetCardByUser: false }),
        (message: any) => set({ errorGetCardByUser: message }),
      );
    }
  },

  findByCardNumber: async (cardNumber: string) => {
    set({ loadingGetCardByCardNumber: true, errorGetCardByCardNumber: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get(`/cards/number/${cardNumber}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      set({
        card: response.data,
        loadingGetCardByCardNumber: false,
        errorGetCardByCardNumber: null,
      });
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingGetCardByCardNumber: false }),
        (message: any) => set({ errorGetCardByCardNumber: message }),
      );
    }
  },

  findByActiveCard: async (search: string, page: number, pageSize: number) => {
    set({ loadingGetActiveCards: true, errorGetActiveCards: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get("/cards/active", {
        params: { page, pageSize, search },
        headers: { Authorization: `Bearer ${token}` },
      });
      set({
        cards: response.data.items,
        pagination: {
          currentPage: response.data.currentPage,
          pageSize: response.data.pageSize,
          totalItems: response.data.totalItems,
          totalPages: response.data.totalPages,
        },
        loadingGetActiveCards: false,
        errorGetActiveCards: null,
      });
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingGetActiveCards: false }),
        (message: any) => set({ errorGetActiveCards: message }),
      );
    }
  },

  findByTrashedCard: async (search: string, page: number, pageSize: number) => {
    set({ loadingGetTrashedCards: true, errorGetTrashedCards: null });
    try {
      const token = getAccessToken();
      const response = await myApi.get("/cards/trashed", {
        params: { page, pageSize, search },
        headers: { Authorization: `Bearer ${token}` },
      });
      set({
        cards: response.data.items,
        pagination: {
          currentPage: response.data.currentPage,
          pageSize: response.data.pageSize,
          totalItems: response.data.totalItems,
          totalPages: response.data.totalPages,
        },
        loadingGetTrashedCards: false,
        errorGetTrashedCards: null,
      });
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingGetTrashedCards: false }),
        (message: any) => set({ errorGetTrashedCards: message }),
      );
    }
  },

  createCard: async (req: CreateCard) => {
    set({ loadingCreateCard: true, errorCreateCard: null });
    try {
      const token = getAccessToken();
      const response = await myApi.post("/cards", req, {
        headers: { Authorization: `Bearer ${token}` },
      });
      set({
        card: response.data,
        loadingCreateCard: false,
        errorCreateCard: null,
      });
      return response.data;
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingCreateCard: false }),
        (message: any) => set({ errorCreateCard: message }),
      );
    }
  },

  updateCard: async (id: number, req: UpdateCard) => {
    set({ loadingUpdateCard: true, errorUpdateCard: null });
    try {
      const token = getAccessToken();
      const response = await myApi.put(`/cards/${id}`, req, {
        headers: { Authorization: `Bearer ${token}` },
      });
      set({
        card: response.data,
        loadingUpdateCard: false,
        errorUpdateCard: null,
      });
      return response.data;
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingUpdateCard: false }),
        (message: any) => set({ errorUpdateCard: message }),
      );
    }
  },

  restoreCard: async (id: number) => {
    set({ loadingRestoreCard: true, errorRestoreCard: null });
    try {
      const token = getAccessToken();
      const response = await myApi.patch(`/cards/restore/${id}`, null, {
        headers: { Authorization: `Bearer ${token}` },
      });
      set({ loadingRestoreCard: false, errorRestoreCard: null });
      return response.data;
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingRestoreCard: false }),
        (message: any) => set({ errorRestoreCard: message }),
      );
    }
  },

  trashedCard: async (id: number) => {
    set({ loadingTrashedCard: true, errorTrashedCard: null });
    try {
      const token = getAccessToken();
      const response = await myApi.patch(`/cards/trashed/${id}`, null, {
        headers: { Authorization: `Bearer ${token}` },
      });
      set({ loadingTrashedCard: false, errorTrashedCard: null });
      return response.data;
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingTrashedCard: false }),
        (message: any) => set({ errorTrashedCard: message }),
      );
    }
  },

  deletePermanentCard: async (id: number) => {
    set({ loadingDeleteCard: true, errorDeleteCard: null });
    try {
      const token = getAccessToken();
      const response = await myApi.delete(`/cards/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      set({ loadingDeleteCard: false, errorDeleteCard: null });
      return response.data;
    } catch (err) {
      handleApiError(
        err,
        () => set({ loadingDeleteCard: false }),
        (message: any) => set({ errorDeleteCard: message }),
      );
    }
  },
}));

export default useCardStore;
