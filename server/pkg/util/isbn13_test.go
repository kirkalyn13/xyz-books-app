package util_test

import (
	"testing"

	"github.com/kirkalyn13/xyz-books-app/pkg/util"
)

func TestIsValidISBN13(t *testing.T) {
	tests := []struct {
		isbn  string
		valid bool
	}{
		{"", false},
		{"123", false},
		{"978123456789", false},
		{"9781234567897", true},
		{"9783127323207", true},
	}

	for _, test := range tests {
		if util.IsValidISBN13(test.isbn) != test.valid {
			t.Errorf("Test Failed: ISBN-13 %s validation failed", test.isbn)
		}
	}
}
