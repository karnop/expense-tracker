# Go CLI Expense Tracker

A lightweight, command-line interface (CLI) tool for tracking personal expenses. Built with Go, this project demonstrates idiomatic Go patterns, a standard project layout, and robust error handling.

---

## 🚀 Features

- **Add Expenses**: Record transactions with a category and amount.
- **Persistence**: Automatically saves data to a local JSON file (`expenses.json`), ensuring data survives between runs.
- **List Expenses**: Displays a formatted, tab-aligned table of all recorded expenses.
- **Delete Expenses**: Remove entries by ID with collision-safe ID generation.
- **Data Safety**: Uses atomic-like read/write operations and handles edge cases (for example, deleting an ID does not corrupt future IDs).

---

## 🛠️ Project Structure

This project follows the **Standard Go Project Layout**:

```
.
├── cmd/
│   └── tracker/
│       └── main.go
├── internal/
│   └── store/
│       └── types.go
|       └── file.go
|       └── file_test.go
├── expenses.json   # Auto-generated data file
└── README.md
```

- `cmd/tracker/`: Application entry point.
- `internal/store/`: Core logic for data handling and file I/O. Using an internal package enforces encapsulation.
- `expenses.json`: Local data store (created automatically).

---

## 📦 Installation & Usage

### Prerequisites

- Go **1.21** or higher

### 1. Clone and Build

```bash
git clone https://github.com/karnop/expense-tracker.git
cd expense-tracker

# Build the binary
go build -o tracker cmd/tracker/main.go
```

### 2. Run Commands

#### Add an Expense

```bash
./tracker add --category "Coffee" --amount 4.50
```

_Output:_

```
Expense added successfully (ID: 1)
```

#### List Expenses

```bash
./tracker list
```

_Output:_

```
ID  Category  Amount
1   Coffee    4.50
```

#### Delete an Expense

```bash
./tracker delete --id 1
```

_Output:_

```
Expense with ID 1 deleted successfully
```

> **Tip:** You can also run commands without building the binary:
>
> ```bash
> go run cmd/tracker/main.go <command>
> ```

---

## 🧪 Running Tests

This project includes unit tests for business logic and integration tests for file system operations.

```bash
go test ./... -v
```

---

## 🧠 Key Technical Concepts Used

- **Dependency Injection**: `FileStore` is injected into handlers, making the logic easy to test.
- **Table-Driven Tests**: Used for comprehensive testing of ID generation logic.
- **Struct Tags**: Enable clean JSON serialization and deserialization.
- **`tabwriter`**: Formats console output into aligned columns.

---
