package model

type Author struct {
	ID         int    `gorm:"column:id;" json:",omitempty"`
	FirstName  string `gorm:"column:first_name;" json:",omitempty"`
	LastName   string `gorm:"column:last_name;" json:",omitempty"`
	MiddleName string `gorm:"column:middle_name;" json:",omitempty"`
	Books      []Book `gorm:"column:books;" json:",omitempty"`
}
