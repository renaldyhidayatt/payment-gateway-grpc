import useUserStore from "@/store/user/user";
import { useNavigate } from "react-router-dom";
import { useToast } from "../use-toast";
import { SubmitHandler } from "react-hook-form";
import { CreateUserFormValues, createUserSchema } from "@/schemas";
import { z } from "zod";

export default function useCreateUser() {
  const {
    createUser,
    setLoadingCreateUser,
    loadingCreateUser,
    setErrorCreateUser,
  } = useUserStore();

  const navigate = useNavigate();
  const { toast } = useToast();

  const onFinish: SubmitHandler<CreateUserFormValues> = async (data) => {
    setLoadingCreateUser(true);

    try {
      const validatedValues = createUserSchema.parse(data);

      await new Promise((resolve) => setTimeout(resolve, 1200));

      const result = await createUser(validatedValues);

      if (result) {
        toast({
          title: "Success",
          description: "User berhasil dibuat",
          variant: "default",
        });

        navigate("/users");
      } else {
        toast({
          title: "Error",
          description: "Gagal membuat user. Silakan coba lagi.",
          variant: "destructive",
        });
      }
    } catch (error: any) {
      if (error instanceof z.ZodError) {
        const errorMessage = error.errors.map((err) => err.message).join(", ");
        setErrorCreateUser(errorMessage);
        toast({
          title: "Validation Error",
          description: errorMessage,
          variant: "destructive",
        });
      } else {
        setErrorCreateUser(
          error?.message || "Terjadi kesalahan saat membuat user",
        );
        toast({
          title: "Error",
          description: error?.message || "Terjadi kesalahan saat membuat user",
          variant: "destructive",
        });
      }
    } finally {
      setLoadingCreateUser(false);
    }
  };

  return {
    onFinish,
    loadingCreateUser,
  };
}
