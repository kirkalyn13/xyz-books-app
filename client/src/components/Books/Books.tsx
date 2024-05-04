import React, { useState, useEffect } from 'react'
import SearchBar from '../SearchBar/SearchBar'
import Table from '../Table/Table'
import { getBooks } from '../../services/bookService'
import { useSearchParams } from 'react-router-dom'

const TITLE = "Book Management"

const columns = [
  { header: 'Title', accessor: 'title' },
  { header: 'Authors', accessor: 'authors' },
  { header: 'ISBN13', accessor: 'isbn13' },
  { header: 'ISBN 10', accessor: 'isbn10' },
  { header: 'Publication Year', accessor: 'publication_year' },
  { header: 'Publisher', accessor: 'publisher' },
  { header: 'Edition', accessor: 'edition' },
  { header: 'Price', accessor: 'list_price' },
]

const Books: React.FC = () => {
  const [ books, setBooks ] = useState([])
  const [ searchParams ] = useSearchParams()

  useEffect(() => {
    getBooks(searchParams.get("q") ?? "")
    .then(res => {
      setBooks(res.data.books)
    })
  },[searchParams])
  
  return (
    <section className="w-full flex flex-col">
        <h2 className="w-full text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center'>
            <SearchBar placeholder='Enter ISBN13...'/>
        </div>
        <Table data={books} columns={columns} />
    </section>
  )
}

export default Books