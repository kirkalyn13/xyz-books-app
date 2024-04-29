package model

type Book struct {
	ID              int       `gorm:"column:id;" json:",omitempty"`
	Title           string    `gorm:"column:title;" json:",omitempty"`
	ISBN13          string    `gorm:"column:isbn13;" json:",omitempty"`
	ISBN10          string    `gorm:"column:isbn10;" json:",omitempty"`
	ListPrice       float64   `gorm:"column:list_price;" json:",omitempty"`
	PublicationYear int       `gorm:"column:publication_year;" json:",omitempty"`
	ImageURL        string    `gorm:"column:image_url;" json:",omitempty"`
	Edition         string    `gorm:"column:edition;" json:",omitempty"`
	Publisher       Publisher `gorm:"column:publisher;" json:",omitempty"`
}
