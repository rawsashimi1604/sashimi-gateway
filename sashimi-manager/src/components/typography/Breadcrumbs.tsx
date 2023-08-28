import React from 'react';
import { useLocation } from 'react-router-dom';
import useBreadcrumbs from 'use-react-router-breadcrumbs';

function Breadcrumbs() {
  const location = useLocation();
  const breadcrumbs = useBreadcrumbs();
  console.log(breadcrumbs);

  return (
    <div className="flex flex-row items-center text-sm gap-1">
      <span className="text-gray-500">manager {'>'}</span>
      {breadcrumbs.map((breadcrumb, i) => {
        return (
          <span
            className={`${
              breadcrumb.key === location.pathname
                ? 'font-semibold'
                : 'text-gray-500'
            }`}
          >
            {breadcrumb.key === '/' ? 'home' : breadcrumb.key.slice(1)}{' '}
            {i + 2 == breadcrumbs.length && '>'}
          </span>
        );
      })}
    </div>
  );
}

export default Breadcrumbs;
