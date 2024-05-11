import React, { useEffect, useState } from 'react'
import Modal from '../../Modal/Modal'
import { Book } from '../../../types/book'
import { addBook, editBook, getBookByID } from '../../../services/bookService'
import { getAuthors } from '../../../services/authorService'
import { getPublishers } from '../../../services/publisherService'
import useSearchID from '../../../hooks/useSearchID';
import { Publisher } from '../../../types/publisher'
import { Author } from '../../../types/author'
import Chip from '../../Chip/Chip'
import { sanitizeData } from '../../../utils/sanitizeData'

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
    publisher: {name: ""},
    authors: []
  })
  const [ authors, setAuthors ] = useState<Author[]>([])
  const [ selectedAuthor, setSelectedAuthor ] = useState("")
  const [ selectedAuthors, setSelectedAuthors ] = useState<Author[]>([])
  const [ publishers, setPublishers ] = useState<Publisher[]>([])
  const [ error, setError ] = useState<string>("")
  const { getSearchID } = useSearchID()

  const addBookHandler = (): void => {
    setError("")
    sanitizeData(book)
    if (!disableSubmit) addBook(book)
        .then((res) => {
          if (res.response) {
            setError(res.response.data.error)
          } else {
            setTimeout(() => closeModal(), 500)
          }
    })
  }

  const editBookHandler = (): void => {
    setError("")
    sanitizeData(book)
    if (!disableSubmit) editBook(book.id!, book)
        .then((res) => {
          if (res.response) {
              setError(res.response.data.error)
          } else {
              closeModal()
          }
      })
  }

  const removeAuthor = (id: string) => {
    const newAuthorList = selectedAuthors.filter((author: Author) => author.id !== id);
    setSelectedAuthors(newAuthorList);
  }

  const getSelectedAuthor = (): Author => {
    return authors.filter((author: Author) => parseInt(author.id!) === parseInt(selectedAuthor))[0]
  }

  const getAuthorName = (author: Author): string => author.first_name + " " + (author.middle_name ? author.middle_name + " " : "") + author.last_name

  let disableSubmit: boolean = book.title === "" ||
    (book.isbn13 === "" && book.isbn10 == "") ||
    book.list_price === 0 ||
    book.publication_year.toString().length !== 4 ||
    book.publisher_id === 0 ||
    book.authors.length === 0

  const submitHandler: Function = () => title.toLowerCase().includes("add") ? addBookHandler() : editBookHandler()

  useEffect(() => {
      getAuthors("").then((res) => {
        setAuthors([{
          first_name: "",
          last_name: "",
          middle_name: ""
      }, ...res.data.authors])
      })
      getPublishers("").then((res) => {
        setPublishers([{name: ""}, ...res.data.publishers])
      })
      if (title.toLowerCase().includes("edit") && getSearchID() !== null) {
          getBookByID(getSearchID()).then((res) => {
              setBook(res.data.book)
              if (res.data.book.authors.length > 0) setSelectedAuthors(res.data.book.authors)
          })
      }
  },[])

  useEffect(() => {
    if (selectedAuthor === "") return
    if (selectedAuthors.some((author: Author) => parseInt(author.id!) === parseInt(selectedAuthor))) return
    setSelectedAuthors((prev) => [...prev, getSelectedAuthor()])
  },[selectedAuthor])
  
  useEffect(() => {
    setBook({...book, authors: selectedAuthors})
  },[selectedAuthors])

  return (
    <Modal disableSubmit={disableSubmit} title={title} closeModal={closeModal} submit={submitHandler} error={error}>
      <>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">Title: </label>
          <input
              className="md:w-1/2 w-full my-1 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.title : ""}
              onChange={(e) => setBook({...book, title: e.target.value})}
              placeholder="Title..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">IBN13: </label>
          <input
              className="md:w-1/2 w-full my-1 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.isbn13 : ""}
              onChange={(e) => setBook({...book, isbn13: e.target.value})}
              placeholder="ISBN13..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">ISBN10: </label>
          <input
              className="md:w-1/2 w-full my-1 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.isbn10 : ""}
              onChange={(e) => setBook({...book, isbn10: e.target.value})}
              placeholder="ISBN10..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">List Price: </label>
          <input
              className="md:w-1/2 w-full my-1 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="number"
              value={book ? book.list_price : 0}
              onChange={(e) => setBook({...book, list_price: parseInt(e.target.value)})}
              placeholder="List Price..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">Publication Year: </label>
          <input
              className="md:w-1/2 w-full my-1 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="number"
              value={book ? book.publication_year : 0}
              onChange={(e) => setBook({...book, publication_year: parseInt(e.target.value)})}
              placeholder="Publication Year..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">Image URL: </label>
          <input
              className="md:w-1/2 w-full my-1 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.image_url : ""}
              onChange={(e) => setBook({...book, image_url : e.target.value})}
              placeholder="Image URL..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">Edition: </label>
          <input
              className="md:w-1/2 w-full my-1 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              type="text"
              value={book ? book.edition : ""}
              onChange={(e) => setBook({...book, edition: e.target.value})}
              placeholder="Edition..." />
        </div>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">Publisher: </label>
          <select 
              className="md:w-1/2 w-full my-1 py-1 px-2
              text-sm text-black border rounded-lg 
              focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 overflow-y-auto"
              value={book.publisher_id} 
              onChange={(e) => setBook({...book, publisher_id: parseInt(e.target.value)})}
              >
            {publishers.map((publisher: Publisher) => (
              <option key={publisher.id} value={publisher.id}>
                {publisher.name}
              </option>
            ))}
          </select>
        </div>
        <div className="flex flex-col md:flex-row justify-between my-1">
          <label className="text-md me-4 flex items-center">Authors: </label>
          <select 
                className="md:w-1/2 w-full my-1 py-1 px-2
                text-sm text-black border rounded-lg 
                focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 overflow-y-auto"
                value={selectedAuthor} 
                onChange={(e) => setSelectedAuthor(e.target.value)}
                >
              {authors.map((author: Author) => (
                <option key={author.id} value={author.id}>
                  {getAuthorName(author)}
                </option>
              ))}
            </select>
        </div>
        <div className="flex flex-wrap gap-2 mb-4 overflow-y-scroll">
            {selectedAuthors.map((author: Author) => (
              <Chip 
                key={author.id}
                id={author.id!} 
                name={getAuthorName(author)}
                removeChip={() => removeAuthor(author.id!)}
                />
            ))}
          </div>
      </>
    </Modal>
  )
}

export default BookModal
