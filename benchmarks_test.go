package guitar_text_notation

import (
	"strings"
	"testing"
)

func BenchmarkScanner(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{
			name: "simple song",
			input: `{title: Simple Song}
{artist: Test Artist}

[Am]Line one of the [C]song
[F]Line two of the [G]song
`,
		},
		{
			name: "long song",
			input: func() string {
				var sb strings.Builder
				sb.WriteString("{title: Long Song}\n{artist: Test Artist}\n\n")
				// Generate 100 lines of alternating chords and lyrics
				for i := 0; i < 100; i++ {
					if i%2 == 0 {
						sb.WriteString("[Am]Line " + strings.Repeat("la ", 20) + "[C]\n")
					} else {
						sb.WriteString("Just some lyrics " + strings.Repeat("oh ", 20) + "\n")
					}
				}
				return sb.String()
			}(),
		},
		{
			name: "many directives",
			input: func() string {
				var sb strings.Builder
				for i := 0; i < 100; i++ {
					sb.WriteString("{directive" + strings.Repeat("x", i%5) + ": value}\n")
				}
				return sb.String()
			}(),
		},
		{
			name:  "blank lines",
			input: strings.Repeat("\n", 1000),
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				scanner := NewChordproScanner(strings.NewReader(bm.input))
				for line, ok := scanner.Line(); ok || line.Type == LineTypeEOF; line, ok = scanner.Line() {
					if line.Type == LineTypeEOF {
						break
					}
				}
			}
		})
	}
}

func BenchmarkIsBlank(b *testing.B) {
	inputs := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"single space", " "},
		{"long whitespace", strings.Repeat(" \t\n", 100)},
		{"non-blank", "Hello"},
		{"long non-blank", strings.Repeat("x", 1000)},
	}

	for _, input := range inputs {
		b.Run(input.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = isBlank(input.input)
			}
		})
	}
}

func BenchmarkIsDirective(b *testing.B) {
	inputs := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"simple directive", "{title: Song}"},
		{"spaced directive", "  {title: Song}  "},
		{"non-directive", "Just some text"},
		{"complex directive", "{" + strings.Repeat("nested: {value} ", 10) + "}"},
	}

	for _, input := range inputs {
		b.Run(input.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = isDirective(input.input)
			}
		})
	}
}
