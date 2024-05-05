import React from 'react'
import Modal from '../../Modal/Modal'

interface AuthorModalProps {
    title: string;
    closeModal: Function;
}

const AuthorModal: React.FC<AuthorModalProps> = ({ title, closeModal }) => {
  return (
    <Modal title={title} closeModal={closeModal}>
        <h3>Add Edit Author Modal</h3>
    </Modal>
  )
}

export default AuthorModal