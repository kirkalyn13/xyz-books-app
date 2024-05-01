package model

// Author entity for authors
type Author struct {
	ID         uint   `gorm:"column:id;type:integer;primaryKey;autoIncrement" json:"id,omitempty"`
	FirstName  string `gorm:"column:first_name;not null;" json:"first_name,omitempty"`
	LastName   string `gorm:"column:last_name;not null;" json:"last_name,omitempty"`
	MiddleName string `gorm:"column:middle_name;" json:"middle_name,omitempty"`

	// Relationships
	Books []*Book `gorm:"many2many:author_book;" json:"books,omitempty"`
}

func (Author) TableName() string {
	return "author"
}
