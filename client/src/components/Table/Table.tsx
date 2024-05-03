import React from 'react';

type Column = {
    header: string;
    accessor: string;
}

interface TableProps {
  data: any[];
  columns: Column[];
}

const Table: React.FC<TableProps> = ({ data, columns }) => {
  return (
    <div className="overflow-x-auto mx-4">
      <table className="table-auto w-full border-collapse border border-gray-800">
        <thead>
          <tr className="bg-black text-white">
            {columns.map((column: Column) => (
              <th key={column.accessor} className="px-4 py-2">
                {column.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="bg-white text-black">
          {data.map((row) => (
            <tr key={row.id}>
              {columns.map((column: Column) => (
                <td key={column.header} className="border border-gray-800 px-4 py-2">
                  {row[column.accessor]}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Table;