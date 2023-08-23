import React from 'react';

interface GatewayLogoProps {
  name: string;
}

function GatewayLogo({ name }: GatewayLogoProps) {
  return (
    <div className="flex flex-row items-center gap-4 text-black">
      <div className="shadow-lg flex items-center justify-center p-2 bg-blue-600 text-white w-12 h-12 text-lg rounded-lg">
        {name.charAt(0)}
      </div>
      <h1 className="text-lg tracking-wide">{name}</h1>
    </div>
  );
}

export default GatewayLogo;
