import React from 'react'
import { FaEdit, FaTrash } from 'react-icons/fa';

type Column = {
    header: string;
    accessor: string;
}

interface TableProps {
  data: any[]
  columns: Column[]
  handleEdit: Function
  deleteItem: Function
}


const Table: React.FC<TableProps> = ({ data, columns, handleEdit, deleteItem }) => {
  const columnsWithActions = (columns: Column[]) => [...columns, { header: 'Actions', accessor: 'actions' }]

  const getActionsRow = (id: number): JSX.Element => { 
    return (
      <div className="w-24 flex justify-around align-center">
        <FaEdit className='text-xl text-zinc-600 hover:text-sky-300' onClick={() => handleEdit(id)}/>
        <FaTrash className='text-xl text-zinc-600 hover:text-sky-300' onClick={() => deleteItem(id)}/>
      </div>
    )
  }

  const getRowData = (accessor: string, rowData: any) => {
    if (accessor === "title") {
      return (
        <a href={`/books/${rowData.isbn13}`}
          className='underline hover:text-sky-300'>
          {rowData.title}
        </a>
      )
    } else if (accessor === "actions") {
      return getActionsRow(rowData.id)
    }
    return parseRowData(rowData[accessor])
  }

  const parseRowData = (rowData: any): string => {
    if (typeof rowData !== 'object') return rowData
    else if (rowData["name"]) return rowData["name"]
    else if (Array.isArray(rowData)) return rowData.map((data: any) => (data["first_name"] + " " + (data["middle_name"] ?? "") + " " + data["last_name"])).join(", ")
    else return ""
  }
  
  return (
    <div className="overflow-x-auto overflow-y-auto mx-4">
      <table className="table-auto w-full border-collapse border border-gray-800">
        <thead>
          <tr className="bg-slate-800 text-white">
            {columnsWithActions(columns).map((column: Column) => (
              <th key={column.accessor} className={`px-4 py-2 ${column.accessor === "actions" && "w-32"}`} >
                {column.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="bg-white text-zinc-600">
          {data.map((row) => (
            <tr key={row.id}>
              {columnsWithActions(columns).map((column: Column) => (
                <td key={column.header} className="border border-gray-800 px-4 py-2">
                  {getRowData(column.accessor, row)}
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