import React, { useState, useEffect } from 'react'
import SearchBar from '../SearchBar/SearchBar'
import Table from '../Table/Table'
import swal from 'sweetalert'
import { deleteBook, getBooks } from '../../services/bookService'
import { useSearchParams } from 'react-router-dom'
import { FaPlusSquare } from 'react-icons/fa'
import BookModal from './BookModal/BookModal'
import useSearchID from '../../hooks/useSearchID'

const TITLE = "Book Management"

const columns = [
  { header: 'Title', accessor: 'title' },
  { header: 'Authors', accessor: 'authors' },
  { header: 'ISBN13', accessor: 'isbn13' },
  { header: 'ISBN10', accessor: 'isbn10' },
  { header: 'Publication Year', accessor: 'publication_year' },
  { header: 'Publisher', accessor: 'publisher' },
  { header: 'Edition', accessor: 'edition' },
  { header: 'Price', accessor: 'list_price' },
]

const Books: React.FC = () => {
  const [ showAddModal, setShowAddModal] = useState(false)
  const [ showEditModal, setShowEditModal] = useState(false)
  const [ books, setBooks ] = useState([])
  const [ searchParams ] = useSearchParams()
  const { updateID, clearID } = useSearchID()

  const loadBooks = (): void => {
    getBooks(searchParams.get("q") ?? "")
    .then(res => {
      setBooks(res.data.books)
    })
  }
  
  const handleEdit = (id: number): void  => {
    updateID(id.toString())
    setShowEditModal(true)
  }

  const closeEditModal = (): void => {
    clearID()
    setShowAddModal(false)
  }

  const handleDelete = (id: number) => {
    swal({
      title: "Are you sure?",
      text: "Once deleted, you will not be able to recover this item.",
      icon: "warning",
      buttons: true,
      dangerMode: true,
    })
    .then((willDelete) => {
      if (willDelete) {
        deleteBook(id.toString()).then(() => {
          swal("Book has been deleted", {
            icon: "success",
          })
          loadBooks()
        })
      } else {
        swal("Book retained.")
      }
    })
  }

  useEffect(() => {
    loadBooks()
  },[searchParams, showAddModal, showEditModal])
  
  return (
    <section className="w-full h-screen flex flex-col">
        {showAddModal ? <BookModal title="Add Book" closeModal={() => closeEditModal()}/> : null}
        {showEditModal ? <BookModal title="Edit Book" closeModal={() => setShowEditModal(false)}/>  : null}
        <h2 className="w-full text-zinc-600 text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center align-center'>
            <SearchBar placeholder='Enter ISBN13...'/>
            <FaPlusSquare 
              className="text-4xl text-slate-800 me-4 hover:text-sky-300"
              onClick={() => setShowAddModal(true)}/>
        </div>
        <Table 
          data={books} 
          columns={columns} 
          handleEdit={handleEdit}  
          deleteItem={handleDelete}/>
    </section>
  )
}

export default Books