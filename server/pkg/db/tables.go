package db

import "github.com/kirkalyn13/xyz-books-app/server/pkg/model"

var (
	// Tables is the list of models to be translated into database tables
	Tables = []interface{}{
		&model.Author{},
		&model.Book{},
		&model.Publisher{},
	}
)
