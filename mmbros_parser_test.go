package guitar_text_notation

import (
	"fmt"
	"testing"
)

func TestMmbParser(t *testing.T) {
	songs := MmbParseFile("./_testdata/Better Word.cho")
	fmt.Print(songs.String())
}
