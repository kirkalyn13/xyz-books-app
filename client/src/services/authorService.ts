import { buildUri } from "../utils/uriBuilder"
import { getEndpoint } from "../utils/getEndpoint"
import { Author } from "../types/author"
import axios from 'axios'

export const getAuthors = async (id: string, searchQuery: string): Promise<any> => {
    try {
        const queryParams = {
            searchQuery
        }
        return await axios.get(buildUri(getEndpoint("authors", id), queryParams))
    } catch(err) {
        console.error(err)
    }
}

export const getPublisherByID = async (id: string): Promise<any> => {
    try {
        return await axios.get(buildUri(getEndpoint("authors", id)))
    } catch(err) {
        console.error(err)
    }
}

export const addBook = async (author: Author): Promise<any> => {
    try {
        return await axios.post(buildUri(getEndpoint("authors", "", "")), author)
    } catch(err) {
        console.error(err)
    }
}

export const editBook = async (author: Author): Promise<any> => {
    try {
        return await axios.put(buildUri(getEndpoint("authors", "", "")), author)
    } catch(err) {
        console.error(err)
    }
}

export const deleteBook = async (id: string): Promise<any> => {
    try {
        return await axios.delete(buildUri(getEndpoint("authors", id)))
    } catch(err) {
        console.error(err)
    }
}