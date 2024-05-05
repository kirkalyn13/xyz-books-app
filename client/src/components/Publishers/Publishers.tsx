import React, { useEffect, useState } from 'react'
import SearchBar from '../SearchBar/SearchBar'
import Table from '../Table/Table'
import swal from 'sweetalert'
import { deletePublisher, getPublishers } from '../../services/publisherService'
import { useSearchParams } from 'react-router-dom'
import { FaPlusSquare } from 'react-icons/fa'
import PublisherModal from './PublisherModal/PublisherModal'

const TITLE = "Publisher Management"

const columns = [
  { header: 'Name', accessor: 'name' },
]

const Publishers: React.FC = () => {
  const [ showAddModal, setShowAddModal] = useState(false)
  const [ showEditModal, setShowEditModal] = useState(false)
  const [ publishers, setPublishers ] = useState([])
  const [ searchParams ] = useSearchParams()

  const loadPublishers = (): void => {
    getPublishers(searchParams.get("q") ?? "")
    .then(res => {
      setPublishers(res.data.publishers)
    })
  }

  const handleEditModal = (): void => {
    setShowEditModal(true)
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
        deletePublisher(id).then(() => {
          swal("Publisher item has been deleted!", {
            icon: "success",
          })
          loadPublishers()
        })
      } else {
        swal("Publisher retained.")
      }
    })
  }
  
  useEffect(() => {
    loadPublishers()
  },[searchParams, showAddModal, showEditModal])

  return (
    <section className="w-full h-screen flex flex-col">
        {showAddModal ? <PublisherModal title="Add Publisher" closeModal={() => setShowAddModal(false)}/> : null}
        {showEditModal ? <PublisherModal title="Edit Publisher" closeModal={() => setShowEditModal(false)}/> : null}
        <h2 className="w-full text-zinc-600 text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center'>
            <SearchBar />
            <FaPlusSquare
              className="text-4xl text-slate-800"
              onClick={() => setShowAddModal(true)}/>
        </div>
        <Table 
          data={publishers} 
          columns={columns} 
          showModal={handleEditModal} 
          deleteItem={handleDelete}/>
    </section>
  )
}

export default Publishers