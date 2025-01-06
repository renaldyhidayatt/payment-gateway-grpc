import { CreateUserFormValues } from "@/schemas";
import { SubmitHandler } from "react-hook-form";

export interface UserCreateFormProps {
  handleSubmit: SubmitHandler<CreateUserFormValues>;
  firstName: string;
  setFirstName: (firstName: string) => void;
  lastName: string;
  setLastName: (lastName: string) => void;
  email: string;
  setEmail: (email: string) => void;
  password: string;
  setPassword: (password: string) => void;
  confirmPassword: string;
  setConfirmPassword: (confirmPassword: string) => void;
}
