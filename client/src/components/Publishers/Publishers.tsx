import React from 'react'
import SearchBar from '../SearchBar/SearchBar'

const TITLE = "Publisher Management"

const Publishers: React.FC = () => {
  return (
    <section className="w-full flex flex-col">
        <h2 className="w-full text-3xl text-center">{TITLE}</h2>
    </section>
  )
}

export default Publishers