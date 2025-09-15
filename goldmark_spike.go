package guitar_text_notation

import (
	"bytes"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func MdToHtml(source string) (output string) {

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(source), &buf); err != nil {
		panic(err)
	}
	output = buf.String()
	return
}

func MdFileToHtml(fname string) (output string) {
	content, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		panic(err)
	}
	output = buf.String()
	return
}

func MdFileToAST(fname string) ast.Node {
	content, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	// Create a parser with our custom extension
	md := goldmark.New(
		goldmark.WithExtensions(
			NewTextExtension(),
		),
	)

	// Create a reader with the content
	reader := text.NewReader(content)

	// Parse the content into an AST
	doc := md.Parser().Parse(reader)
	return doc
}
