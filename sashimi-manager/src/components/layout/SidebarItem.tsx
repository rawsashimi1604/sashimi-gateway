import React from 'react';
import { Link } from 'react-router-dom';

interface SidebarItemProps {
  item: string;
  isSelected?: boolean;
  href?: string;
}

function SidebarItem({
  item,
  isSelected = false,
  href = '/'
}: SidebarItemProps) {
  return (
    <li
      className={`${
        isSelected && 'bg-blue-100 py-0.5 pl-2'
      } hover:bg-slate-100 transition-all hover:pl-4  duration-150  hover:cursor-pointer rounded-lg tracking-tighter`}
    >
      <Link to={href}>
        <span className="hover:transition-all duration-150">{item}</span>
      </Link>
    </li>
  );
}

export default SidebarItem;
