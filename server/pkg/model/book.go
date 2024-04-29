package model

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	ISBN13          string    `json:"isbn13"`
	ISBN10          string    `json:"isbn10"`
	ListPrice       float64   `json:"list_price"`
	PublicationYear int       `json:"publication_year"`
	ImageURL        string    `json:"image_url"`
	Edition         string    `json:"edition"`
	Publisher       Publisher `json:"publisher"`
}
