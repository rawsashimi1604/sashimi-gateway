import React from 'react';
import { Link } from 'react-router-dom';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import Table from './Table';

function Services() {
  return (
    <Container>
      <Header text="gateway services" align="left" size="sm" />

      <div className="flex flex-row items-center justify-between mb-3">
        <h3 className="text-xs -mt-2 text-gray-600 ">gateway services</h3>
        <Link to="/services/register">
          <button className="flex-end text-xs py-1 px-2 bg-blue-500 text-white shadow-md rounded-lg font-sans border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg">
            <span>add service</span>
          </button>
        </Link>
      </div>

      <Table />
    </Container>
  );
}

export default Services;
