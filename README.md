# Ascii-Art

A command-line program written in Go that turns text into ASCII art banners.

## How It Works

The program reads a string from the command line and prints it using one of three banner styles, where each character is drawn as an 8-line-tall block built from a `.txt` font file (`standard.txt`, `shadow.txt`, `thinkertoy.txt`).

## Usage

```bash
go run . "your text here" [banner]
```

- **`your text here`** — the string to render (required, must be quoted if it contains spaces)
- **`banner`** — optional banner style: `standard` (default), `shadow`, or `thinkertoy`

### Examples

```bash
# Uses the default "standard" banner
go run . "Hello"

# Uses the "shadow" banner
go run . "Hello" shadow

# Uses the "thinkertoy" banner
go run . "Hello" thinkertoy
```

### Newlines

Use `\n` inside the text to print on multiple lines:

```bash
go run . "Hello\nWorld"
```

## Project Structure

| File | Description |
|---|---|
| `main.go` | Program entry point and core logic |
| `main_test.go` | Unit tests |
| `standard.txt` | Standard banner font |
| `shadow.txt` | Shadow banner font |
| `thinkertoy.txt` | Thinkertoy banner font |
| `go.mod` | Go module definition |

## Functions

- **`main`** — parses command-line arguments, picks the banner file, and calls `Render`
- **`LoadBanner(filename string) ([]string, error)`** — reads a banner file and returns its lines
- **`GetCharLines(lines []string, c rune) []string`** — returns the 8 lines that make up a given character
- **`RenderLine(banner []string, text string)`** — prints one line of text as ASCII art, row by row
- **`Render(banner []string, input string)`** — handles `\n`-separated input and calls `RenderLine` for each part

## Requirements

- Go 1.25.6 or later

## Running Tests

```bash
go test ./...
```

## Author

[Demiana Ayad](https://github.com/demianaayad8)
