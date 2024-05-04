import React from 'react'
import SearchBar from '../SearchBar/SearchBar'
import { useSearchParams } from 'react-router-dom'

const TITLE = "Author Management"

const Authors: React.FC = () => {
  const [ searchParams ] = useSearchParams()

  return (
    <section className="w-full flex flex-col">
        <h2 className="w-full text-3xl text-center">{TITLE}</h2>
        <div className='w-full mt-4 text-3xl flex justify-center'>
            <SearchBar />
        </div>
    </section>
  )
}

export default Authors