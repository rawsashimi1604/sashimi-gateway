import React from 'react';

interface CardProps {
  header: string;
  data: string;
}

function Card({ header, data }: CardProps) {
  return (
    <div className="w-full border-gray-200 shadow-md border rounded-lg text-sm p-2 h-24 flex items-center">
      <div>
        <h2 className="text-lg tracking-tighter">{header}</h2>
        <h3 className="font-light tracking-wider text-3xl">{data}</h3>
      </div>

      {/* Insert some logo here */}
    </div>
  );
}

export default Card;
