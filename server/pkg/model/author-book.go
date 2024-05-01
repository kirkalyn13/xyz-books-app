package model

// AuthorBook many-to-many relationships of Authors and Books
type AuthorBook struct {
	AuthorID uint
	BookID   uint
}

func (AuthorBook) TableName() string {
	return "author_book"
}
