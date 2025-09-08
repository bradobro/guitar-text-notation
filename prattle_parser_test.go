package guitar_text_notation

import "testing"

func TestScanner(t *testing.T) {
	// Instantiate
	scan := NewScanner()
	NotNil(t, scan)
	parse := NewParser()
	NotNil(t, parse)

	// Use
	source := "1 + 23 + 456 + 7890"
	scan.InitWithString(source)
	parse.Init(&scan).Parse(0)
}
