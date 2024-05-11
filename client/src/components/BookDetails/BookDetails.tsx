import React,{ useState, useEffect } from 'react'
import { BiArrowToLeft } from "react-icons/bi"
import { useNavigate, useParams } from 'react-router-dom'
import { Book } from '../../types/book'
import { getBookByISBN13 } from '../../services/bookService'
import NotFound from '../NotFound/NotFound'
import { Author } from '../../types/author'

const PLACEHOLDER = "/image-placeholder.svg"

const BookDetails: React.FC = () => {
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
    const [ isError, setIsError] = useState<boolean>(false)
    const { isbn } = useParams()
    const navigate = useNavigate()

    const goToHome = (): void => navigate("/")
    
    const formatAuthors = (authors: Author[]): string => {
        return authors
            .map((author: Author) => author.first_name + " " + (author["middle_name"] ?? "") + " " + author.last_name)
            .join(", ")
    }

    useEffect(() => {
        getBookByISBN13(isbn!)
        .then((res) => {
            if (res.response) {
                setIsError(res.response.data.error)
            } else {
                setBook(res.data.book)
            }
        })
    },[isbn])
    
    return (!isError ? 
            (<section className="w-screen h-auto flex flex-col">
            <div className='flex justify-center'>
                <div className="w-full md:w-1/3 container text-center space-y-4 mb-8 mx-2">
                    <div className='w-full mt-4 text-3xl flex justify-center align-center'>
                        <h2 className="w-full md:w-1/3 font-bold text-zinc-600 text-3xl text-center">
                            {book.title}
                        </h2>
                        <BiArrowToLeft 
                            className="text-4xl text-slate-800 hover:text-sky-300 ms-8 me-2"
                            onClick={() => goToHome()} />
                    </div>
                    <div className='w-full flex justify-center align-center'>
                        <img 
                            width="300"
                            height="500"
                            className='border border-md my-4' 
                            src={!book.image_url ? PLACEHOLDER : book.image_url } 
                            alt="book-image" />
                    </div>
                    <div className='w-full flex justify-between align-center'>
                        <span className='text-xl font-bold'>ISBN 13:</span>
                        <span className='text-xl'>{book.isbn13}</span>
                    </div>
                    <div className='w-full flex justify-between align-center'>
                        <span className='text-xl font-bold'>ISBN 10:</span>
                        <span className='text-xl'>{book.isbn10}</span>
                    </div>
                    <div className='w-full flex justify-between align-center'>
                        <span className='text-xl font-bold'>List Price:</span>
                        <span className='text-xl'>{book.list_price}</span>
                    </div>
                    <div className='w-full flex justify-between align-center'>
                        <span className='text-xl font-bold'>Publication Year:</span>
                        <span className='text-xl'>{book.publication_year}</span>
                    </div>
                    <div className='w-full flex justify-between align-center'>
                        <span className='text-xl font-bold'>Edition:</span>
                        <span className='text-xl'>{book.edition}</span>
                    </div>
                    <div className='w-full flex justify-between align-center'>
                        <span className='text-xl font-bold'>Publisher:</span>
                        <span className='text-xl'>{book.publisher.name}</span>
                    </div>
                    <div className='w-full flex justify-between align-center'>
                        <span className='text-xl font-bold'>Authors:</span>
                        <span className='text-xl'>{formatAuthors(book.authors!)}</span>
                    </div>
                </div>
            </div>
        </section>) : <NotFound />
    )
}

export default BookDetails