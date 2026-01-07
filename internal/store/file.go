package store

import (
	"encoding/json"
	"fmt"
	"os"
)

// GetNextId finds the highest ID in the list and adds 1
// This prevents ID collisions when items are deleted.
func (fs *FileStore) GetNextId(expenses []Expense) int {
	maxID := 0
	for _, e := range expenses {
		if e.ID > maxID {
			maxID = e.ID
		}
	}

	return maxID + 1
}

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

// Remove deletes an expense by its ID
// returns an error if ID is not found
func (fs *FileStore) Remove(id int) error {
	expenses, err := fs.Load()
	if err != nil {
		return err
	}

	index := -1
	for i, e := range expenses {
		if e.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	// removing element from a slice
	expenses = append(expenses[:index], expenses[index+1:]...)
	return fs.Save(expenses)
}
