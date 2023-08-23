import React from 'react';

interface SidebarItemProps {
  item: string;
  isSelected?: boolean;
}

function SidebarItem({ item, isSelected = false }: SidebarItemProps) {
  return (
    <li
      className={`${
        isSelected && 'bg-blue-100 py-0.5 pl-2'
      } hover:bg-slate-100 transition-all hover:pl-4  duration-150  hover:cursor-pointer rounded-lg tracking-tighter`}
    >
      <span className="hover:transition-all duration-150">{item}</span>
    </li>
  );
}

export default SidebarItem;
