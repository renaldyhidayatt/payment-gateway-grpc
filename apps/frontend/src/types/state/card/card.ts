import { CreateCard, UpdateCard } from "../../domain/request";
import { Card } from "../../model/card";

export interface CardStore {
  cards: Card[] | null;
  card: Card | null;

  pagination: {
    currentPage: number;
    pageSize: number;
    totalItems: number;
    totalPages: number;
  };

  loadingGetCards: boolean;
  loadingGetCard: boolean;
  loadingGetCardByUser: boolean;
  loadingGetActiveCards: boolean;
  loadingGetTrashedCards: boolean;
  loadingGetCardByCardNumber: boolean;

  loadingCreateCard: boolean;
  loadingUpdateCard: boolean;
  loadingTrashedCard: boolean;
  loadingRestoreCard: boolean;
  loadingDeleteCard: boolean;

  errorGetCards: string | null;
  errorGetCard: string | null;
  errorGetCardByUser: string | null;
  errorGetActiveCards: string | null;
  errorGetTrashedCards: string | null;
  errorGetCardByCardNumber: string | null;

  errorCreateCard: string | null;
  errorUpdateCard: string | null;
  errorTrashedCard: string | null;
  errorRestoreCard: string | null;
  errorDeleteCard: string | null;

  setErrorGetCards: (value: string | null) => void;
  setErrorGetCard: (value: string | null) => void;
  setErrorGetCardByUser: (value: string | null) => void;
  setErrorGetActiveCards: (value: string | null) => void;
  setErrorGetTrashedCards: (value: string | null) => void;
  setErrorGetCardByCardNumber: (value: string | null) => void;

  setErrorCreateCard: (value: string | null) => void;
  setErrorUpdateCard: (value: string | null) => void;
  setErrorTrashedCard: (value: string | null) => void;
  setErrorRestoreCard: (value: string | null) => void;
  setErrorDeleteCard: (value: string | null) => void;

  setLoadingGetCards: (value: boolean) => void;
  setLoadingGetCard: (value: boolean) => void;
  setLoadingGetCardByUser: (value: boolean) => void;
  setLoadingGetActiveCards: (value: boolean) => void;
  setLoadingGetTrashedCards: (value: boolean) => void;
  setLoadingGetCardByCardNumber: (value: boolean) => void;

  setLoadingCreateCard: (value: boolean) => void;
  setLoadingUpdateCard: (value: boolean) => void;
  setLoadingTrashedCard: (value: boolean) => void;
  setLoadingRestoreCard: (value: boolean) => void;
  setLoadingDeleteCard: (value: boolean) => void;

  findAllCards: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByIdCard: (id: number) => Promise<void>;

  findByUser: (
    userId: number,
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByCardNumber: (cardNumber: string) => Promise<void>;

  findByActiveCard: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByTrashedCard: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;

  createCard: (req: CreateCard) => Promise<boolean>;
  updateCard: (id: number, req: UpdateCard) => Promise<boolean>;
  trashedCard: (id: number) => Promise<boolean>;
  restoreCard: (id: number) => Promise<boolean>;
  deletePermanentCard: (id: number) => Promise<boolean>;
}
