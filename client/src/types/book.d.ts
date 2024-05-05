import { Author } from "./author"

export type Book = {
    id?: string
    title: string
    isbn13: string
    isbn10: string
    list_price: number
    publication_year: number
    image_url: string
    edition: string
    publisher_id: number
    authors: Author[]
}