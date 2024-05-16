package util_test

import (
	"testing"

	"github.com/kirkalyn13/xyz-books-app/server/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestIsValidISBN13(t *testing.T) {
	tests := []struct {
		isbn  string
		valid bool
	}{
		{"", false},
		{"123", false},
		{"978123456789", false},
		{"97812345678922", false},
		{"9781234567897", true},
		{"9783127323207", true},
		{"9781891830020", true},
	}

	for _, test := range tests {
		if util.IsValidISBN13(test.isbn) != test.valid {
			assert.Equal(t, test.valid, test.isbn)
		}
	}
}

func TestIsValidISBN10(t *testing.T) {
	tests := []struct {
		isbn     string
		expected bool
	}{
		{"", false},
		{"123", false},
		{"160309454", false},
		{"16030945477", false},
		{"1603094547", true},
		{"160309038X", true},
		{"1891830023", true},
	}

	for _, test := range tests {
		if util.IsValidISBN10(test.isbn) != test.expected {
			assert.Equal(t, test.expected, test.isbn)
		}
	}
}
