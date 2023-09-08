import React from 'react';

import { Route } from '../../api/services/admin/models/Route';
import ApiMethodTag from '../../components/tags/ApiMethodTag';
import { ApiMethod } from '../../types/api/ApiMethod.interface';

interface TableProps {
  routes: Route[];
}

function Table({ routes }: TableProps) {
  return (
    <div>
      <div className="relative overflow-x-auto font-sans">
        <table className="w-full text-sm text-left text-gray-500">
          <thead className="text-xs text-sashimi-deepgray lowercase bg-sashimi-gray tracking-tighter">
            <tr>
              <th scope="col" className="px-3 py-4">
                Id
              </th>
              <th scope="col" className="px-3 py-4">
                Service
              </th>
              <th scope="col" className="px-3 py-4">
                Method
              </th>
              <th scope="col" className="px-3 py-4">
                Path
              </th>
              <th scope="col" className="px-3 py-4">
                Description
              </th>
            </tr>
          </thead>
          <tbody>
            {routes &&
              routes.map((route) => {
                return (
                  <tr className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 cursor-pointer text-xs">
                    <td className="px-3 py-3">{route.id}</td>
                    <td className="px-3 py-3">{route.serviceId}</td>
                    <td className="px-3 py-3">
                      <ApiMethodTag method={route.method as ApiMethod} />
                    </td>
                    <td className="px-3 py-3 italic">{route.path}</td>
                    <td className="px-3 py-3">{route.description}</td>
                  </tr>
                );
              })}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Table;
