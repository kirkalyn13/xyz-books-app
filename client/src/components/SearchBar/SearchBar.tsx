import { FaSearch } from "react-icons/fa"
import { useState, useEffect } from "react";
import React from "react";

const SearchBar: React.FC = () => {
    return (
        <div
            className="w-full md:w-1/3 bg-white my-8 md:mt-0 mx-4 border-2 border-black rounded-sm
                flex align-center justify-start">
            <FaSearch className="text-xl m-2 text-black"/>
            <input 
                className="text-xl outline-none focus:outline-none text-slate-900"
                type="text"
                placeholder="Type to search..."
                />
        </div>
    )
}

export default SearchBar