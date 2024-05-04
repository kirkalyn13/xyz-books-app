import React, { useState } from 'react'
import { BiMenu } from 'react-icons/bi'

const LOGO_SRC = "/favicon.ico"
const Header : React.FC = () => {
  const [ showMenu, setShowMenu ] = useState(false)
  
  return (
    <nav className="w-screen p-4 z-10">
        <div className="container mx-auto flex flex-wrap justify-between items-center">
            <div className="flex items-center">
                <img
                    className="me-4"
                    src={LOGO_SRC}
                    width={50}
                    height={50}
                    alt="logo"/>
            </div>

            <div className="md:hidden mt-4">
                <button 
                    className="text-black"
                    onClick={() => setShowMenu(!showMenu)}>
                    <BiMenu className="text-4xl"/>
                </button>
            </div>

            <ul className="hidden md:flex space-x-4 align-center items-center">
                <li><a href="/" className="text-black hover:text-amber-500">Books</a></li>
                <li><a href="/authors" className="text-black hover:text-amber-500">Authors</a></li>
                <li><a href="/publishers" className="text-black hover:text-amber-500">Publishers</a></li>
            </ul>
        </div>

            { showMenu ?
                (<ul className="space-y-2 text-center">
                    <li><a href="/" className="text-black border-t-1 hover:text-amber-500">Books</a></li>
                    <li><a href="/authors" className="text-black hover:text-amber-500">Authors</a></li>
                    <li><a href="/publishers" className="text-black hover:text-amber-500">Publishers</a></li>
                </ul>)
            : null }

    </nav>
  )
}

export default Header