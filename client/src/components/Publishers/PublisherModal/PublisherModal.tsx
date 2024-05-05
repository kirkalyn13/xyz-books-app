import React, { useEffect, useState } from 'react'
import Modal from '../../Modal/Modal'
import { Publisher } from '../../../types/publisher';
import { addPublisher, editPublisher } from '../../../services/publisherService';

interface PublisherModalProps {
    title: string;
    closeModal: Function;
    data?: Publisher;
}

const PublisherModal: React.FC<PublisherModalProps> = ({ title, closeModal, data }) => {
    const [publisher, setPublisher] = useState<Publisher | null>(null)

    const addPublisherHandler = (): void => {
        console.log(publisher)
        if (publisher) addPublisher(publisher)
        closeModal()
    }

    const editPublisherHandler = (): void => {
        if (publisher) editPublisher(publisher)
    }

    const submitHandler: Function = () => title.toLowerCase().includes("add") ? addPublisherHandler() : editPublisherHandler()

    useEffect(() => {
        if (data) setPublisher(data)
    },[])

    return (
        <Modal title={title} closeModal={closeModal} submit={submitHandler}>
            <div className="flex flex-col md:flex-row justify-between my-4">
                <label className="text-md me-4 flex items-center">Publisher Name: </label>
                <input
                    className="md:w-1/2 w-full my-2 py-1 px-2
                    text-sm text-black border rounded-lg 
                    focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    type="text"
                    value={publisher ? publisher.name : ""}
                    onChange={(e) => setPublisher({name: e.target.value})}
                    placeholder="First Name..." />
            </div>
        </Modal>
    )
}

export default PublisherModal