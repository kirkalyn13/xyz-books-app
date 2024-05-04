import React, { useEffect, useState } from 'react'
import SearchBar from '../SearchBar/SearchBar'
import Table from '../Table/Table'
import { getPublishers } from '../../services/publisherService'
import { useSearchParams } from 'react-router-dom'

const TITLE = "Publisher Management"

const columns = [
  { header: 'Name', accessor: 'name' },
]

const Publishers: React.FC = () => {
  const [ publishers, setPublishers ] = useState([])
  const [ searchParams ] = useSearchParams()
  
  useEffect(() => {
    getPublishers(searchParams.get("q") ?? "")
    .then(res => {
      setPublishers(res.data.publishers)
    })
  },[searchParams])

  return (
    <section className="w-full flex flex-col">
        <h2 className="w-full text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center'>
            <SearchBar />
        </div>
        <Table data={publishers} columns={columns} />
    </section>
  )
}

export default Publishers