package guitar_text_notation

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// TextLineNode represents our custom node type for special text lines
type TextLineNode struct {
	ast.BaseBlock
	// Add any custom fields you need here
}

// Dump implements Node.Dump .
func (n *TextLineNode) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// KindTextLine is a unique kind ID for our custom node
var KindTextLine = ast.NewNodeKind("TextLine")

// Kind implements Node.Kind.
func (n *TextLineNode) Kind() ast.NodeKind {
	return KindTextLine
}

// textLineParser is our custom parser for special text lines
type textLineParser struct {
}

func (s *textLineParser) Trigger() []byte {
	// Return the character(s) that trigger this parser
	// For example, return []byte{'-'} to trigger on lines starting with -
	return []byte{'-'}
}

// Parse parses the current line and returns a new node if it matches our criteria
func (s *textLineParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, _ := block.PeekLine()
	// Add your custom parsing logic here
	// For example, check if the line matches a certain pattern
	if len(line) > 0 && line[0] == '-' {
		node := &TextLineNode{}
		// Add the text as a child node
		pos, _ := block.Position()
		text := ast.NewTextSegment(text.NewSegment(pos, pos+len(line)))
		node.AppendChild(node, text)
		block.Advance(len(line))
		return node
	}
	return nil
}

// textExtension is our custom extension
type textExtension struct {
}

// New returns a new TextExtension
func NewTextExtension() goldmark.Extender {
	return &textExtension{}
}

// Extend implements goldmark.Extender.
func (e *textExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(&textLineParser{}, 100),
		),
	)
}