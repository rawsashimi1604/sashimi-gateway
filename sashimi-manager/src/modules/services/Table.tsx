import React from 'react';

import { Service } from '../../api/services/admin/responses/Service';

interface TableProps {
  services: Service[];
}

function Table({ services }: TableProps) {
  console.log(services);

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
                Name
              </th>
              <th scope="col" className="px-3 py-4">
                Path
              </th>
              <th scope="col" className="px-3 py-4">
                Target URL
              </th>
              <th scope="col" className="px-3 py-4">
                Description
              </th>
            </tr>
          </thead>
          <tbody>
            {services &&
              services.map((service) => {
                return (
                  <tr
                    key={service.id}
                    className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 cursor-pointer text-xs"
                  >
                    <td className="px-3 py-4">{service.id}</td>
                    <td className="px-3 py-4">{service.name}</td>
                    <td className="px-3 py-4 italic">{service.path}</td>
                    <td className="px-3 py-4">{service.targetUrl}</td>
                    <td className="px-3 py-4">{service.description}</td>
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
