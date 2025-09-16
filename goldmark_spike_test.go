package guitar_text_notation

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/yassinebenaid/godump"
)

const BETTER_WORD = "_testdata/better-word.md"

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
	html1 := MdToHtml(md1)
	fmt.Println(html1)
}

func TestMarkdownFile(t *testing.T) {
	t.SkipNow()
	html1 := MdFileToHtml(BETTER_WORD)
	fmt.Println(html1)
}

func TestSpikeAST(t *testing.T) {
	t.Skip("not reeady yet")
	simple := false
	ast := MdFileToAST(BETTER_WORD)
	if simple {
		godump.Dump(ast)
	} else {
		spew.Dump(ast)
	}
}
