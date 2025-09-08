package guitar_text_notation

/* SPIKE: evaluate the prattle PRAT parser.

I think it could do the job at the parser level. I'd have to build my own
driver to handle each of the tokens, and it's not clear what the precedence does
in the context of parsing song lines.

*/

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/askeladdk/prattle"
)

func NewScanner() prattle.Scanner {
	scanner := prattle.Scanner{
		Scan: func(s *prattle.Scanner) (kind int) {
			s.ExpectAny(unicode.IsSpace)
			s.Skip()

			// Scan the next token.
			switch {
			case s.Done(): // Stop when the entire input has been consumed.
				return 0
			case s.Expect('+'): // Scan the addition operator.
				return 1
			case s.ExpectOne(unicode.IsDigit): // Scan a number consisting of one or more digits.
				s.ExpectAny(unicode.IsDigit)
				return 2
			}

			// Invalid token.
			s.Advance()
			return -1
		},
	}
	return scanner
}

// Define the parsing Driver.
type driver struct {
	stack []int
}

func (d *driver) push(i int) {
	d.stack = append(d.stack, i)
}

func (d *driver) pop() (i int) {
	n := len(d.stack)
	i, d.stack = d.stack[n-1], d.stack[:n-1]
	return
}

func (d *driver) number(p *prattle.Parser, t prattle.Token) error {
	n, _ := strconv.Atoi(t.Text)
	d.push(n)
	return nil
}

func (d *driver) add(p *prattle.Parser, t prattle.Token) error {
	// Parse the right hand operator.
	_ = p.Parse(d.Precedence(t.Kind))

	right := d.pop()
	left := d.pop()
	sum := left + right
	fmt.Printf("%d + %d = %d\n", left, right, sum)
	d.push(sum)
	return nil
}

func (d *driver) Prefix(kind int) prattle.ParseFunc {
	if kind == 2 {
		return d.number
	}
	return nil
}

func (d *driver) Infix(kind int) prattle.ParseFunc {
	if kind == 1 {
		return d.add
	}
	return nil
}

// Precedence makes numbers (token 2) bind more tightly than
// the plus operator
func (d *driver) Precedence(kind int) int {
	return kind
}

func (d *driver) ParseError(t prattle.Token) error {
	return fmt.Errorf("%s", t)
}

func NewParser() prattle.Parser {
	return prattle.Parser{
		Driver: &driver{},
	}
}
