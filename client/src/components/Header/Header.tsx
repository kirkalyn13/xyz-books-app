import React from 'react'
import SearchBar from '../SearchBar/SearchBar'

const Header : React.FC = () => {
  return (
    <header className="w-100 m-4 z-10 pb-4 text-black flex flex-col justify-center align-center md:flex-row md:justify-between">
        <h1 className="text-center mb-4 md: mb-0 md:ms-4 text-3xl">XYZ Books</h1>
      <SearchBar />
    </header>
  )
}

export default Header