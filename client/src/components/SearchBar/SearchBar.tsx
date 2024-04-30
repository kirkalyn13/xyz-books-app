import { FaSearch } from "react-icons/fa"
import { useState, useEffect } from "react";
import React from "react";

const SearchBar: React.FC = () => {
    return (
        <div
        className="bg-white w-100 mt-8 md:mt-0 border-2 border-black rounded-md
            flex align-center justify-start">
        <FaSearch className="text-xl mt-4 ms-2 text-black"/>
        <input 
            className="ps-4 text-xl focus:outline-none text-black"
            type="text"
            placeholder="Enter ISBN 13..."
            />
        </div>
    )
}

export default SearchBar