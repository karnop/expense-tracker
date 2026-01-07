package store

import (
	"encoding/json"
	"os"
)

// FileStore handles reading and writing expenses to a local file system
type FileStore struct {
	Filename string
}

// NewFileStore creates a new instance of FileStore with the specified filename.
func NewFileStore(filename string) *FileStore {
	return &FileStore{Filename: filename}
}

// Save writes a slice of Expense objects to the file, overwriting existing content.
func (fs *FileStore) Save(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.Filename, data, 0644)
}

// Load reads expenses from the file
// It returns an empty slice if the file does not exist.

func (fs *FileStore) Load() ([]Expense, error) {
	data, err := os.ReadFile(fs.Filename)
	if os.IsNotExist(err) {
		return []Expense{}, nil
	}

	if err != nil {
		return nil, err
	}

	var expenses []Expense
	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}
