import React from 'react';

import { ReactComponent as LogoSvg } from '../../assets/sushi_logo.svg';

function Logo() {
  return (
    <div className="flex flex-col justify-center items-center font-light text-2xl font-sans">
      <LogoSvg className="transform scale-110" />
      <div className="relative">
        <h1 className="-mt-2 font-cabin">SASHIMI GATEWAY</h1>
        <div className="font-cabin tracking-wider absolute text-sm border border-blue-500 shadow-lg px-2 py-0.5 -right-5 top-5.5 font-bold">
          manager
        </div>
      </div>
    </div>
  );
}

export default Logo;
