import { Author } from "./author"

export type Book = {
    id?: number
    title: string
    isbn13: string
    isbn10: string
    listPrice: number
    publicationYear: number
    imageURL: string
    edition: string
    publisherID: number
    authors: Author[]
}