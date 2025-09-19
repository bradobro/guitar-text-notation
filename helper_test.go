package guitar_text_notation

import "testing"

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single space", " ", true},
		{"multiple spaces", "   ", true},
		{"tabs and spaces", "\t  \t", true},
		{"newlines", "\n\r\n", true},
		{"mixed whitespace", " \t\n\r  ", true},
		{"single character", "x", false},
		{"text with spaces", "  hello  ", false},
		{"special characters", "\u2003\u2002", true}, // em space and en space
		{"zero-width space", "\u200b", true},
		{"non-breaking space", "\u00A0", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBlank(tt.input); got != tt.expected {
				t.Errorf("isBlank(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsDirective(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"just brace", "{", true},
		{"simple directive", "{title: Song}", true},
		{"directive with spaces", "  {artist: Name}  ", true},
		{"nested braces", "{chord: [Am]}", true},
		{"text with braces", "Hello {world}", false},
		{"commented directive", "# {title: Song}", false},
		{"brace in middle", "Some {text} here", false},
		{"escaped brace", "\\{not a directive}", false},
		{"multiple directives", "{title: Song}{key: C}", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDirective(tt.input); got != tt.expected {
				t.Errorf("isDirective(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
