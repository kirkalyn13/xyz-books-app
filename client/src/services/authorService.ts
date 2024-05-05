import { buildUri } from "../utils/uriBuilder"
import { getEndpoint } from "../utils/getEndpoint"
import { Author } from "../types/author"
import axios from 'axios'

export const getAuthors = async (q: string): Promise<any> => {
    try {
        const queryParams = {
            q
        }
        return await axios.get(buildUri(getEndpoint("authors", ""), queryParams))
    } catch(err) {
        console.error(err)
    }
}

export const getAuthorByID = async (id: string): Promise<any> => {
    try {
        return await axios.get(buildUri(getEndpoint("authors", id)))
    } catch(err) {
        console.error(err)
    }
}

export const addAuthor = async (author: Author): Promise<any> => {
    try {
        return await axios.post(buildUri(getEndpoint("authors", "", "")), author)
    } catch(err) {
        console.error(err)
    }
}

export const editAuthor = async (id: string, author: Author): Promise<any> => {
    try {
        return await axios.put(buildUri(getEndpoint("authors", id, "")), author)
    } catch(err) {
        console.error(err)
    }
}

export const deleteAuthor = async (id: string): Promise<any> => {
    try {
        return await axios.delete(buildUri(getEndpoint("authors", id)))
    } catch(err) {
        console.error(err)
    }
}