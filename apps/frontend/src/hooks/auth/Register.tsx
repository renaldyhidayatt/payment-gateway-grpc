import useAuthStore from "@/store/auth";
import { useToast } from "../use-toast";
import { useNavigate } from "react-router-dom";
import { useState } from "react";

export default function useRegister(){
    const [firstName, setFirstName] = useState<string>("");
    const [lastName, setLastName] = useState<string>("");
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [confirmPassword, setConfirmPassword] = useState<string>("");
    const navigate = useNavigate();

    const {
        register,
        setLoadingRegister,
        loadingRegister,
    } = useAuthStore();
    const {
        toast
    } = useToast();

    const onFinish = async(event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        setLoadingRegister(true);

        if (!firstName || !lastName || !email || !password || !confirmPassword) {
            toast({
                title: "Error",
                description: "Semua field harus diisi",
                variant: "destructive",
            });
            setLoadingRegister(false);
            return;
        }

        try{
            const data = {
                firstname: firstName,
                lastname: lastName,
                email: email,
                password: password,
                confirm_password: confirmPassword
            }

            await register(data);

            toast({
                title: "Success",
                description: "Register Berhasil",
                variant: "default",
            });
            navigate("/auth/login");
        } catch (error: any) {
            toast({
                title: "Error",
                description: error?.message || "Terjadi kesalahan saat register",
                variant: "destructive",
            });
            setLoadingRegister(false);

        }
    }

    return {
        firstName,
        setFirstName,
        lastName,
        setLastName,
        email,
        setEmail,
        password,
        setPassword,
        confirmPassword,
        setConfirmPassword,
        onFinish,
        loadingRegister
    }

}