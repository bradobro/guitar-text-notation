package guitar_text_notation

import (
	"fmt"
	"testing"
)

func TestGoldmark1(t *testing.T) {
	t.SkipNow()
	md1 := `
# Heading 1

This is paragraph 1. It's followed by a list:

1. List item 1
2. List item 2
1. List item 3

## Heading Level 2

That was a second level heading.
`
	html1 := TryGoldMark1(md1)
	fmt.Println(html1)
}

func TestMarkdownFile(t *testing.T) {
	// read file
}
