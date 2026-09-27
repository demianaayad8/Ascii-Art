package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: go run . <string> [standard|shadow|thinkertoy]")
		return
	}

	input := os.Args[1]
	bannerFile := "standard.txt"

	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "standard":
			bannerFile = "standard.txt"
		case "shadow":
			bannerFile = "shadow.txt"
		case "thinkertoy":
			bannerFile = "thinkertoy.txt"
		default:
			fmt.Println("Error: unknown banner")
			return
		}
	}

	err := ValidateInput(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = Render(lines, input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func ValidateInput(input string) error {
	for _, c := range input {
		if c < 32 || c > 126 {
			return fmt.Errorf("unprintable character")
		}
	}
	return nil
}

func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	if len(lines) < 855 {
		return nil, fmt.Errorf("invalid banner file: not enough lines")
	}

	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}

	return lines, nil
}

func GetCharLines(lines []string, c rune) ([]string, error) {
	startLine := (int(c)-32)*9 + 1

	return lines[startLine : startLine+8], nil
}

func RenderLine(banner []string, text string) error {
	var builder strings.Builder

	for row := 0; row < 8; row++ {
		builder.Reset()

		for _, c := range text {
			charLines, err := GetCharLines(banner, c)
			if err != nil {
				return err
			}

			builder.WriteString(charLines[row])
		}

		fmt.Println(builder.String())
	}

	return nil
}

func Render(banner []string, input string) error {
	if input == "" {
		return nil
	}

	parts := strings.Split(input, `\n`)

	if strings.ReplaceAll(input, "\\n", "") == "" {
		parts = parts[:len(parts)-1]
	}

	for _, part := range parts {
		if part == "" {
			fmt.Println()
		} else if err := RenderLine(banner, part); err != nil {
			return err
		}
	}

	return nil
}