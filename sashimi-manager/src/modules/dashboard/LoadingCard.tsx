import React from 'react';

import LoadingSpinner from '../../components/utils/LoadingSpinner';

interface LoadingCardProps {
  header: string;
}

function LoadingCard({ header }: LoadingCardProps) {
  return (
    <div className="w-full shadow-md rounded-lg text-sm p-2 px-4 h-24 flex flex-col justify-center bg-sashimi-blue">
      <div>
        <h2 className="text-sm mb-4 tracking-wide">{header}</h2>
        <div className="flex flex-row items-center justiy-start gap-2 text-sashimi-deepgray">
          <LoadingSpinner size={12} color="#505050" />
          <h3 className="font-cabin tracking-wider text-[18px] font-bold">loading...</h3>
        </div>
      </div>
    </div>
  );
}

export default LoadingCard;
