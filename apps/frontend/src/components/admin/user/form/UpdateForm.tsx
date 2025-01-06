import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { createUserSchema, UpdateUserFormValues } from "@/schemas";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { forwardRef } from "react";

interface UpdateUserFormProps {
  onSubmit: (data: UpdateUserFormValues) => void;
}

const UpdateUserForm = forwardRef<HTMLFormElement, UpdateUserFormProps>(
  ({ onSubmit }, ref) => {
    const {
      register,
      handleSubmit,
      formState: { errors },
    } = useForm<UpdateUserFormValues>({
      resolver: zodResolver(createUserSchema),
    });

    return (
      <form ref={ref} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
        <div>
          <Label htmlFor="firstname" className="text-sm font-medium">
            First Name
          </Label>
          <Input
            id="firstname"
            type="text"
            placeholder="Enter first name"
            className="mt-1"
            {...register("firstname")}
          />
          {errors.firstname && (
            <p className="text-sm text-red-500">{errors.firstname.message}</p>
          )}
        </div>
        <div>
          <Label htmlFor="lastname" className="text-sm font-medium">
            Last Name
          </Label>
          <Input
            id="lastname"
            type="text"
            placeholder="Enter last name"
            className="mt-1"
            {...register("lastname")}
          />
          {errors.lastname && (
            <p className="text-sm text-red-500">{errors.lastname.message}</p>
          )}
        </div>
        <div>
          <Label htmlFor="email" className="text-sm font-medium">
            Email
          </Label>
          <Input
            id="email"
            type="email"
            placeholder="Enter email"
            className="mt-1"
            {...register("email")}
          />
          {errors.email && (
            <p className="text-sm text-red-500">{errors.email.message}</p>
          )}
        </div>
        <div>
          <Label htmlFor="password" className="text-sm font-medium">
            Password
          </Label>
          <Input
            id="password"
            type="password"
            placeholder="Enter password"
            className="mt-1"
            {...register("password")}
          />
          {errors.password && (
            <p className="text-sm text-red-500">{errors.password.message}</p>
          )}
        </div>
        <div>
          <Label htmlFor="confirm_password" className="text-sm font-medium">
            Confirm Password
          </Label>
          <Input
            id="confirm_password"
            type="password"
            placeholder="Confirm password"
            className="mt-1"
            {...register("confirm_password")}
          />
          {errors.confirm_password && (
            <p className="text-sm text-red-500">
              {errors.confirm_password.message}
            </p>
          )}
        </div>
      </form>
    );
  },
);

UpdateUserForm.displayName = "CreateUserForm";

export default UpdateUserForm;
