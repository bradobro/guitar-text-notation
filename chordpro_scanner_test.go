package guitar_text_notation

import (
	"strings"
	"testing"
)

func TestChordProScanner(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []ChordProLine
	}{
		{
			name:  "empty input",
			input: "",
			expected: []ChordProLine{
				{Type: LineTypeEOF},
			},
		},
		{
			name:  "single blank line",
			input: "\n",
			expected: []ChordProLine{
				{Type: LineTypeBlank, Text: ""},
				{Type: LineTypeEOF},
			},
		},
		{
			name:  "whitespace line",
			input: "   \t  \n",
			expected: []ChordProLine{
				{Type: LineTypeBlank, Text: "   \t  "},
				{Type: LineTypeEOF},
			},
		},
		{
			name:  "directive line",
			input: "{title: My Song}\n",
			expected: []ChordProLine{
				{Type: LineTypeDirective, Text: "{title: My Song}"},
				{Type: LineTypeEOF},
			},
		},
		{
			name:  "directive with leading space",
			input: "  {key: C}\n",
			expected: []ChordProLine{
				{Type: LineTypeDirective, Text: "  {key: C}"},
				{Type: LineTypeEOF},
			},
		},
		{
			name:  "music line with chords",
			input: "[Am] [C] [F]\n",
			expected: []ChordProLine{
				{Type: LineTypeMusic, Text: "[Am] [C] [F]"},
				{Type: LineTypeEOF},
			},
		},
		{
			name:  "music line with lyrics",
			input: "Hello darkness my old friend\n",
			expected: []ChordProLine{
				{Type: LineTypeMusic, Text: "Hello darkness my old friend"},
				{Type: LineTypeEOF},
			},
		},
		{
			name:  "music line with chords and lyrics",
			input: "[Am]Hello [C]darkness my old [F]friend\n",
			expected: []ChordProLine{
				{Type: LineTypeMusic, Text: "[Am]Hello [C]darkness my old [F]friend"},
				{Type: LineTypeEOF},
			},
		},
		{
			name: "multiple line types",
			input: `{title: The Sound of Silence}
{artist: Simon & Garfunkel}

[Am]Hello darkness my old [G]friend
I've come to talk with you [Am]again
`,
			expected: []ChordProLine{
				{Type: LineTypeDirective, Text: "{title: The Sound of Silence}"},
				{Type: LineTypeDirective, Text: "{artist: Simon & Garfunkel}"},
				{Type: LineTypeBlank, Text: ""},
				{Type: LineTypeMusic, Text: "[Am]Hello darkness my old [G]friend"},
				{Type: LineTypeMusic, Text: "I've come to talk with you [Am]again"},
				{Type: LineTypeBlank, Text: ""},
				{Type: LineTypeEOF},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewChordproScanner(strings.NewReader(tt.input))
			var got []ChordProLine

			// Collect all lines from the scanner
			for line, ok := scanner.Line(); ok || line.Type == LineTypeEOF; line, ok = scanner.Line() {
				got = append(got, line)
				if line.Type == LineTypeEOF {
					break
				}
			}

			// Compare results
			if len(got) != len(tt.expected) {
				t.Errorf("got %d lines, expected %d lines", len(got), len(tt.expected))
				return
			}

			for i, line := range got {
				if line.Type != tt.expected[i].Type {
					t.Errorf("line %d: got type %v, expected %v", i, line.Type, tt.expected[i].Type)
				}
				if line.Text != tt.expected[i].Text {
					t.Errorf("line %d: got text %q, expected %q", i, line.Text, tt.expected[i].Text)
				}
			}
		})
	}
}

// Test edge cases and error handling
func TestChordProScannerEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "very long line",
			input: strings.Repeat("[Am]La ", 1000) + "\n", // Should handle lines up to bufio.MaxScanTokenSize
		},
		{
			name:  "no final newline",
			input: "{title: Song}[Am]Hello there",
		},
		{
			name:  "mixed line endings",
			input: "Line 1\rLine 2\r\nLine 3\nLine 4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewChordproScanner(strings.NewReader(tt.input))
			for line, ok := scanner.Line(); ok || line.Type == LineTypeEOF; line, ok = scanner.Line() {
				// Just verify we can read all lines without panicking
				if line.Type == LineTypeEOF {
					break
				}
			}
		})
	}
}
