import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

export default function CreateUserForm({
  formData = {
    firstname: '',
    lastname: '',
    email: '',
    password: '',
    confirm_password: '',
  },
  onFormChange,
  formErrors = {},
}: {
  formData: any;
  onFormChange: (field: string, value: any) => void;
  formErrors: Record<string, string>;
}) {
  return (
    <div className="space-y-4">
      {/* First Name */}
      <div>
        <Label htmlFor="firstname" className="text-sm font-medium">
          First Name
        </Label>
        <Input
          id="firstname"
          type="text"
          placeholder="Enter first name"
          className="mt-1"
          value={formData?.firstname || ''}
          onChange={(e) => onFormChange('firstname', e.target.value)}
        />
        {formErrors.firstname && (
          <p className="text-red-500 text-sm mt-1">{formErrors.firstname}</p>
        )}
      </div>

      {/* Last Name */}
      <div>
        <Label htmlFor="lastname" className="text-sm font-medium">
          Last Name
        </Label>
        <Input
          id="lastname"
          type="text"
          placeholder="Enter last name"
          className="mt-1"
          value={formData?.lastname || ''}
          onChange={(e) => onFormChange('lastname', e.target.value)}
        />
        {formErrors.lastname && (
          <p className="text-red-500 text-sm mt-1">{formErrors.lastname}</p>
        )}
      </div>

      {/* Email */}
      <div>
        <Label htmlFor="email" className="text-sm font-medium">
          Email
        </Label>
        <Input
          id="email"
          type="email"
          placeholder="Enter email"
          className="mt-1"
          value={formData?.email || ''}
          onChange={(e) => onFormChange('email', e.target.value)}
        />
        {formErrors.email && (
          <p className="text-red-500 text-sm mt-1">{formErrors.email}</p>
        )}
      </div>

      {/* Password */}
      <div>
        <Label htmlFor="password" className="text-sm font-medium">
          Password
        </Label>
        <Input
          id="password"
          type="password"
          placeholder="Enter password"
          className="mt-1"
          value={formData?.password || ''}
          onChange={(e) => onFormChange('password', e.target.value)}
        />
        {formErrors.password && (
          <p className="text-red-500 text-sm mt-1">{formErrors.password}</p>
        )}
      </div>

      {/* Confirm Password */}
      <div>
        <Label htmlFor="confirm_password" className="text-sm font-medium">
          Confirm Password
        </Label>
        <Input
          id="confirm_password"
          type="password"
          placeholder="Confirm password"
          className="mt-1"
          value={formData?.confirm_password || ''}
          onChange={(e) => onFormChange('confirm_password', e.target.value)}
        />
        {formErrors.confirm_password && (
          <p className="text-red-500 text-sm mt-1">
            {formErrors.confirm_password}
          </p>
        )}
      </div>
    </div>
  );
}
