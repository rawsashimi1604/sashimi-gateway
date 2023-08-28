import React from 'react';
import { AiOutlineLoading3Quarters } from 'react-icons/ai';

function Notifcations() {
  return (
    <div className="p-4 pt-5 flex flex-col relative z-0 w-full h-full">
      {/* border */}
      <div className="absolute left-0 border-l border-gray-200 h-full w-full -z-10"></div>
      <div className="flex flex-row items-center justify-between">
        <h1 className="font-cabin font-light tracking-tight">notifications</h1>

        <div className="font-cabin flex flex-row border px-2 rounded-lg shadow-sm bg-gray-50 py-0.5 text-sm gap-2 items-center justify-between text-gray-600">
          <span>listening</span>
          <AiOutlineLoading3Quarters className="animate-spin" />
        </div>
      </div>
    </div>
  );
}

export default Notifcations;
