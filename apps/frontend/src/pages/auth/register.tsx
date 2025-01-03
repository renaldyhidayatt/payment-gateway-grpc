import RegisterForm from "@/components/auth/register/RegisterForm";
import LinkRegister from "@/components/auth/register/LinkRegister";
import useRegister from "@/hooks/auth/Register";

export default function RegisterPage() {
  const {
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
    onFinish
  } = useRegister();


  return (
    <div className="flex w-full lg:w-1/2 items-center justify-center p-8 bg-background">
      <div className="mx-auto w-full max-w-sm space-y-6">
        <div className="flex flex-col space-y-2 text-center">
          <h1 className="text-2xl font-semibold tracking-tight">
            Welcome back
          </h1>
          <p className="text-sm text-muted-foreground">
            Enter your email to sign in to your account
          </p>
        </div>
        <RegisterForm
          handleSubmit={onFinish}
          firstName={firstName}
          setFirstName={setFirstName}
          lastName={lastName}
          setLastName={setLastName}
          email={email}
          setEmail={setEmail}
          password={password}
          setPassword={setPassword}
          confirmPassword={confirmPassword}
          setConfirmPassword={setConfirmPassword}
        />
        <LinkRegister />
      </div>
    </div>

  )

}
