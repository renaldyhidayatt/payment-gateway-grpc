import { Withdraw } from "@/types/model/withdraw";


export interface WithdrawStore{
    withdraws: Withdraw[] | null;
    withdraw: Withdraw| null;

    pagination: {
        currentPage: number;
        pageSize: number;
        totalItems: number;
        totalPages: number;
    }

    loadingGetWithdraws: boolean;
    loadingGetWithdraw: boolean;
    loadingGetCardNumberWithdraw: boolean;
    loadingGetActiveWithdraw: boolean;
    loadingGetTrashedWithdraw: boolean;
    
    loadingCreateWithdraw: boolean;
    loadingUpdateWithdraw: boolean;
    loadingTrashedWithdraw: boolean;
    loadingRestoreWithdraw: boolean;
    loadingPermanentWithdraw: boolean;

    // Error states
    errorGetWithdraws: string | null;
    errorGetWithdraw: string | null;
    errorGetCardNumberWithdraw: string | null;
    errorGetActiveWithdraw: string | null;
    errorGetTrashedWithdraw: string | null;

    errorCreateWithdraw: string | null;
    errorUpdateWithdraw: string | null;
    errorTrashedWithdraw: string | null;
    errorRestoreWithdraw: string | null;
    errorPermanentWithdraw: string | null;

    
    setLoadingGetWithdraws: (value: boolean) => void;
    setLoadingGetWithdraw: (value: boolean) => void;
    setLoadingGetCardNumberWithdraw: (value: boolean) => void;
    setLoadingGetActiveWithdraw: (value: boolean) => void;
    setLoadingGetTrashedWithdraw: (value: boolean) => void;

    setLoadingCreateWithdraw: (value: boolean) => void;
    setLoadingUpdateWithdraw: (value: boolean) => void;
    setLoadingTrashedWithdraw: (value: boolean) => void;
    setLoadingRestoreWithdraw: (value: boolean) => void;
    setLoadingPermanentWithdraw: (value: boolean) => void;

    
    setErrorGetWithdraws: (value: string | null) => void;
    setErrorGetWithdraw: (value: string | null) => void;
    setErrorGetCardNumberWithdraw: (value: string | null) => void;
    setErrorGetActiveWithdraw: (value: string | null) => void;
    setErrorGetTrashedWithdraw: (value: string | null) => void;

    setErrorCreateWithdraw: (value: string | null) => void;
    setErrorUpdateWithdraw: (value: string | null) => void;
    setErrorTrashedWithdraw: (value: string | null) => void;
    setErrorRestoreWithdraw: (value: string | null) => void;
    setErrorPermanentWithdraw: (value: string | null) => void;

   
    findAllWithdraws: (search: string, page: number, pageSize: number) => Promise<void>;
    findByIdWithdraw: (id: number) => Promise<void>;
    findByCardNumberWithdraw: (cardNumber: string) => Promise<void>;
    findByActiveWithdraw: (search: string, page: number, pageSize: number) => Promise<void>;
    findByTrashedWithdraw: (search: string, page: number, pageSize: number) => Promise<void>;
    createWithdraw: (req: any) => Promise<void>;
    updateWithdraw: (id: number, req: any) => Promise<void>;
    restoreWithdraw: (id: number) => Promise<void>;
    trashedWithdraw: (id: number) => Promise<void>;
    deleteWithdrawPermanent: (id: number) => Promise<void>;
    
}   