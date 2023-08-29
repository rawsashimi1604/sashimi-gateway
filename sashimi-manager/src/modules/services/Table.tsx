import React from 'react';

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
                Name
              </th>
              <th scope="col" className="px-2 py-1.5">
                Path
              </th>
              <th scope="col" className="px-2 py-1.5">
                Target URL
              </th>
              <th scope="col" className="px-2 py-1.5">
                Description
              </th>
            </tr>
          </thead>
          <tbody>
            <tr className="bg-white border-b">
              <td className="px-2 py-1.5">1</td>
              <td className="px-2 py-1.5">Salmon</td>
              <td className="px-2 py-1.5 italic">/salmon</td>
              <td className="px-2 py-1.5">http://localhost:8081</td>
              <td className="px-2 py-1.5">
                The salmon microservice used to learn how to create a golang api
                gateway infrastructure.
              </td>
            </tr>
            <tr className="bg-white border-b">
              <td className="px-2 py-1.5">2</td>
              <td className="px-2 py-1.5">Tuna</td>
              <td className="px-2 py-1.5 italic">/tuna</td>
              <td className="px-2 py-1.5">http://localhost:8082</td>
              <td className="px-2 py-1.5">
                The tuna microservice used to learn how to create a golang api
                gateway infrastructure.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Table;
