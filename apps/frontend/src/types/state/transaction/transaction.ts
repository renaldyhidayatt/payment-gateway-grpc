import { CreateTransaction, UpdateTransaction } from "../../domain/request";
import { Transaction } from "../../model/transaction";

export interface TransactionStore {
  transactions: Transaction[] | null;
  transaction: Transaction | null;

  pagination: {
    currentPage: number;
    pageSize: number;
    totalItems: number;
    totalPages: number;
  };

  loadingGetTransactions: boolean;
  loadingGetTransaction: boolean;
  loadingGetCardNumberTransaction: boolean;
  loadingGetMerchantTransaction: boolean;
  loadingGetActiveTransaction: boolean;
  loadingGetTrashedTransaction: boolean;

  loadingCreateTransaction: boolean;
  loadingUpdateTransaction: boolean;
  loadingRestoreTransaction: boolean;
  loadingTrashedTransaction: boolean;
  loadingDeletePermanentTransaction: boolean;

  errorGetTransactions: string | null;
  errorGetTransaction: string | null;
  errorGetCardNumberTransaction: string | null;
  errorGetMerchantTransaction: string | null;
  errorGetActiveTransaction: string | null;
  errorGetTrashedTransaction: string | null;

  errorCreateTransaction: string | null;
  errorUpdateTransaction: string | null;
  errorRestoreTransaction: string | null;
  errorTrashedTransaction: string | null;
  errorDeletePermanentTransaction: string | null;

  setLoadingGetTransactions: (value: boolean) => void;
  setLoadingGetTransaction: (value: boolean) => void;
  setLoadingGetCardNumberTransaction: (value: boolean) => void;
  setLoadingGetMerchantTransaction: (value: boolean) => void;
  setLoadingGetActiveTransaction: (value: boolean) => void;
  setLoadingGetTrashedTransaction: (value: boolean) => void;

  setLoadingCreateTransaction: (value: boolean) => void;
  setLoadingUpdateTransaction: (value: boolean) => void;
  setLoadingRestoreTransaction: (value: boolean) => void;
  setLoadingTrashedTransaction: (value: boolean) => void;
  setLoadingDeletePermanentTransaction: (value: boolean) => void;

  setErrorGetTransactions: (value: string | null) => void;
  setErrorGetTransaction: (value: string | null) => void;
  setErrorGetCardNumberTransaction: (value: string | null) => void;
  setErrorGetMerchantTransaction: (value: string | null) => void;
  setErrorGetActiveTransaction: (value: string | null) => void;
  setErrorGetTrashedTransaction: (value: string | null) => void;

  setErrorCreateTransaction: (value: string | null) => void;
  setErrorUpdateTransaction: (value: string | null) => void;
  setErrorRestoreTransaction: (value: string | null) => void;
  setErrorTrashedTransaction: (value: string | null) => void;
  setErrorDeletePermanentTransaction: (value: string | null) => void;

  findAllTransactions: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByIdTransaction: (id: number) => Promise<void>;
  findByCardNumberTransaction: (cardNumber: string) => Promise<void>;
  findByMerchantTransaction: (merchantId: number) => Promise<void>;
  findByActiveTransaction: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByTrashedTransaction: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;

  createTransaction: (req: CreateTransaction) => Promise<boolean>;
  updateTransaction: (id: number, req: UpdateTransaction) => Promise<boolean>;
  restoreTransaction: (id: number) => Promise<boolean>;
  trashedTransaction: (id: number) => Promise<boolean>;
  deleteTransactionPermanent: (id: number) => Promise<boolean>;
}
