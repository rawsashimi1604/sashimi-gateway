import React from 'react';

import { Consumer } from '../../api/services/admin/models/Consumer';
import { parseDateString } from '../../utils/parseDate';

interface TableProps {
  consumers: Consumer[];
}

function Table({ consumers }: TableProps) {
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
                Username
              </th>
              <th scope="col" className="px-3 py-4">
                Services
              </th>
              <th scope="col" className="px-3 py-4">
                Created At
              </th>
              <th scope="col" className="px-3 py-4">
                Updated At
              </th>
            </tr>
          </thead>
          <tbody>
            {consumers &&
              consumers.map((consumer) => {
                return (
                  <tr
                    key={consumer.id}
                    className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 cursor-pointer text-xs"
                    onClick={() => console.log('to be implemented, navigate')}
                  >
                    <td className="px-3 py-4">{consumer.id}</td>
                    <td className="px-3 py-4">{consumer.username}</td>
                    <td className="px-3 py-4">{consumer.services.map((service) => service.name).join(', ')}</td>
                    <td className="px-3 py-4">{parseDateString(consumer.createdAt)}</td>
                    <td className="px-3 py-4">{parseDateString(consumer.updatedAt)}</td>
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
