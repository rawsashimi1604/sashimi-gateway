import React from 'react';
import { useNavigate } from 'react-router-dom';

import { GetServiceByIdResponse } from '../../api/services/admin/responses/GetServiceById';
import ApiMethodTag from '../../components/tags/ApiMethodTag';
import Header from '../../components/typography/Header';
import { ApiMethod } from '../../types/api/ApiMethod.interface';
import { parseDateString } from '../../utils/parseDate';

interface ServiceProps {
  data: GetServiceByIdResponse;
}

function ServiceInformation({ data }: ServiceProps) {
  const navigate = useNavigate();

  return (
    <section className="text-sans">
      <div className="flex flex-row justify-between gap-6">
        <div className="grow">
          <Header text="details" align="left" size="sm" />
          <div className="relative overflow-x-auto font-sans mt-3">
            <table className="w-full text-sm text-left text-gray-500">
              <thead className="text-xs text-sashimi-deepgray lowercase bg-sashimi-gray tracking-tighter">
                <tr>
                  <th scope="col" className="px-3 py-2">
                    Parameter
                  </th>
                  <th scope="col" className="px-3 py-2">
                    Value
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 text-xs">
                  <td className="px-3 py-2">id</td>
                  <td className="px-3 py-2">{data.service.id}</td>
                </tr>
                <tr className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 text-xs">
                  <td className="px-3 py-2">name</td>
                  <td className="px-3 py-2">{data.service.name}</td>
                </tr>
                <tr className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 text-xs">
                  <td className="px-3 py-2">path</td>
                  <td className="px-3 py-2">{data.service.path}</td>
                </tr>
                <tr className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 text-xs">
                  <td className="px-3 py-2">target url</td>
                  <td className="px-3 py-2">{data.service.targetUrl}</td>
                </tr>
                <tr className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 text-xs">
                  <td className="px-3 py-2">description</td>
                  <td className="px-3 py-2">{data.service.description}</td>
                </tr>
                <tr className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 text-xs">
                  <td className="px-3 py-2">created at</td>
                  <td className="px-3 py-2">{parseDateString(data.service.createdAt)}</td>
                </tr>
                <tr className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 text-xs">
                  <td className="px-3 py-2">updated at</td>
                  <td className="px-3 py-2">{parseDateString(data.service.updatedAt)}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div className="grow">
          <div className="flex flex-row items-center justify-between">
            <Header text="routes" align="left" size="sm" />
            <button
              type="button"
              onClick={() => navigate('/routes/register')}
              className="text-xs py-1 px-2 bg-blue-500 text-white shadow-md rounded-lg font-sans border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg"
            >
              <span>add route</span>
            </button>
          </div>
          <div className="relative overflow-x-auto font-sans mt-1">
            <table className="w-full text-sm text-left text-gray-500">
              <thead className="text-xs text-sashimi-deepgray lowercase bg-sashimi-gray tracking-tighter">
                <tr>
                  <th scope="col" className="px-3 py-2">
                    id
                  </th>
                  <th scope="col" className="px-3 py-2">
                    method
                  </th>
                  <th scope="col" className="px-3 py-2">
                    path
                  </th>
                  <th scope="col" className="px-3 py-2">
                    description
                  </th>
                  <th scope="col" className="px-3 py-2">
                    created at
                  </th>
                  <th scope="col" className="px-3 py-2">
                    updated at
                  </th>
                </tr>
              </thead>
              <tbody>
                {data.service.routes.map((route) => {
                  return (
                    <tr
                      key={route.id}
                      className="transition-all duration-150 hover:pl-10 bg-white border-b hover:bg-sashimi-gray/50 text-xs"
                    >
                      <td className="px-3 py-2">{route.id}</td>
                      <td className="px-3 py-2">
                        <ApiMethodTag method={route.method as ApiMethod} />
                      </td>
                      <td className="px-3 py-2">{route.path}</td>
                      <td className="px-3 py-2">{route.description}</td>
                      <td className="px-3 py-2">{parseDateString(route.createdAt)}</td>
                      <td className="px-3 py-2">{parseDateString(route.updatedAt)}</td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </section>
  );
}

export default ServiceInformation;
