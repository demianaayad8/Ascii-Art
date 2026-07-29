# ASCII Art

ASCII Art is a Go command-line program that converts text into large ASCII art letters.

The program reads banner files and prints the input text using the selected ASCII art style.

## Project Description

This project takes a string as input and displays it as ASCII art.

Each supported character is stored inside a banner file.
Each character is represented using 8 lines.

The program reads the correct banner file, finds the ASCII art version of each character, and prints the final result line by line.

## Features

* Converts text into ASCII art
* Supports different banner styles
* Handles escaped new lines using `\n`
* Supports printable ASCII characters
* Rejects unsupported characters
* Handles invalid banner names
* Includes unit tests
* Uses only the Go standard library

## Supported Banners

The program supports three banner styles:

| Banner     | File             |
| ---------- | ---------------- |
| Standard   | `standard.txt`   |
| Shadow     | `shadow.txt`     |
| Thinkertoy | `thinkertoy.txt` |

If no banner is selected, the program uses the `standard` banner by default.

## Requirements

To run this project, you need Go installed on your computer.

Check your Go version:

```bash
go version
```

## How to Run

Run the program with:

```bash
go run . "Hello"
```

This will print `Hello` using the default `standard` banner.

## Choose a Banner

You can choose a banner by adding the banner name after the input.

### Standard

```bash
go run . "Hello" standard
```

### Shadow

```bash
go run . "Hello" shadow
```

### Thinkertoy

```bash
go run . "Hello" thinkertoy
```

## New Line Support

To print text on more than one line, use `\n` inside the input.

Example:

```bash
go run . "Hello\nWorld"
```

This prints `Hello` first, then prints `World` on a new ASCII art line.

You can also use multiple escaped new lines:

```bash
go run . "Hello\n\nWorld"
```

## Examples

### Example 1

Command:

```bash
go run . "Hello"
```

Output:

```text
 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/
```

### Example 2

Command:

```bash
go run . "Hello" shadow
```

This prints `Hello` using the shadow banner.

### Example 3

Command:

```bash
go run . "Hello\nWorld" thinkertoy
```

This prints two lines using the thinkertoy banner.

## Error Handling

The program handles several error cases.

### Wrong Number of Arguments

If the program is run without input or with too many arguments, it prints:

```text
Usage: go run . <string> [standard|shadow|thinkertoy]
```

### Unknown Banner

If the user enters a banner name that is not supported, the program prints an error.

Example:

```bash
go run . "Hello" random
```

Output:

```text
Error: unknown banner
```

### Unsupported Characters

The program only accepts printable ASCII characters from space to `~`.

It rejects characters such as:

* Real new line characters
* Arabic letters
* Accented letters like `é`
* Emojis
* Other unprintable characters

Example:

```bash
go run . "é"
```

Output:

```text
Error: unprintable character
```

## Testing

Run the tests with:

```bash
go test
```

Expected result:

```text
PASS
```

The tests check:

* Valid input
* Different banners
* Empty input
* Escaped new lines
* Symbols
* Numbers
* Mixed input
* Invalid characters
* Missing banner files

## Project Structure

```text
.
├── go.mod
├── main.go
├── unit_test.go
├── standard.txt
├── shadow.txt
└── thinkertoy.txt
```

## Files Explanation

### `main.go`

Contains the main program logic.

Main functions:

* `main()`
  Handles command-line arguments, banner selection, input validation, banner loading, and rendering.

* `ValidateInput()`
  Checks that the input contains only printable ASCII characters.

* `LoadBanner()`
  Reads the selected banner file and splits it into lines.

* `GetCharLines()`
  Gets the 8 ASCII art lines for one character.

* `RenderLine()`
  Prints one line of text as ASCII art.

* `Render()`
  Handles the full input, including escaped new lines.

### `unit_test.go`

Contains unit tests for the program.

### `standard.txt`

Contains the standard ASCII art banner.

### `shadow.txt`

Contains the shadow ASCII art banner.

### `thinkertoy.txt`

Contains the thinkertoy ASCII art banner.

### `go.mod`

Defines the Go module.

## Author

Created by Demiana Ayad.
