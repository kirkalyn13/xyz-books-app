package model

// Book entity for books
type Book struct {
	ID              uint    `gorm:"column:id;type:integer;primaryKey;not null" json:"id,omitempty"`
	Title           string  `gorm:"column:title;not null;" json:"title,omitempty"`
	ISBN13          string  `gorm:"column:isbn13;not null;unique;" json:"isbn13,omitempty"`
	ISBN10          string  `gorm:"column:isbn10;not null;unique;" json:"isbn10,omitempty"`
	ListPrice       float64 `gorm:"column:list_price;not null;type:real" json:"list_price,omitempty"`
	PublicationYear int     `gorm:"column:publication_year;not null;type:integer" json:"publication_year,omitempty"`
	ImageURL        string  `gorm:"column:image_url;" json:"image_url,omitempty"`
	Edition         string  `gorm:"column:edition;" json:"edition,omitempty"`
	PublisherID     *uint   `gorm:"column:publisher_id;" json:"publisher_id,omitempty"`

	// Relationships
	Author    []*Author `gorm:"many2many:book_authors;" json:"authors,omitempty"`
	Publisher Publisher `gorm:"references:id;foreignKey:publisher_id;" json:"publisher,omitempty"`
}

func (Book) TableName() string {
	return "book"
}
