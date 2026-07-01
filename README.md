# calc — CLI Calculator

> A command-line calculator written in Go. Learning Project #1 as part of preparation for a Go Backend Developer role.

---

## For the Recruiter

### What It Is and Why

The first project in the Go transition roadmap. The goal is not the calculator itself, but **mastering the core constructs of the language** in isolation from frameworks and libraries: working with `os.Args`, static typing, error handling via `stderr` + exit codes, and basic control flow.

### What This Project Demonstrates

| Skill | Implementation |
|---|---|
| Handling CLI Arguments | `os.Args` without third-party parsers |
| Static Typing | `float64` for all numeric operations |
| Error Handling | `fmt.Fprintf(os.Stderr)` + `os.Exit(1)` |
| Control Flow | `switch` statement with an explicit `default` case for unknown operators |
| Standard Library | Restricted to `fmt`, `os`, and `strconv` |

### Stack

- **Language:** Go 1.22+
- **Dependencies:** Standard library only
- **Platform:** Linux / macOS / Windows

---

## For the Developer

### Architectural Decisions

**Why keep everything in a single `main.go`?**
The project is intentionally minimalistic. Splitting 50 lines of code across packages would be over-engineering. For a project of this scale, a single file is the correct decision, not technical debt.

**Why `float64` instead of `int`?**
`10 / 3` should evaluate to `3.3333...`, not `3`. Enforcing `float64` via `strconv.ParseFloat` handles this elegantly without adding extra logic.

**Why output errors to `stderr` instead of `stdout`?**
This is a UNIX convention: normal output goes to `stdout`, while errors go to `stderr`. This allows the calculator to be used seamlessly in shell pipes and scripts:
```bash
result=$(./calc 10 + 5)   # $result == "15", errors won't pollute the variable

```

**Using Exit Code 1 on error** is standard practice for CLI tools. It allows the invoking script to check for success using `$?`.

### Structure

```
calc/
└── main.go   # entry point + all logic

```

### Installation and Setup

```bash
git clone [https://github.com/Shipovmax/calc](https://github.com/Shipovmax/calc)
cd calc
go build -o calc .
./calc 10 + 5

```

### Usage

```bash
./calc <number> <operator> <number>

```

**Supported Operators:** `+`, `-`, `*`, `/`

**Numbers:** Any values parseable as `float64` (integers, decimals, negative numbers)

### Examples

```bash
./calc 10 + 5      # 15
./calc 20 - 3      # 17
./calc 6 * 7       # 42
./calc 10 / 3      # 3.3333333333333335
./calc 2.5 + 1.5   # 4
./calc -5 + 3      # -2

```

### Error Handling

```bash
./calc 10 / 0
# stderr: error: division by zero
# exit code: 1

./calc 10 ^ 2
# stderr: error: unknown operator "^"
# exit code: 1

./calc 10 5
# stderr: error: exactly 3 arguments are required (received: 2)
# exit code: 1

./calc abc + 5
# stderr: error: "abc" is not a number
# exit code: 1

```

### Running Without Building

```bash
go run main.go 10 + 5

```

