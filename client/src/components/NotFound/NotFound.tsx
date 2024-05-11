import React from 'react'
import { BiErrorCircle } from 'react-icons/bi'

const NotFound: React.FC = () => {
    return (
        <section className="w-screen h-screen flex flex-col">
            <div className='flex justify-center mt-24'>
                <div className="w-full md:w-1/3 container text-center text-red-500 space-y-8">
                    <div className="my-4 flex justify-center align-center">
                        <BiErrorCircle className='text-red-500 text-6xl me-4' />
                        <h1 className="text-5xl">404 - Not Found</h1>
                    </div>
                    <div className='mb-8'>
                        <p className="text-3xl">Invalid ISBN 13.</p>
                    </div>
                    <div className='mt-32'>
                        <a href="/" className="text-xl underline">Return to Home</a>
                    </div>
                </div>
            </div>
        </section>
    )
}

export default NotFound