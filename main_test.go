package main

import "testing"

func TestLoadBanner(t *testing.T) {
	lines, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(lines) != 856 {
		t.Fatalf("expected 856 lines, got %d", len(lines))
	}
}

func TestGetCharLines(t *testing.T) {
	lines, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	charLines := GetCharLines(lines, 'A')

	if len(charLines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(charLines))
	}
}

func TestGetCharLinesBang(t *testing.T) {
	lines, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	charLines := GetCharLines(lines, '!')

	if len(charLines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(charLines))
	}
}

func TestGetCharLinesSpace(t *testing.T) {
	lines, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	charLines := GetCharLines(lines, ' ')

	if len(charLines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(charLines))
	}
}

func TestGetCharLinesInvalidRune(t *testing.T) {
	lines, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	charLines := GetCharLines(lines, '😊')

	if len(charLines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(charLines))
	}

	for _, line := range charLines {
		if line != "" {
			t.Fatalf("expected empty line for invalid rune, got %q", line)
		}
	}
}

func TestRenderEmpty(t *testing.T) {
	lines, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	Render(lines, "")
}
