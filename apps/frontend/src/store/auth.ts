import { create } from "zustand";
import { persist } from "zustand/middleware";
import myApi from "@/helpers/api";
import { handleApiError } from "@/helpers/handleApi";
import { AuthStore } from "@/types/state/auth";
import { LoginRequest, RegisterRequest } from "@/types/domain/request";

const useAuthStore = create<AuthStore>()(
    persist<AuthStore>(
        (set, get) => ({
            user: null,
            accessToken: null,
            refreshToken: null,
            isAuthenticated: false,

            refreshTimer: null,

            loadingLogin: false,
            loadingRegister: false,
            loadingLogout: false,
            loadingGetMe: false,
            loadingRefreshAccessToken: false,

            errorLogin: null,
            errorRegister: null,
            errorLogout: null,
            errorGetMe: null,
            errorRefreshAccessToken: null,

            setLoadingLogin: (value: boolean) => set({ loadingLogin: value }),
            setLoadingRegister: (value: boolean) => set({ loadingRegister: value }),
            setLoadingLogout: (value: boolean) => set({ loadingLogout: value }),
            setLoadingGetMe: (value: boolean) => set({ loadingGetMe: value }),
            setLoadingRefreshAccessToken: (value: boolean) =>
                set({ loadingRefreshAccessToken: value }),

          
            setErrorLogin: (value: string | null) => set({ errorLogin: value }),
            setErrorRegister: (value: string | null) => set({ errorRegister: value }),
            setErrorLogout: (value: string | null) => set({ errorLogout: value }),
            setErrorGetMe: (value: string | null) => set({ errorGetMe: value }),
            setErrorRefreshAccessToken: (value: string | null) =>
                set({ errorRefreshAccessToken: value }),

            login: async (req: LoginRequest) => {
                set({
                    loadingLogin: true,
                    errorLogin: null,
                    isAuthenticated: false
                })

                try {
                    const response = await myApi.post("/auth/login", req);

                    if (response.status == 200) {
                        const { access_token, refresh_token } = response.data.data;

                        set({ accessToken: access_token, refreshToken: refresh_token, isAuthenticated: true });
                    }

                    const timer = setInterval(
                        () => get().refreshAccessToken?.(),
                        15 * 60 * 1000
                    )
                    
                    set({
                        refreshTimer: timer
                    })

                    return response.data.data;
                } catch (error) {
                    handleApiError(
                        error,
                        () => set({
                            loadingLogin: false,
                        }),
                        (messsage: any) => set({
                            errorLogin: messsage
                        })
                    )
                }
            },
            logout: async () => {
                set({
                    loadingLogout: true,
                    errorLogout: null,
                })
                try{
                    set({
                        user: null,
                        accessToken: null,
                        refreshToken: null,
                    });
                }catch(error: any){
                    set({
                        errorLogout: error?.message || "Failed to log out",
                    })
                }

            },

            register: async (req: RegisterRequest) => {
                set({
                    loadingRegister: true,
                    errorRegister: null,
                })

                try {
                    const response = await myApi.post("/auth/register", req);

                    if(response.status == 201){
                        set({
                            loadingRegister: false,
                            errorRegister: null
                        })
                    }

                    return response.data.data;
                } catch (error) {
                    handleApiError(
                        error,
                        () => set({
                            loadingRegister: false,
                        }),
                        (messsage: any) => set({
                            errorRegister: messsage
                        })
                    )
                }
            },

            getMe: async () => {
                set({
                    loadingGetMe: true,
                    errorGetMe: null
                })
                try {
                    const response = await myApi.get("/auth/me", {
                        headers: {
                            Authorization: `Bearer ${get().accessToken}`,
                        },
                    });
                    if(response.status == 200){
                        set({
                            loadingGetMe: false,
                            errorGetMe: null,
                            user: response.data.data
                        })
                    }

                    return response.data.data;
                } catch (error) {
                    handleApiError(
                        error,
                        () => set({
                            loadingGetMe: false,
                        }),
                        (messsage: any) => set({
                            errorGetMe: messsage
                        })
                    )
                }
            },

            refreshAccessToken: async () => {
                set({
                    loadingRefreshAccessToken: true,
                    errorRefreshAccessToken: null
                })
                try {
                    const response = await myApi.post("/auth/refresh", {
                        refreshToken: get().refreshToken,
                    });

                    if(response.status == 200){
                        set({
                            loadingRefreshAccessToken: false,
                            errorRefreshAccessToken: null,
                            refreshToken: response.data.data.refreshToken,
                            accessToken: response.data.data.accessToken
                        })
                    }

                    return response.data.data;
                } catch (error) {
                    handleApiError(
                        error,
                        () => set({
                            loadingRefreshAccessToken: false,
                        }),
                        (messsage: any) => set({
                            errorRefreshAccessToken: messsage
                        })
                    )
                }
            },
        }),
        {
            name: "auth",
        }
    )
);

export const getAccessToken = () => {
    const { accessToken } = useAuthStore.getState();

    if (!accessToken) {
        throw new Error("Access token not found");
    }

    return accessToken;
};

export default useAuthStore;
