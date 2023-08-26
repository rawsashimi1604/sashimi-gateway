import React, { ReactNode } from 'react';

import Footer from './Footer';
import Navbar from './Navbar';
import Sidebar from './Sidebar';

interface ContainerProps {
  children: React.ReactElement | React.ReactElement[] | ReactNode;
}

function Container({ children }: ContainerProps) {
  return (
    <div className="flex h-screen w-screen ">
      <div className="min-w-[300px]">
        <Sidebar />
      </div>
      <main className="flex-grow p-6">
        <Navbar />
        {children}
        {/* <Footer /> */}
      </main>
    </div>
  );
}

export default Container;
