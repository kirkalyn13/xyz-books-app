import React, { useState } from 'react'
import { BiMenu } from 'react-icons/bi'

const LOGO_SRC = "/favicon.ico"
const TITLE = "XYZ Books"

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
                <h1 className="text-2xl text-bold text-zinc-600">{TITLE}</h1>
            </div>

            <div className="md:hidden mt-4">
                <button 
                    className="text-zinc-600"
                    onClick={() => setShowMenu(!showMenu)}>
                    <BiMenu className="text-4xl"/>
                </button>
            </div>

            <ul className="hidden md:flex space-x-4 align-center items-center">
                <li><a href="/" className="text-zinc-600 hover:text-sky-300">BOOKS</a></li>
                <li><a href="/authors" className="text-zinc-600 hover:text-sky-300">AUTHORS</a></li>
                <li><a href="/publishers" className="text-zinc-600 hover:text-sky-300">PUBLISHERS</a></li>
            </ul>
        </div>

            { showMenu ?
                (<ul className="space-y-2 text-center">
                    <li><a href="/" className="text-zinc-600 border-t-1 hover:text-sky-300">BOOKS</a></li>
                    <li><a href="/authors" className="text-zinc-600 hover:text-sky-300">AUTHORS</a></li>
                    <li><a href="/publishers" className="text-zinc-600 hover:text-sky-300">PUBLISHERS</a></li>
                </ul>)
            : null }

    </nav>
  )
}

export default Header