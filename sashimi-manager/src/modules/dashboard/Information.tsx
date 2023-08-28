import c3 from 'c3';
import React from 'react';

function Information() {
  var chart = c3.generate({
    bindto: '#chart',
    data: {
      columns: [
        ['data1', 30, 200, 100, 400, 150, 250],
        ['data2', 50, 20, 10, 40, 15, 25]
      ]
    }
  });

  return <div>Information</div>;
}

export default Information;
