package model

// Response types
type (
	// BooksResponse books response
	BooksResponse struct {
		Books []Book `json:"books,omitempty"`
	}

	// AuthorsResponse authors response
	AuthorsResponse struct {
		Authors []Author `json:"authors,omitempty"`
	}

	// PublishersResponse publishers response
	PublishersResponse struct {
		Publishers []Publisher `json:"publishers,omitempty"`
	}

	// BookResponse book response
	BookResponse struct {
		Book Book `json:"book,omitempty"`
	}

	// AuthorResponse author response
	AuthorResponse struct {
		Author Author `json:"author,omitempty"`
	}

	// PublisherResponse publisher response
	PublisherResponse struct {
		Publisher Publisher `json:"publisher,omitempty"`
	}
)
