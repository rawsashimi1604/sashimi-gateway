import React from 'react';

interface TextInputProps {
  id: string;
  name: string;
  placeholder?: string;
  value?: string;
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
  error?: string;
}

function TextInput({
  id,
  name,
  placeholder,
  value,
  onChange,
  error
}: TextInputProps) {
  return (
    <>
      <input
        id={id}
        name={name}
        type="text"
        className="font-sans tracking-tighter indent-1 rounded-lg w-full bg-sashimi-gray border border-sashimi-deepgray px-1 py-0.5 text-sm focus:outline-none"
        placeholder={placeholder}
        value={value}
        onChange={onChange}
      />
      {error && <span className="text-red-500 text-xs">{error}</span>}
    </>
  );
}

export default TextInput;
