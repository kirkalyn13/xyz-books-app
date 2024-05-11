import { buildUri } from "../utils/uriBuilder"
import { getEndpoint } from "../utils/getEndpoint"
import { Book } from "../types/book"
import axios from 'axios'

export const getBooks = async (q: string): Promise<any> => {
    try {
        const queryParams = {
            q
        }
        return await axios.get(buildUri(getEndpoint("books", ""), queryParams))
    } catch(err) {
        return err
    }
}

export const getBookByISBN13 = async (isbn13: string): Promise<any> => {
    try {
        return await axios.get(buildUri(getEndpoint("books", isbn13, "isbn13")))
    } catch(err) {
        return err
    }
}

export const getBookByID = async (id: string): Promise<any> => {
    try {
        return await axios.get(buildUri(getEndpoint("books", id)))
    } catch(err) {
        return err
    }
}

export const addBook = async (book: Book): Promise<any> => {
    try {
        return await axios.post(buildUri(getEndpoint("books", "", "")), book)
    } catch(err) {
        return err
    }
}

export const editBook = async (id: string, book: Book): Promise<any> => {
    try {
        return await axios.put(buildUri(getEndpoint("books", id, "")), book)
    } catch(err) {
        return err
    }
}

export const deleteBook = async (id: string): Promise<any> => {
    try {
        return await axios.delete(buildUri(getEndpoint("books", id)))
    } catch(err) {
        return err
    }
}