package guitar_text_notation

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

// LineType represents different types of lines in a ChordPro file
type LineType int

const (
	// LineTypeEOF indicates end of file has been reached
	LineTypeEOF LineType = iota
	// LineTypeBlank represents an empty line or one containing only whitespace
	LineTypeBlank
	// LineTypeDirective represents a line containing ChordPro directives (e.g., {title:})
	LineTypeDirective
	// LineTypeMusic represents a line containing musical content (chords, lyrics, or both)
	LineTypeMusic
)

// String returns the string representation of a LineType
func (lt LineType) String() string {
	switch lt {
	case LineTypeEOF:
		return "EOF"
	case LineTypeBlank:
		return "Blank"
	case LineTypeMusic:
		return "Music"
	case LineTypeDirective:
		return "Directive"
	default:
		return "Unknown"
	}
}

// ChordProLine represents a single line from a ChordPro file
type ChordProLine struct {
	Type LineType // The type of content in this line
	Text string   // The raw text content of the line
}

// ChordproScanner is a line-level scanner for chordpro lines
type ChordproScanner struct {
	scanner *bufio.Scanner // default max 64k bytes/line
}

func NewChordproScanner(r io.Reader) (sc *ChordproScanner) {
	return &ChordproScanner{
		scanner: bufio.NewScanner(r),
	}
}

// Line reads the next line from the scanner and returns it as a ChordProLine.
// Returns ok=false when end of file is reached or an error occurs.
// isBlank returns true if the string contains only whitespace
func isBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// isDirective returns true if the line is a ChordPro directive (starts with '{')
func isDirective(s string) bool {
	return strings.HasPrefix(strings.TrimSpace(s), "{")
}

func (sc *ChordproScanner) Line() (line ChordProLine, ok bool) {
	if ok = sc.scanner.Scan(); !ok {
		return ChordProLine{Type: LineTypeEOF}, false
	}

	text := sc.scanner.Text()

	// Determine line type based on content
	switch {
	case len(text) == 0 || isBlank(text):
		return ChordProLine{Type: LineTypeBlank, Text: text}, true
	case isDirective(text):
		return ChordProLine{Type: LineTypeDirective, Text: text}, true
	default:
		return ChordProLine{Type: LineTypeMusic, Text: text}, true
	}
}
