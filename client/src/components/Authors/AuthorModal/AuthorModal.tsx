import React, { useEffect, useState } from 'react'
import Modal from '../../Modal/Modal'
import { addAuthor, editAuthor, getAuthorByID } from '../../../services/authorService'
import { Author } from '../../../types/author'
import useSearchID from '../../../hooks/useSearchID'
import { sanitizeData } from '../../../utils/sanitizeData'

interface AuthorModalProps {
    title: string;
    closeModal: Function;
    data?: Author;
}

const AuthorModal: React.FC<AuthorModalProps> = ({ title, closeModal, data }) => {
    const [author, setAuthor] = useState<Author>({
        first_name: "",
        last_name: "",
        middle_name: ""
    })
    const [ error, setError ] = useState<string>("")
    const { getSearchID } = useSearchID()

    const addAuthorHandler = (): void => {
        setError("")
        sanitizeData(author)
        if (!disableSubmit) addAuthor(author)
            .then((res) => {
                if (res.response) {
                    setError(res.response.data.error)
                } else {
                    closeModal()
                }
            })
    }

    const editAuthorHandler = (): void => {
        setError("")
        sanitizeData(author)
        if (!disableSubmit) editAuthor(author.id!, author)
            .then((res) => {
                if (res.response) {
                    setError(res.response.data.error)
                } else {
                    closeModal()
                }
            })
    }

    let disableSubmit: boolean = author.first_name === "" || author.last_name === ""
    const submitHandler: Function = () => title.toLowerCase().includes("add") ? addAuthorHandler() : editAuthorHandler()

    useEffect(() => {
        if (title.toLowerCase().includes("edit") && getSearchID() !== null) {
            getAuthorByID(getSearchID()).then((res) => {
                setAuthor(res.data.author)
            })
        }
    },[])

    return (
        <Modal disableSubmit={disableSubmit} title={title} closeModal={closeModal} submit={submitHandler} error={error}>
            <>
            <div className="flex flex-col md:flex-row justify-between my-4">
                <label className="text-md me-4 flex items-center">First Name: </label>
                <input
                    className="md:w-1/2 w-full my-2 py-1 px-2
                    text-sm text-black border rounded-lg 
                    focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    type="text"
                    value={author ? author.first_name : ""}
                    onChange={(e) => setAuthor({...author, first_name: e.target.value})}
                    placeholder="First Name..." />
            </div>
            <div className="flex flex-col md:flex-row justify-between my-4">
                <label className="text-md me-4 flex items-center">Middle Name: </label>
                <input
                    className="md:w-1/2 w-full my-2 py-1 px-2
                    text-sm text-black border rounded-lg 
                    focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    type="text"
                    value={author ? author.middle_name : ""}
                    onChange={(e) => setAuthor({...author, middle_name: e.target.value})}
                    placeholder="Middle Name..." />
            </div>
            <div className="flex flex-col md:flex-row justify-between my-4">
                <label className="text-md me-4 flex items-center">Last Name: </label>
                <input
                    className="md:w-1/2 w-full my-2 py-1 px-2
                    text-sm text-black border rounded-lg 
                    focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    type="text"
                    value={author ? author.last_name : ""}
                    onChange={(e) => setAuthor({...author, last_name: e.target.value})}
                    placeholder="Last Name..." />
            </div>
            </>
        </Modal>
    )
}

export default AuthorModal