package store

import (
	"testing"
)

// Unit test
func TestGetNextID(t *testing.T) {
	// this struct represents a single test case
	tests := []struct {
		name     string    // Name of the test case (for logging)
		input    []Expense // The input we give to the function
		expected int       // The result we expect to get back
	}{
		{
			name:     "empty list should return 1",
			input:    []Expense{},
			expected: 1,
		},
		{
			name: "list with IDs 1 and 2 should return 3",
			input: []Expense{
				{ID: 1},
				{ID: 2},
			},
			expected: 3,
		},
		{
			name: "list with gap (IDs 1 and 5) should return 6",
			input: []Expense{
				{ID: 1},
				{ID: 5}, // Deleted 2, 3, 4
			},
			expected: 6,
		},
	}

	// Execution loop
	// verifying each testcase in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// creating a filestore instance to call the method
			fs := FileStore{}

			// calling the function
			got := fs.GetNextId(tt.input)

			// Assertion
			if got != tt.expected {
				t.Errorf("GetNextID() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Integration test
func TestFileStore_SaveAndLoad(t *testing.T) {
	// creating a temp directory
	dir := t.TempDir()
	storeFile := dir + "/test_expenses.json"
	fs := NewFileStore(storeFile)

	// defining data to save
	expenses := []Expense{
		{ID: 1, Category: "Test", Amount: 100},
	}

	// saving data
	err := fs.Save(expenses)
	if err != nil {
		t.Fatalf("Save() failed: %v", err) 
		// Fatalf stops the test immediately
	}

	// Loading it back
	loadedExpenses, err := fs.Load()
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	// verifying data integrity
	if len(loadedExpenses) != 1 {
		t.Errorf("Expected 1 expense, got %d", len(loadedExpenses))
	}

	if loadedExpenses[0].Category != "Test" {
		t.Errorf("Expected category 'Test', got '%s'", loadedExpenses[0].Category)
	}
}