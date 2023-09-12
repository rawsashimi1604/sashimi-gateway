import React, { ReactNode } from 'react';

import { useRedirectLogin } from '../../hooks/useRedirectLogin';
import Footer from './Footer';
import Navbar from './Navbar';
import Notifications from './Notifications';
import Sidebar from './Sidebar';

interface ContainerProps {
  children: React.ReactElement | React.ReactElement[] | ReactNode;
}

function Container({ children }: ContainerProps) {
  // Redirect all traffic to login page when jwt token is absent.
  useRedirectLogin();

  return (
    <div className="flex h-screen w-screen ">
      <div className="min-w-[300px]">
        <Sidebar />
      </div>
      <main className="flex-grow p-6">
        <Navbar />
        <div className="mt-3 container overflow-y-scroll">{children}</div>
        {/* <Footer /> */}
      </main>
      <div className="min-w-[300px]">
        <Notifications />
      </div>
    </div>
  );
}

export default Container;
