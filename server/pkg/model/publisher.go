package model

type Publisher struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Books []Book `josn:"books"`
}
