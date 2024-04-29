package model

type Publisher struct {
	ID    int    `gorm:"column:id;" json:",omitempty"`
	Name  string `gorm:"column:name;" json:",omitempty"`
	Books []Book `gorm:"column:books;" json:",omitempty"`
}
