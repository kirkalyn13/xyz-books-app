package model

// Publisher entity for publishers
type Publisher struct {
	ID   uint   `gorm:"column:id;type:integer;primaryKey;not null" json:"id,omitempty"`
	Name string `gorm:"column:name;unique;not null;" json:"name,omitempty"`

	// Relationships
	Books []*Book `gorm:"references:id;foreignKey:publisher_id;" json:"books,omitempty"`
}

func (Publisher) TableName() string {
	return "publisher"
}
