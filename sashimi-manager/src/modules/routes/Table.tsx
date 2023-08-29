import React from 'react';

import ApiMethodTag from '../../components/tags/ApiMethodTag';

function Table() {
  return (
    <div>
      <div className="relative overflow-x-auto font-sans">
        <table className="w-full text-sm text-left text-gray-500">
          <thead className="text-xs text-black lowercase bg-sashimi-gray tracking-tighter">
            <tr>
              <th scope="col" className="px-2 py-1.5">
                Id
              </th>
              <th scope="col" className="px-2 py-1.5">
                Service
              </th>
              <th scope="col" className="px-2 py-1.5">
                Method
              </th>
              <th scope="col" className="px-2 py-1.5">
                Path
              </th>
              <th scope="col" className="px-2 py-1.5">
                Description
              </th>
            </tr>
          </thead>
          <tbody>
            <tr className="bg-white border-b hover:bg-sashimi-gray/50 cursor-pointer">
              <td className="px-2 py-1.5">1</td>
              <td className="px-2 py-1.5">Salmon</td>
              <td className="px-2 py-1.5">
                <ApiMethodTag method="GET" />
              </td>
              <td className="px-2 py-1.5 italic">/</td>
              <td className="px-2 py-1.5">Get all salmon dishes</td>
            </tr>
            <tr className="bg-white border-b hover:bg-sashimi-gray/50 cursor-pointer">
              <td className="px-2 py-1.5">2</td>
              <td className="px-2 py-1.5">Salmon</td>
              <td className="px-2 py-1.5">
                <ApiMethodTag method="GET" />
              </td>
              <td className="px-2 py-1.5 italic">/:id</td>
              <td className="px-2 py-1.5">Get all salmon dishes</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Table;
