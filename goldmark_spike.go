package guitar_text_notation

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func TryGoldMark1(source string) (output string) {

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(source), &buf); err != nil {
		panic(err)
	}
	output = buf.String()
	return
}
