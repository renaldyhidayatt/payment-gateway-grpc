import { Saldo } from "@/types/model/saldo";
import { CreateSaldo, UpdateSaldo } from "@/types/domain/request";

export interface SaldoStore {
  saldos: Saldo[] | null;
  saldo: Saldo | null;

  pagination: {
    currentPage: number;
    pageSize: number;
    totalItems: number;
    totalPages: number;
  };

  loadingGetSaldos: boolean;
  loadingGetSaldo: boolean;
  loadingGetActiveSaldo: boolean;
  loadingGetTrashedSaldo: boolean;
  loadingGetCardNumberSaldo: boolean;

  loadingCreateSaldo: boolean;
  loadingUpdateSaldo: boolean;
  loadingTrashedSaldo: boolean;
  loadingRestoreSaldo: boolean;
  loadingDeletePermanent: boolean;

  errorGetSaldos: string | null;
  errorGetSaldo: string | null;
  errorGetActiveSaldo: string | null;
  errorGetTrashedSaldo: string | null;
  errorGetCardNumberSaldo: string | null;

  errorCreateSaldo: string | null;
  errorUpdateSaldo: string | null;
  errorTrashedSaldo: string | null;
  errorRestoreSaldo: string | null;
  errorDeletePermanent: string | null;

  setLoadingGetSaldos: (value: boolean) => void;
  setLoadingGetSaldo: (value: boolean) => void;
  setLoadingGetActiveSaldo: (value: boolean) => void;
  setLoadingGetTrashedSaldo: (value: boolean) => void;
  setLoadingGetCardNumberSaldo: (value: boolean) => void;

  setLoadingCreateSaldo: (value: boolean) => void;
  setLoadingUpdateSaldo: (value: boolean) => void;
  setLoadingTrashedSaldo: (value: boolean) => void;
  setLoadingRestoreSaldo: (value: boolean) => void;
  setLoadingDeletePermanent: (value: boolean) => void;

  setErrorGetSaldos: (value: string | null) => void;
  setErrorGetSaldo: (value: string | null) => void;
  setErrorGetActiveSaldo: (value: string | null) => void;
  setErrorGetTrashedSaldo: (value: string | null) => void;
  setErrorGetCardNumberSaldo: (value: string | null) => void;

  setErrorCreateSaldo: (value: string | null) => void;
  setErrorUpdateSaldo: (value: string | null) => void;
  setErrorTrashedSaldo: (value: string | null) => void;
  setErrorRestoreSaldo: (value: string | null) => void;
  setErrorDeletePermanent: (value: string | null) => void;

  findAllSaldos: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByIdSaldo: (id: number) => Promise<void>;
  findByActiveSaldo: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByTrashedSaldo: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByCardNumberSaldo: (cardNumber: string) => Promise<void>;

  createSaldo: (req: CreateSaldo) => Promise<boolean>;
  updateSaldo: (id: number, req: UpdateSaldo) => Promise<boolean>;
  restoreSaldo: (id: number) => Promise<boolean>;
  trashedSaldo: (id: number) => Promise<boolean>;
  deleteSaldoPermanent: (id: number) => Promise<boolean>;
}
