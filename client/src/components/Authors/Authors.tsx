import React, { useEffect, useState } from 'react'
import SearchBar from '../SearchBar/SearchBar'
import { getAuthors } from '../../services/authorService'
import { useSearchParams } from 'react-router-dom'
import Table from '../Table/Table'

const TITLE = "Author Management"

const columns = [
  { header: 'First Name', accessor: 'first_name' },
  { header: 'Middle Name', accessor: 'middle_name' },
  { header: 'Last Name', accessor: 'last_name' },
]

const Authors: React.FC = () => {
  const [ authors, setAuthors ] = useState([])
  const [ searchParams ] = useSearchParams()
  
  useEffect(() => {
    getAuthors(searchParams.get("q") ?? "")
    .then(res => {
      setAuthors(res.data.authors)
    })
  },[searchParams])

  return (
    <section className="w-full h-screen flex flex-col">
        <h2 className="w-full text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center'>
            <SearchBar />
        </div>
        <Table data={authors} columns={columns} />
    </section>
  )
}

export default Authors