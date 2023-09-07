import React from 'react';

import { GetServiceByIdResponse } from '../../api/services/admin/responses/GetServiceById';
import ApiMethodTag from '../../components/tags/ApiMethodTag';
import Header from '../../components/typography/Header';
import { ApiMethod } from '../../types/api/ApiMethod.interface';

interface ServiceProps {
  data: GetServiceByIdResponse;
}

function ServiceInformation({ data }: ServiceProps) {
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
              </tbody>
            </table>
          </div>
        </div>
        <div className="grow">
          <Header text="routes" align="left" size="sm" />
          <div className="relative overflow-x-auto font-sans mt-3">
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
                      <td className="px-3 py-2">{new Date(route.createdAt).toLocaleString()}</td>
                      <td className="px-3 py-2">{new Date(route.updatedAt).toLocaleString()}</td>
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
