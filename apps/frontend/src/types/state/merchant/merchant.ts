import {
    CreateMerchant,
    UpdateMerchant
} from "../../domain/request"
import {
    Merchant
} from "../../model/merchant"

export interface MerchantStore{
    merchants: Merchant[] | null;
    merchant: Merchant | null;

    pagination: {
        currentPage: number;
        pageSize: number;
        totalItems: number;
        totalPages: number;
    }

    loadingGetMerchants: boolean;
    loadingGetMerchant: boolean;
    loadingGetApiKey: boolean;
    loadingGetMerchantUser: boolean;
    loadingGetActiveMerchant: boolean;
    loadingGetTrashedMerchant: boolean;

    loadingCreateMerchant: boolean;
    loadingUpdateMerchant: boolean;
    loadingTrashedMerchant: boolean;
    loadingRestoreMerchant: boolean;
    loadingDeletePermanentMerchant: boolean;

    

    errorGetMerchants: string | null;
    errorGetMerchant: string | null;
    errorGetApiKey: string | null;
    errorGetMerchantUser: string | null;
    errorGetActiveMerchant: string | null;
    errorGetTrashedMerchant: string | null;

    errorCreateMerchant: string | null;
    errorUpdateMerchant: string | null;
    errorTrashedMerchant: string | null;
    errorRestoreMerchant: string | null;
    errorDeletePermanentMerchant: string | null;

    setLoadingGetMerchants: (value: boolean) => void;
    setLoadingGetMerchant: (value: boolean) => void;
    setLoadingGetApiKey: (value: boolean) => void;
    setLoadingGetMerchantUser: (value: boolean) => void;
    setLoadingGetActiveMerchant: (value: boolean) => void;
    setLoadingGetTrashedMerchant: (value: boolean) => void;
    setLoadingCreateMerchant: (value: boolean) => void;
    setLoadingUpdateMerchant: (value: boolean) => void;
    setLoadingTrashedMerchant: (value: boolean) => void;
    setLoadingRestoreMerchant: (value: boolean) => void;
    setLoadingDeletePermanentMerchant: (value: boolean) => void;

    
    setErrorGetMerchants: (value: string | null) => void;
    setErrorGetMerchant: (value: string | null) => void;
    setErrorGetApiKey: (value: string | null) => void;
    setErrorGetMerchantUser: (value: string | null) => void;
    setErrorGetActiveMerchant: (value: string | null) => void;
    setErrorGetTrashedMerchant: (value: string | null) => void;
    setErrorCreateMerchant: (value: string | null) => void;
    setErrorUpdateMerchant: (value: string | null) => void;
    setErrorTrashedMerchant: (value: string | null) => void;
    setErrorRestoreMerchant: (value: string | null) => void;
    setErrorDeletePermanentMerchant: (value: string | null) => void;

    findAllMerchants: (search: string, page: number, pageSize: number) => Promise<void>;
    findById: (id: number) => Promise<void>;
    findByApiKey: (api_key: string) => Promise<void>;
    findByMerchantUser: (user_id: number) => Promise<void>;
    findByActive: (search: string, page: number, pageSize: number) => Promise<void>;
    findByTrashed: (search: string, page: number, pageSize: number) => Promise<void>;
    createMerchant: (req: CreateMerchant) => Promise<void>;
    updateMerchant: (id: number, req: UpdateMerchant) => Promise<void>;
    restoreMerchant: (id: number) => Promise<void>;
    trashedMerchant:  (id: number) => Promise<void>;
    deleteMerchantPermanent: (id: number) => Promise<void>;


}