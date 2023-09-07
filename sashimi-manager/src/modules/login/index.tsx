import Logo from '../../components/sashimi-gateway/Logo';
import LoginForm from './LoginForm';

function Login() {
  return (
    <div className="flex flex-col items-center justify-center w-screen h-screen">
      <div className="mb-12">
        <Logo />
      </div>
      <LoginForm />
    </div>
  );
}

export default Login;
