package guitar_text_notation

// Test the mmbros custom parser (which is loosely based on Rob Pike's parsing talk)

import (
	"os"

	"github.com/mmbros/chordpro/pkg/chordpro"
)

func MmbParseStr(src string) (songs chordpro.Songs) {
	songs = chordpro.ParseText(src)
	return
}

func MmbParseFile(fname string) (songs chordpro.Songs) {
	content, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	songs = MmbParseStr(string(content))
	return
}
