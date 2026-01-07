package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/karnop/expense-tracker/internal/store"
)

func main() {
	// Initializing the FileStore
	// We will store data in a file named "expenses.json" in the current directory
	storage := store.NewFileStore("expenses.json")

	// checking if the user provides enough arguments
	if len(os.Args) < 2 {
		fmt.Println("expected 'add' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		runAdd(storage)

	case "list":
		runList(storage)

	case "delete":
		runDelete(storage)

	default:
		fmt.Println("expected 'add' subcommand")
		os.Exit(1)
	}
}

func runAdd(storage *store.FileStore) {
	// Defining the add subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// defining flags for add subcommand
	addCategory := addCmd.String("category", "", "The category of the expense.")
	addAmount := addCmd.Float64("amount", 0, "The amount of the expense")

	addCmd.Parse(os.Args[2:])

	// validating input
	if *addCategory == "" || *addAmount == 0 {
		fmt.Println("Usage: tracker add --category <name> --amount <value>")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	// Loading existing expenses to prevent overwriting
	expenses, err := storage.Load()
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
		os.Exit(1)
	}

	// creating the new expense object
	newExpense := store.Expense{
		ID:        storage.GetNextId(expenses),
		Category:  *addCategory,
		Amount:    *addAmount,
		CreatedAt: time.Now(),
	}

	// appending the new expense to the list
	expenses = append(expenses, newExpense)

	// saving the updated list back to the file
	if err := storage.Save(expenses); err != nil {
		fmt.Printf("Error saving expense: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", newExpense.ID)
}

func runList(storage *store.FileStore) {
	expenses, err := storage.Load()
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
		os.Exit(1)
	}

	// Creating a new tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// printing the header
	fmt.Fprintln(w, "ID\tCategory\tAmount")

	for _, e := range expenses {
		fmt.Fprintf(w, "%d\t%s\t%.2f\n", e.ID, e.Category, e.Amount)
	}

	// flushing the writer to ensure everything is printed
	w.Flush()
}

func runDelete(storage *store.FileStore) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("id", 0, "The Id of the expense to delete")

	deleteCmd.Parse(os.Args[2:])

	if *deleteID == 0 {
		fmt.Println("Usage: tracker delete --id <value>")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}

	if err := storage.Remove(*deleteID); err != nil {
		fmt.Printf("Error deleting expense: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Expense with ID %d deleted successfully\n", *deleteID)
}
