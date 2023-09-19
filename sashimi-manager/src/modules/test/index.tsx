import React from 'react';

function Test() {
  const items = [
    {
      tag: 'Information',
      item: 'random1234'
    },
    {
      tag: 'Informatio1',
      item: 'random1234'
    },
    {
      tag: 'Information2',
      item: 'random1234'
    },
    {
      tag: 'Information3',
      item: 'random1235'
    },
    {
      tag: 'Information4',
      item: 'random1236'
    },
    {
      tag: 'Information5',
      item: 'random1237'
    }
  ];

  return (
    <div
      style={{
        background: 'yellow',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center'
      }}
    >
      <h1 className="text-2xl font-bold" style={{ paddingBottom: '20px' }}>
        Loan application
      </h1>

      <div style={{ width: '25%', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        {items.map((item) => {
          return (
            <div style={{ marginBottom: '10px' }}>
              <span style={{ display: 'inline-block', width: '150px' }}>{item.tag}</span>
              <span>{item.item}</span>
            </div>
          );
        })}
      </div>
    </div>
  );
}

export default Test;
