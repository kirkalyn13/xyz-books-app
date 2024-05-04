import React from 'react'
import { FaEdit, FaTrash } from 'react-icons/fa';

type Column = {
    header: string;
    accessor: string;
}

interface TableProps {
  data: any[];
  columns: Column[];
}


const Table: React.FC<TableProps> = ({ data, columns }) => {
  const columnsWithActions = (columns: Column[]) => [{ header: 'Actions', accessor: 'actions' }, ...columns]

  const actionsRow = (
    <div className="flex justify-around align-center">
      <FaEdit className='text-xl text-black'/>
      <FaTrash className='text-xl text-black' />
    </div>
  )

  const parseRowData = (rowData: any): any => {
    if (typeof rowData !== 'object') return rowData
    else if (rowData["name"]) return rowData["name"]
    else if (Array.isArray(rowData)) return rowData.map((data: any) => (data["first_name"] + " " + (data["middle_name"] ?? "") + " " + data["last_name"])).join(", ")
    else return ""
  }
  
  return (
    <div className="overflow-x-auto mx-4">
      <table className="table-auto w-full border-collapse border border-gray-800">
        <thead>
          <tr className="bg-black text-white">
            {columnsWithActions(columns).map((column: Column) => (
              <th key={column.accessor} className="px-4 py-2">
                {column.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="bg-white text-black">
          {data.map((row) => (
            <tr key={row.id}>
              {columnsWithActions(columns).map((column: Column) => (
                <td key={column.header} className="border border-gray-800 px-4 py-2">
                  {column.accessor === "actions" ? actionsRow : parseRowData(row[column.accessor])}
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