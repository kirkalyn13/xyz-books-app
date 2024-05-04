import React from 'react'
import SearchBar from '../SearchBar/SearchBar'
import Table from '../Table/Table'
import { useSearchParams } from 'react-router-dom'

const TITLE = "Book Management"

const data = [
    { id: 1, name: 'John Doe', age: 30 },
    { id: 2, name: 'Jane Smith', age: 25 },
    { id: 3, name: 'Bob Johnson', age: 40 },
  ];

  const columns = [
    { header: 'ID', accessor: 'id' },
    { header: 'Name', accessor: 'name' },
    { header: 'Age', accessor: 'age' },
  ];

const Books: React.FC = () => {
  const [ searchParams ] = useSearchParams()
  
  return (
    <section className="w-full flex flex-col">
        <h2 className="w-full text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center'>
            <SearchBar placeholder='Enter ISBN13...'/>
        </div>
        <Table data={data} columns={columns} />
    </section>
  )
}

export default Books