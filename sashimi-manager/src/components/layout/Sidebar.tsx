import React from 'react';

import GatewayLogo from '../sashimi-gateway/GatewayLogo';
import Logo from '../sashimi-gateway/Logo';
import SidebarItem from './SidebarItem';

function Sidebar() {
  return (
    <nav className="relative w-full h-full p-6 pt-8 pr-8 flex flex-col justify-between z-0">
      {/* border */}
      <div className="absolute right-0 border-r border-gray-200 h-full w-full -z-10"></div>
      <div>
        <div className="mb-10">
          <Logo />
        </div>

        <div className="border-b border-gray-200 pb-5">
          {/* Reverse proxy name */}
          <GatewayLogo gateway="Sushi Gateway" user="admin" />
        </div>
        {/* Services, Routes */}
        <div className="mt-4">
          <h2 className="font-bold text-gray-800 tracking-tighter">gateway</h2>
          <ul className="flex flex-col gap-0">
            <SidebarItem item="Home" isSelected={true} />
            <SidebarItem item="Services" />
            <SidebarItem item="Routes" />
            <SidebarItem item="Logs" />
          </ul>
        </div>
        {/* Settings , Logout */}
        <div className="mt-4">
          <h2 className="font-bold text-gray-800 tracking-tighter">account</h2>
          <ul className="flex flex-col gap-0">
            <SidebarItem item="Account" />
            <SidebarItem item="Settings" />
          </ul>
        </div>
      </div>

      <button className="w-full flex-end py-2 bg-blue-500 text-white shadow-md rounded-full font-sans border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg">
        logout
      </button>
    </nav>
  );
}

export default Sidebar;
