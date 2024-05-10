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
        return err
    }
}

export const getPublisherByID = async (id: string): Promise<any> => {
    try {
        return await axios.get(buildUri(getEndpoint("publishers", id)))
    } catch(err) {
        return err
    }
}

export const addPublisher = async (publisher: Publisher): Promise<any> => {
    try {
        return await axios.post(buildUri(getEndpoint("publishers", "", "")), publisher)
    } catch(err) {
        return err
    }
}

export const editPublisher = async (id: string, publisher: Publisher): Promise<any> => {
    try {
        return await axios.put(buildUri(getEndpoint("publishers", id, "")), publisher)
    } catch(err) {
        return err
    }
}

export const deletePublisher = async (id: string): Promise<any> => {
    try {
        return await axios.delete(buildUri(getEndpoint("publishers", id)))
    } catch(err) {
        return err
    }
}