# Installing Go (Golang) on Windows and Mac

## 📥 Installation Overview

This guide walks you through installing **Go** (Golang) on:

- Windows
- macOS (Intel & Apple Silicon)

---

## 🖥️ Install Go on Windows

### 1. Download the Go Installer

- Visit the official Go download page:  
  👉 [https://go.dev/dl/](https://go.dev/dl/)

- Download the **Windows Installer** (`.msi` file) for your system (64-bit most likely).

### 2. Run the Installer

- Open the downloaded `.msi` file.
- Follow the installation steps:
  - Accept the license agreement
  - Choose the destination folder (default is fine)
  - Complete the installation.

### 3. Verify Installation

- Open **Command Prompt** (press `Win + R`, type `cmd`, hit Enter).
- Type the following command:

  ```bash
  go version
  ```

---

# Hello World in Go

This repository contains a basic "Hello World" program written in the Go programming language.

## Code Explanation

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

### Breakdown:

- `package main`:  
  Defines the package name. In Go, the `main` package is special — it tells the Go compiler that this is an executable program, not a shared library.

- `import "fmt"`:  
  Imports the `fmt` package, which contains functions for formatting text, including printing to the console.

- `func main() { ... }`:  
  Declares the `main` function — the entry point of a Go program. When you run the program, the code inside `main` gets executed.

- `fmt.Println("Hello World")`:  
  Prints the text `Hello World` to the standard output (your terminal) followed by a newline.

## How to Run

1. Make sure you have [Go installed](https://golang.org/dl/).
2. Save the code in a file named `main.go`.
3. Open a terminal and navigate to the folder containing `main.go`.
4. Run the following command:

   ```bash
   go run main.go
   ```

5. You should see the output:

   ```
   Hello World
   ```
