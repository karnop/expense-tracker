package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Defining the add subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// defining flags for add subcommand
	addCategory := addCmd.String("category", "", "The category of the expense.")
	addAmount := addCmd.Float64("amount", 0, "The amount of the expense")

	// checking if the user provides enough arguments
	if len(os.Args) < 2 {
		fmt.Println("expected 'add' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add" :
		addCmd.Parse(os.Args[2:])

		// validating input
		if *addCategory == "" || * addAmount == 0 {
			fmt.Println("Usage: tracker add --category <name> --amount <value>")
			addCmd.PrintDefaults()
			os.Exit(1)
		}

		// temporary : printing output to verify working
		fmt.Printf("Adding Expense: Category=%s, Amount=%.2f\n", *addCategory, *addAmount)

	default:
		fmt.Println("expected 'add' subcommand")
		os.Exit(1)
	}
}