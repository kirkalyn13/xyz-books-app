import React from 'react'
import Button from '../Button/Button'
import { BiX } from 'react-icons/bi'

interface ModalProps {
    closeModal: Function;
}

const Modal: React.FC<ModalProps> = ({ closeModal }) => {
    const handleSubmit = () => {

    }
    
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-80">
        <div className="w-screen md:w-1/4 h-2/3 md:h-auto mx-8 bg-white text-zinc-600 rounded-lg p-8 overflow-y-auto">
            <div className="flex flex-row justify-between align-center mb-4">
                <h2 className="text-2xl font-bold">Modal</h2>
                <BiX className="text-4xl" onClick={() => closeModal()}/>
            </div>
            <div className='flex space-x-2 align-center justify-center'>
                <Button text="Save" color="bg-sky-300" handleOnClick={() => handleSubmit()}/>
                <Button
                    text="Close" 
                    color="bg-gray-300"
                    handleOnClick={() => closeModal()}/>
        </div>
        </div>
    </div>
  )
}

export default Modal