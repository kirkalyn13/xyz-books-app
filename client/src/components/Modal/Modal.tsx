import React from 'react'
import Button from '../Button/Button'
import { BiX, BiErrorCircle } from 'react-icons/bi'

interface ModalProps {
    title: string;
    closeModal: Function;
    children: JSX.Element;
    submit: Function;
    disableSubmit: boolean;
    error: string;
}

const Modal: React.FC<ModalProps> = ({ title, closeModal, children, submit, disableSubmit, error = "" }) => {
    const handleSubmit = (e): void => {
        e.preventDefault()
        submit()
    }

    const renderError = (): JSX.Element => (
        <div className="my-4 flex justify-center align-center">
            <BiErrorCircle className='text-red-400 text-2xl me-4' />
            <span className='text-red-400'>{error}</span>
        </div>
    )
    return (
        <form className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-80" onSubmit={handleSubmit}>
            <div className="w-screen md:w-1/4 h-2/3 md:h-auto mx-8 bg-white text-zinc-600 rounded-lg p-8 overflow-y-auto">
                <div className="flex flex-row justify-between align-center mb-4">
                    <h2 className="text-2xl font-bold">{title}</h2>
                    <BiX className="text-4xl" onClick={() => closeModal()}/>
                </div>
                {children}
                {error && renderError()}
                <div className='flex space-x-2 align-center justify-center'>
                    <input disabled={disableSubmit} type="submit" className={`${!disableSubmit ? "bg-sky-300" : "bg-gray-300"} px-8 py-2 rounded-lg focus:outline-none`}/>
                    <Button
                        text="Close" 
                        color="bg-gray-300"
                        handleOnClick={() => closeModal()}/>
                </div>
            </div>
        </form>
    )
}

export default Modal