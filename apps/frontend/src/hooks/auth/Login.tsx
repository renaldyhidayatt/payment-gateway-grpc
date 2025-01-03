import useAuthStore from "@/store/auth";
import { useToast } from "../use-toast";
import { useNavigate } from "react-router-dom";
import { useState } from "react";

export default function useLogin() {
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    const { login, setLoadingLogin, loadingLogin, setErrorLogin } = useAuthStore();
    const navigate = useNavigate();
    const { toast } = useToast();

    const onFinish = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault(); 
        setLoadingLogin(true);

        if (!email || !password) {
            toast({
                title: "Error",
                description: "Email dan Password harus diisi",
                variant: "destructive",
            });
            setLoadingLogin(false);
            return;
        }

        try {
            const data = { email, password };
            
            await login(data);

            toast({
                title: "Success",
                description: "Login Berhasil",
                variant: "default",
            });

            navigate("/dashboard");
        } catch (error: any) {
            setErrorLogin(error?.message || "Terjadi kesalahan saat login");
            toast({
                title: "Error",
                description: error?.message || "Terjadi kesalahan saat login",
                variant: "destructive",
            });
        } finally {
            setLoadingLogin(false);
        }
    };

    return {
        email,
        setEmail,
        password,
        setPassword,
        onFinish,
        loadingLogin,
    };
}
