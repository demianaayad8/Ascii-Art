package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
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

	lines, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	Render(lines, input)
}

func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := string(data)

	lines := strings.Split(content, "\n")

	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}

	return lines, nil
}
func GetCharLines(lines []string, c rune) []string {
	if c < 32 || c > 126 {
		return []string{"", "", "", "", "", "", "", ""}
	}

	startLine := (int(c)-32)*9 + 1

	return lines[startLine : startLine+8]
}
func RenderLine(banner []string, text string) {
	for row := 0; row < 8; row++ {
		var builder strings.Builder

		for _, c := range text {
			charLines := GetCharLines(banner, c)
			builder.WriteString(charLines[row])
		}

		fmt.Println(builder.String())
	}
}
func Render(banner []string, input string) {
	if input == "" {
		return
	}

	parts := strings.Split(input, "\\n")

	allEmpty := true
	for _, part := range parts {
		if part != "" {
			allEmpty = false
			break
		}
	}

	if allEmpty {
		for i := 0; i < len(parts)-1; i++ {
			fmt.Println()
		}
		return
	}

	for _, part := range parts {
		if part == "" {
			fmt.Println()
			continue
		}

		RenderLine(banner, part)
	}
}
