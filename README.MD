# InvoiceTrack

**InvoiceTrack** is a CLI tool designed to process and track invoice balances from daily statement files (e.g., XML or JSON). It ensures accurate aggregation and analysis of billing data by leveraging Go’s concurrency model.

## 🧠 Overview

InvoiceTrack scans a directory for files with a given extension (e.g., `.xml`), extracts invoice data, and computes the total balances.
### Key Features


- 🗂 Process multiple files concurrently.
- 🔄 Built with **pipelines** and **semaphores** to control concurrency.
- 📄 Supports different file formats (`xml`, `json`, etc.).
- 🛠 Extensible design for additional processing logic.

---

## Run Project

This project provides a CLI tool to process and aggregate invoice balances from statement files using concurrent patterns like **pipelines** and **semaphores**.

### 📦 Available Command

- `compute-invoice-balances`: Reads all statement files from a directory (e.g., XML or JSON) and calculates the total amounts based on specific filtering criteria.

### Build: 

```bash
go build -tags invoice-track -o ./build/invoice-track cmd/main.go
```

### ▶️ Run the CLI

#### 🧩 Optional Flags

This command supports the following flags:

- `--dir`: Path to the folder containing the input files.  
  **Default**: `${HOME}/Documents/xml_repo`

- `--ext`: File extension to filter (e.g., `xml`, `json`).  
  **Default**: `xml`


*Run with default values:* 

```bash
./build/invoice-track compute-invoice-balances
```

```bash
go run ./cmd/main.go compute-invoice-balances
```

*Run with values:* 

```bash
./build/invoice-track  compute-invoice-balances --dir <path-to-files> --ext <file-extension>
```

```bash
go run ./cmd/main.go  compute-invoice-balances --dir <path-to-files> --ext <file-extension>
```

## ℹ️ Help

```bash
./build/invoice-track  compute-invoice-balances --help
```

```bash
go run ./cmd/main.go compute-invoice-balances --help
```

### Run Benchmark:

```bash
go test -bench ^BenchmarkTotalAmountCalculator$ -benchmem -v ./...
```

## ⚙️ Architecture

This tool uses Go’s native concurrency primitives:

- **Pipelines**: to decouple stages like file discovery, parsing, and aggregation.
- **Semaphores**: to limit the number of files being processed concurrently, improving performance and stability.

## 🎨 Processing Flow
```bash
                    +--------------+
                    |   Reader     |
                    | Reads files  |
                    +--------------+
                            |
                     (chan string)
                            v
                    +--------------+
                    |    Parser    |
                    |   Parse XML  |
                    +--------------+
                            |
                (chan InvoiceStatement)
                            v
                +------------------------+
                |    FilterSettlements   |
                |    Only num_liq = "0"  |
                +------------------------+
                            |
                  (chan FilteredInvoice)
                            v
            +-------------------------------+
            |  SumTotalAmountByFUECDAndType |
            |  Group by FUECD & type        |
            +-------------------------------+
                            |
                        Output
                (Totals per FUECD/type)
```

## 📚 References

This section lists the sources, documentation, and tools that were referenced or used in the development of this project:

1. [Pipelines](https://go.dev/blog/pipelines)
2. [Semaphore](https://www.codingexplorations.com/blog/understanding-and-implementing-the-semaphore-pattern-in-go)