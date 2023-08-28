import React from 'react';

interface SelectInputProps {
  options: string[];
  onChange?: (value: string) => void;
}

function SelectInput({ options, onChange }: SelectInputProps) {
  return (
    <select
      className="rounded-lg border border-sashimi-deepgray px-1 py-0.5 text-xs focus:outline-none"
      onChange={(e) => onChange && onChange(e.target.value)}
    >
      {options?.map((option) => {
        return (
          <option value={option} key={option}>
            {option}
          </option>
        );
      })}
    </select>
  );
}

export default SelectInput;
