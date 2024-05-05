import React, { useEffect, useState } from 'react'
import Modal from '../../Modal/Modal'
import { Book } from '../../../types/book'
import { addBook, editBook, getBookByID } from '../../../services/bookService'
import useSearchID from '../../../hooks/useSearchID';

interface BookModalProps {
    title: string;
    closeModal: Function;
}

const BookModal: React.FC<BookModalProps> = ({ title, closeModal }) => {
  const [ book, setBook ] = useState<Book>({
    title: "",
    isbn13: "",
    isbn10: "",
    list_price: 0,
    publication_year: 0,
    image_url: "",
    edition: "",
    publisher_id: 0,
    authors: []
  })
  const { getSearchID } = useSearchID()

  const addBookHandler = (): void => {
    if (!disableSubmit) addBook(book)
        .then(() => closeModal())
    
  }

  const editBookHandler = (): void => {
      if (!disableSubmit) editBook(book.id!, book)
          .then(() => closeModal())
  }

  let disableSubmit: boolean = book.title === "" ||
    (book.isbn13 === "" && book.isbn10 == "") ||
    book.list_price === 0 ||
    book.publication_year === 0 

  const submitHandler: Function = () => title.toLowerCase().includes("add") ? addBookHandler() : editBookHandler()

  useEffect(() => {
      if (title.toLowerCase().includes("edit") && getSearchID() !== null) {
          getBookByID(getSearchID()).then((res) => {
              setBook(res.data.book)
          })
      }
  },[])

  return (
    <Modal disableSubmit={disableSubmit} title={title} closeModal={closeModal} submit={submitHandler}>
      <>
        <div className="flex flex-col md:flex-row justify-between my-4">
          <label className="text-md me-4 flex items-center">Title: </label>
          <input
              className="md:w-1/2 w-full my-2 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.title : ""}
              onChange={(e) => setBook({...book, title: e.target.value})}
              placeholder="Title..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-4">
          <label className="text-md me-4 flex items-center">IBN13: </label>
          <input
              className="md:w-1/2 w-full my-2 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.isbn13 : ""}
              onChange={(e) => setBook({...book, isbn13: e.target.value})}
              placeholder="ISBN13..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-4">
          <label className="text-md me-4 flex items-center">ISBN10: </label>
          <input
              className="md:w-1/2 w-full my-2 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.isbn10 : ""}
              onChange={(e) => setBook({...book, isbn10: e.target.value})}
              placeholder="ISBN10..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-4">
          <label className="text-md me-4 flex items-center">List Price: </label>
          <input
              className="md:w-1/2 w-full my-2 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="number"
              value={book ? book.list_price : 0}
              onChange={(e) => setBook({...book, list_price: parseInt(e.target.value)})}
              placeholder="List Price..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-4">
          <label className="text-md me-4 flex items-center">Publication Year: </label>
          <input
              className="md:w-1/2 w-full my-2 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="number"
              value={book ? book.publication_year : 0}
              onChange={(e) => setBook({...book, publication_year: parseInt(e.target.value)})}
              placeholder="Publication Year..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-4">
          <label className="text-md me-4 flex items-center">Image URL: </label>
          <input
              className="md:w-1/2 w-full my-2 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.image_url : ""}
              onChange={(e) => setBook({...book, image_url : e.target.value})}
              placeholder="Image URL..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-4">
          <label className="text-md me-4 flex items-center">Book Name: </label>
          <input
              className="md:w-1/2 w-full my-2 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.edition : ""}
              onChange={(e) => setBook({...book, edition: e.target.value})}
              placeholder="Edition..." />
        </div>
      </>
    </Modal>
  )
}

export default BookModal
