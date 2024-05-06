import { FaSearch } from "react-icons/fa"
import React from "react";
import { useSearchParams } from 'react-router-dom'

interface SearchBarProps {
    placeholder?: string
}

const DEFAULT_PLACEHOLDER = "Type to search..."

const SearchBar: React.FC<SearchBarProps> = ({ placeholder = DEFAULT_PLACEHOLDER }) => {
    const [ searchParams, setSearchParams ] = useSearchParams({q: ''})

    return (
        <div
            className="w-full md:w-1/3 bg-white md:my-8 md:mt-0 mb-4 mx-4 border-2 border-slate-900 rounded-sm
                flex align-start justify-start">
            <FaSearch className="text-xl m-2 text-zinc-600"/>
            <input 
                className="text-xl outline-none focus:outline-none text-slate-900"
                type="text"
                placeholder={placeholder}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => setSearchParams((prev: URLSearchParams) => {
                    prev.set("q", e.target.value)
                    return prev
                }, { replace: true })}
                />
        </div>
    )
}

export default SearchBar