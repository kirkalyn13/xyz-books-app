import React from 'react'
import Modal from '../../Modal/Modal'

interface BookModalProps {
    title: string;
    closeModal: Function;
}

const BookModal: React.FC<BookModalProps> = ({ title, closeModal }) => {
  return (
    <Modal title={title} closeModal={closeModal}>
        <h3>Add Edit Book Modal</h3>
    </Modal>
  )
}

export default BookModal