import { buildUri } from "../utils/uriBuilder"
import { getEndpoint } from "../utils/getEndpoint"
import { Publisher } from "../types/publisher"
import axios from 'axios'

export const getPublishers = async (q: string): Promise<any> => {
    try {
        const queryParams = {
            q
        }
        return await axios.get(buildUri(getEndpoint("publishers", ""), queryParams))
    } catch(err) {
        console.error(err)
    }
}

export const getPublisherByID = async (id: string): Promise<any> => {
    try {
        return await axios.get(buildUri(getEndpoint("publishers", id)))
    } catch(err) {
        console.error(err)
    }
}

export const addBook = async (publisher: Publisher): Promise<any> => {
    try {
        return await axios.post(buildUri(getEndpoint("publishers", "", "")), publisher)
    } catch(err) {
        console.error(err)
    }
}

export const editBook = async (publisher: Publisher): Promise<any> => {
    try {
        return await axios.put(buildUri(getEndpoint("publishers", "", "")), publisher)
    } catch(err) {
        console.error(err)
    }
}

export const deleteBook = async (id: string): Promise<any> => {
    try {
        return await axios.delete(buildUri(getEndpoint("publishers", id)))
    } catch(err) {
        console.error(err)
    }
}