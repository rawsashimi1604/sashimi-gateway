import React from 'react';

import GatewayLogo from '../sashimi-gateway/GatewayLogo';
import SidebarItem from './SidebarItem';

function Sidebar() {
  return (
    <nav className="relative w-full h-full p-6 pt-8 pr-8 flex flex-col justify-between z-0">
      {/* border */}
      <div className="absolute right-0 border-r border-gray-200 h-full w-full -z-10"></div>
      <div>
        {/* Reverse proxy name */}
        <GatewayLogo name="Sushi Gateway" />
        {/* Services, Routes */}
        <div className="mt-4">
          <h2 className="font-bold text-gray-800 tracking-tighter">gateway</h2>
          <ul className="flex flex-col gap-0">
            <SidebarItem item="Home" isSelected={true} />
            <SidebarItem item="Services" />
            <SidebarItem item="Routes" />
          </ul>
        </div>
        {/* Settings , Logout */}
        <div className="mt-4">
          <h2 className="font-bold text-gray-800 tracking-tighter">account</h2>
          <ul className="flex flex-col gap-0">
            <SidebarItem item="Configure Account" />
          </ul>
        </div>
      </div>

      <button className="w-full flex-end py-2 bg-blue-500 text-white shadow-md rounded-full font-sans border-0">
        logout
      </button>
    </nav>
  );
}

export default Sidebar;