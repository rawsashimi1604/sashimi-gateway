import React from 'react';

interface CardProps {
  header: string;
  data: string;
}

function Card({ header, data }: CardProps) {
  return (
    <div className="w-full shadow-md rounded-lg text-sm p-2 px-4 h-24 flex flex-col justify-center bg-sashimi-blue">
      <div>
        <h2 className="text-sm tracking-wide">{header}</h2>
        <h3 className="font-light tracking-wider text-4xl">{data}</h3>
      </div>

      {/* Insert some icon here */}
    </div>
  );
}

export default Card;
