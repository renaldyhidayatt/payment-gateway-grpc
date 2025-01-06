import { CreateUser, UpdateUser } from "@/types/domain/request";
import { User } from "@/types/model/user";

export interface UserStore {
  users: User[] | null;
  user: User | null;

  pagination: {
    currentPage: number;
    pageSize: number;
    totalItems: number;
    totalPages: number;
  };
  loadingGetUsers: boolean;
  loadingGetUser: boolean;
  loadingGetActiveUsers: boolean;
  loadingGetTrashedUsers: boolean;

  loadingCreateUser: boolean;
  loadingUpdateUser: boolean;
  loadingTrashedUser: boolean;
  loadingRestoreUser: boolean;
  loadingDeleteUser: boolean;

  errorGetUsers: string | null;
  errorGetUser: string | null;
  errorGetActiveUsers: string | null;
  errorGetTrashedUsers: string | null;
  errorCreateUser: string | null;
  errorUpdateUser: string | null;
  errorTrashedUser: string | null;
  errorRestoreUser: string | null;
  errorDeleteUser: string | null;

  setErrorGetUsers: (value: string | null) => void;
  setErrorGetUser: (value: string | null) => void;
  setErrorGetActiveUsers: (value: string | null) => void;
  setErrorGetTrashedUsers: (value: string | null) => void;

  setErrorCreateUser: (value: string | null) => void;
  setErrorUpdateUser: (value: string | null) => void;
  setErrorTrashedUser: (value: string | null) => void;
  setErrorRestoreUser: (value: string | null) => void;
  setErrorDeleteUser: (value: string | null) => void;

  setLoadingGetUsers: (value: boolean) => void;
  setLoadingGetUser: (value: boolean) => void;
  setLoadingGetActiveUsers: (value: boolean) => void;
  setLoadingGetTrashedUsers: (value: boolean) => void;

  setLoadingCreateUser: (value: boolean) => void;
  setLoadingUpdateUser: (value: boolean) => void;
  setLoadingTrashedUser: (value: boolean) => void;
  setLoadingRestoreUser: (value: boolean) => void;
  setLoadingDeleteUser: (value: boolean) => void;

  findAllUsers: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findById: (id: number) => Promise<void>;
  findByActive: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;
  findByTrashed: (
    search: string,
    page: number,
    pageSize: number,
  ) => Promise<void>;

  createUser: (req: CreateUser) => Promise<boolean>;
  updateUser: (id: number, req: UpdateUser) => Promise<boolean>;
  restoreUser: (id: number) => Promise<boolean>;
  trashedUser: (id: number) => Promise<boolean>;
  deleteUserPermanent: (id: number) => Promise<boolean>;
}
