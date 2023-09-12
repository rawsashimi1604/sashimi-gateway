import React, { useState } from 'react';
import * as yup from 'yup';

import AdminAuth from '../../api/services/admin/AdminAuth';
import TextInput from '../../components/input/TextInput';

// Define validation schema using yup for login
const loginValidationSchema = yup.object().shape({
  username: yup.string().required('Username is required.'),
  password: yup.string().required('Password is required.')
});

function LoginForm() {
  const [loginData, setLoginData] = useState({
    username: '',
    password: ''
  });
  const [validationErrors, setValidationErrors] = useState<{ [key: string]: string }>({});

  const handleChange = (name: string, value: string) => {
    setLoginData((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      await loginValidationSchema.validate(loginData, { abortEarly: false });
      console.log('Login data is valid:', loginData);
      setValidationErrors({});

      // Authenticate the user, e.g. send a request to the server
      const loginRes = await AdminAuth.login(loginData.username, loginData.password);
      console.log({ loginRes });
    } catch (err) {
      if (err instanceof yup.ValidationError) {
        const errorObj: { [key: string]: string } = {};
        for (let error of err.inner) {
          errorObj[error.path as string] = error.message;
        }
        setValidationErrors(errorObj);
      }
    }
  };

  return (
    <div className="font-sans w-[300px]">
      <form className="flex flex-col gap-3" onSubmit={handleLogin}>
        {/* Username */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="username" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1 ">username</span>
          </label>
          <div className="">
            <TextInput
              id="username"
              name="username"
              value={loginData.username}
              onChange={(e) => handleChange('username', e.target.value)}
              error={validationErrors.username}
            />
          </div>
        </div>

        {/* Password */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="password" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1 ">password</span>
          </label>
          <div className="">
            <TextInput
              type="password"
              id="password"
              name="password"
              value={loginData.password}
              onChange={(e) => handleChange('password', e.target.value)}
              error={validationErrors.password}
            />
          </div>
        </div>

        <button
          type="submit"
          className="w-[80px] mt-2 text-xs py-1.5 px-2 pb-2 text-white bg-sashimi-deepblue shadow-md rounded-lg font-sans tracking-wider border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg"
        >
          <span>login</span>
        </button>
      </form>
    </div>
  );
}

export default LoginForm;
