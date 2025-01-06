import { Transfer } from "../../model/transfer";

export interface TransferStore {
  transfers: Transfer[] | null;
  transfer: Transfer | null;

  pagination: {
    currentPage: number;
    pageSize: number;
    totalItems: number;
    totalPages: number;
  };

  loadingGetTransfers: boolean;
  loadingGetTransfer: boolean;
  loadingGetTransferFrom: boolean;
  loadingGetTransferTo: boolean;
  loadingGetActiveTransfer: boolean;
  loadingGetTrashedTransfer: boolean;

  loadingCreateTransfer: boolean;
  loadingUpdateTransfer: boolean;
  loadingTrashedTransfer: boolean;
  loadingRestoreTransfer: boolean;
  loadingPermanentTransfer: boolean;

  errorGetTransfers: string | null;
  errorGetTransfer: string | null;
  errorGetTransferFrom: string | null;
  errorGetTransferTo: string | null;
  errorGetActiveTransfer: string | null;
  errorGetTrashedTransfer: string | null;

  errorCreateTransfer: string | null;
  errorUpdateTransfer: string | null;
  errorTrashedTransfer: string | null;
  errorRestoreTransfer: string | null;
  errorPermanentTransfer: string | null;

  setLoadingGetTransfers: (value: boolean) => void;
  setLoadingGetTransfer: (value: boolean) => void;
  setLoadingGetTransferFrom: (value: boolean) => void;
  setLoadingGetTransferTo: (value: boolean) => void;
  setLoadingGetActiveTransfer: (value: boolean) => void;
  setLoadingGetTrashedTransfer: (value: boolean) => void;

  setLoadingCreateTransfer: (value: boolean) => void;
  setLoadingUpdateTransfer: (value: boolean) => void;
  setLoadingTrashedTransfer: (value: boolean) => void;
  setLoadingRestoreTransfer: (value: boolean) => void;
  setLoadingPermanentTransfer: (value: boolean) => void;

  setErrorGetTransfers: (value: string | null) => void;
  setErrorGetTransfer: (value: string | null) => void;
  setErrorGetTransferFrom: (value: string | null) => void;
  setErrorGetTransferTo: (value: string | null) => void;
  setErrorGetActiveTransfer: (value: string | null) => void;
  setErrorGetTrashedTransfer: (value: string | null) => void;

  setErrorCreateTransfer: (value: string | null) => void;
  setErrorUpdateTransfer: (value: string | null) => void;
  setErrorTrashedTransfer: (value: string | null) => void;
  setErrorRestoreTransfer: (value: string | null) => void;
  setErrorPermanentTransfer: (value: string | null) => void;

  findAllTransfers: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByIdTransfer: (id: number) => Promise<void>;
  findByTransferFrom: (fromAccountId: number) => Promise<void>;
  findByTransferTo: (toAccountId: number) => Promise<void>;
  findByActiveTransfer: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByTrashedTransfer: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;

  createTransfer: (req: any) => Promise<boolean>;
  updateTransfer: (id: number, req: any) => Promise<boolean>;
  restoreTransfer: (id: number) => Promise<boolean>;
  trashedTransfer: (id: number) => Promise<boolean>;
  deleteTransferPermanent: (id: number) => Promise<boolean>;
}
