package db

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDatabase(t *testing.T) {
	err := os.Remove(filepath.Join(".", "gorm.db"))
	assert.NoError(t, err)

	err = LoadDatabase()
	assert.NoError(t, err)
}
