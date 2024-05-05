import React, { useEffect, useState } from 'react'
import SearchBar from '../SearchBar/SearchBar'
import Table from '../Table/Table'
import swal from 'sweetalert'
import { deleteAuthor, getAuthors } from '../../services/authorService'
import { useSearchParams } from 'react-router-dom'
import { FaPlusSquare } from 'react-icons/fa'
import AuthorModal from './AuthorModal/AuthorModal'
import useSearchID from '../../hooks/useSearchID'

const TITLE = "Author Management"

const columns = [
  { header: 'First Name', accessor: 'first_name' },
  { header: 'Middle Name', accessor: 'middle_name' },
  { header: 'Last Name', accessor: 'last_name' },
]

const Authors: React.FC = () => {
  const [ showAddModal, setShowAddModal] = useState(false)
  const [ showEditModal, setShowEditModal] = useState(false)
  const [ authors, setAuthors ] = useState([])
  const [ searchParams ] = useSearchParams()
  const { updateID, clearID } = useSearchID()

  const loadAuthors = ():void => {
    getAuthors(searchParams.get("q") ?? "")
    .then(res => {
      setAuthors(res.data.authors)
    })
  }

  const handleEdit = (id: number): void => {
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
        deleteAuthor(id.toString()).then(() => {
          swal("Author has been deleted.", {
            icon: "success",
          })
          loadAuthors()
        })
      } else {
        swal("Author retained.")
      }
    })
  }
  
  useEffect(() => {
    loadAuthors()
  },[searchParams, showAddModal, showEditModal])

  return (
    <section className="w-full h-screen flex flex-col"> 
        {showAddModal ? <AuthorModal title="Add Author" closeModal={() => closeEditModal()}/> : null}
        {showEditModal ? <AuthorModal title="Edit Author" closeModal={() => setShowEditModal(false)}/> : null}
        <h2 className="w-full text-zinc-600 text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center'>
            <SearchBar />
            <FaPlusSquare 
              className="text-4xl text-slate-800 hover:bg-sky-300"
              onClick={() => setShowAddModal(true)}/>
        </div>
        <Table 
          data={authors} 
          columns={columns} 
          handleEdit={handleEdit} 
          deleteItem={handleDelete}/>
    </section>
  )
}

export default Authors