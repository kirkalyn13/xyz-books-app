package db

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDatabase(t *testing.T) {
	gormFile := filepath.Join(".", "gorm.db")

	if _, err := os.Stat(gormFile); os.IsNotExist(err) {
		os.Remove(gormFile)
	}

	err := LoadDatabase()
	assert.NoError(t, err)
}
