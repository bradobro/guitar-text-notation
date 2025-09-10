package guitar_text_notation

import "testing"

func TestScanner(t *testing.T) {
	t.SkipNow()
	// Instantiate
	scan := NewScanner()
	NotNil(t, scan)
	parse := NewParser()
	NotNil(t, parse)

	// Use
	source := "1 + 23 + 456 + 7890"
	scan.InitWithString(source)
	// from the lib docs: Parse parses using the TDOP algorithm until it encounters a token with an equal or lower precedence than least. It may be called in a mutual recursive manner by the parsing functions provided by the Driver.
	parse.Init(&scan).Parse(0)
}
