package store

import (
	"time"
)

// Expense represents a single financial transaction record.
type Expense struct {
	ID        int       `json:"id"`
	Category  string    `json:"category"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
