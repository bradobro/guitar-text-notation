package guitar_text_notation

// ChordPro Constants
const (
	ChordproCommentBegin     = '#'
	ChordproChordBegin       = '['
	ChordproChordEnd         = ']'
	ChordproDirectiveBegin   = '{'
	ChordproDirectiveEnd     = '}'
	ChordproDirectiveNameSep = ":"
	ChordproDirectiveMetaSep = " "
)

// ChordProPlus Extensions
const (
	Bar                = "|"
	RepeatBegin        = "|:"
	RepeatEnd          = ":|"
	PolyphonyOpen      = "("
	PolyphonyClose     = ")"
	RelativeNotesBegin = "." // [C .12 23 5- 43]
	FretHintSep        = ":" // [C 23.23 z- 6- -- --] string 2 fret 3 re mi tI- la
)
