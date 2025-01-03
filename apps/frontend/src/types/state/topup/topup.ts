import { Topup } from "../../model/topup";

export interface TopupStore {
    topups: Topup[] | null;
    topup: Topup | null;

    pagination: {
        currentPage: number;
        pageSize: number;
        totalItems: number;
        totalPages: number;
    }

    loadingGetTopups: boolean;
    loadingGetTopup: boolean;
    loadingGetActiveTopup: boolean;
    loadingGetTrashedTopup: boolean;
    loadingGetCardNumberTopup: boolean;

    loadingCreateTopup: boolean;
    loadingUpdateTopup: boolean;
    loadingTrashedTopup: boolean;
    loadingRestoreTopup: boolean;
    loadingPermanentTopup: boolean;


    errorGetTopups: string | null;
    errorGetTopup: string | null;
    errorGetActiveTopup: string | null;
    errorGetTrashedTopup: string | null;
    errorGetCardNumberTopup: string | null;

    errorCreateTopup: string | null;
    errorUpdateTopup: string | null;
    errorTrashedTopup: string | null;
    errorRestoreTopup: string | null;
    errorPermanentTopup: string | null;

    setLoadingGetTopups: (value: boolean) => void;
    setLoadingGetTopup: (value: boolean) => void;
    setLoadingGetActiveTopup: (value: boolean) => void;
    setLoadingGetTrashedTopup: (value: boolean) => void;
    setLoadingGetCardNumberTopup: (value: boolean) => void;

    setLoadingCreateTopup: (value: boolean) => void;
    setLoadingUpdateTopup: (value: boolean) => void;
    setLoadingTrashedTopup: (value: boolean) => void;
    setLoadingRestoreTopup: (value: boolean) => void;
    setLoadingPermanentTopup: (value: boolean) => void;

    setErrorGetTopups: (value: string | null) => void;
    setErrorGetTopup: (value: string | null) => void;
    setErrorGetActiveTopup: (value: string | null) => void;
    setErrorGetTrashedTopup: (value: string | null) => void;
    setErrorGetCardNumberTopup: (value: string | null) => void;

    setErrorCreateTopup: (value: string | null) => void;
    setErrorUpdateTopup: (value: string | null) => void;
    setErrorTrashedTopup: (value: string | null) => void;
    setErrorRestoreTopup: (value: string | null) => void;
    setErrorPermanentTopup: (value: string | null) => void;

   
    findAllTopups: (search: string, page: number, pageSize: number) => Promise<void>;
    findByIdTopup: (id: number) => Promise<void>;
    findByActiveTopup: (search: string, page: number, pageSize: number) => Promise<void>;
    findByTrashedTopup: (search: string, page: number, pageSize: number) => Promise<void>;
    findByCardNumberTopup: (cardNumber: string) => Promise<void>;
    createTopup: (req: any) => Promise<void>;
    updateTopup: (id: number, req: any) => Promise<void>;
    restoreTopup: (id: number) => Promise<void>;
    trashedTopup: (id: number) => Promise<void>;
    deleteTopupPermanent: (id: number) => Promise<void>;
}