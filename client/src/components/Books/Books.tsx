import React, { useState, useEffect } from 'react'
import SearchBar from '../SearchBar/SearchBar'
import Table from '../Table/Table'
import Modal from '../Modal/Modal'
import swal from 'sweetalert'
import { getBooks } from '../../services/bookService'
import { useSearchParams } from 'react-router-dom'
import { FaPlusSquare } from 'react-icons/fa'

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

  const handleEditModal = (): void  => {
    setShowEditModal(true)
  }

  const handleDelete = () => {
    swal({
      title: "Are you sure?",
      text: "Once deleted, you will not be able to recover this item!",
      icon: "warning",
      buttons: true,
      dangerMode: true,
    })
    .then((willDelete) => {
      if (willDelete) {
        // Perform delete action here
        swal("Your item has been deleted!", {
          icon: "success",
        });
      } else {
        swal("Your item is safe!")
      }
    })
  }

  useEffect(() => {
    getBooks(searchParams.get("q") ?? "")
    .then(res => {
      setBooks(res.data.books)
    })
  },[searchParams])
  
  return (
    <section className="w-full h-screen flex flex-col">
        {showAddModal ? <Modal closeModal={() => setShowAddModal(false)}/> : null}
        {showEditModal ? <Modal closeModal={() => setShowEditModal(false)} /> : null}
        <h2 className="w-full text-zinc-600 text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center align-center'>
            <SearchBar placeholder='Enter ISBN13...'/>
            <FaPlusSquare 
              className="text-4xl text-slate-800 me-4"
              onClick={() => setShowAddModal(true)}/>
        </div>
        <Table 
          data={books} 
          columns={columns} 
          showModal={handleEditModal} 
          deleteItem={handleDelete}/>
    </section>
  )
}

export default Books