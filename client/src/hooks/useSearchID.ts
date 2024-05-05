import React from 'react'
import { useSearchParams } from 'react-router-dom'

const useSearchID = () => {
    const [ searchParams, setSearchParams ] = useSearchParams()

    const updateID = (id: string): void => {
        setSearchParams((prev: URLSearchParams) => {
            prev.set("id", id)
            return prev
          }, { replace: true })
        }

    const clearID = (): void => {
        setSearchParams((prev: URLSearchParams) => {
            prev.set("id", "")
            return prev
            }, { replace: true })
        }
    
    const getSearchID = (): string => searchParams.get("id")!

    return { updateID, clearID, getSearchID }
}

export default useSearchID