import React from 'react';
import { BsDot } from 'react-icons/bs';
import { useNavigate } from 'react-router-dom';

import { Service } from '../../api/services/admin/models/Service';
import { ServiceHealth } from '../../types/api/ServiceHealth.interface';
import { serviceHealthColor } from '../../utils/serviceHealthColor';

interface TableProps {
  services: Service[];
}

function Table({ services }: TableProps) {
  const navigate = useNavigate();

  function handleNavigate(id: any) {
    navigate('/services/' + id);
  }

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
              <th scope="col" className="px-3 py-4">
                Health
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
                    onClick={() => handleNavigate(service.id)}
                  >
                    <td className="px-3 py-4">{service.id}</td>
                    <td className="px-3 py-4">{service.name}</td>
                    <td className="px-3 py-4 italic">{service.path}</td>
                    <td className="px-3 py-4">{service.targetUrl}</td>
                    <td className="px-3 py-4">{service.description}</td>
                    <td className="pr-3 py-4 tracking-wider">
                      <div className="flex flex-row items-center justify-start gap-1">
                        <BsDot
                          className={`${
                            serviceHealthColor(service.health as ServiceHealth).text
                          } w-8 h-8 animate-pulse`}
                        />
                        <span className={`${serviceHealthColor(service.health as ServiceHealth).text}`}>
                          {service.health}
                        </span>
                      </div>
                    </td>
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
