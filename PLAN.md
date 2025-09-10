# Plan Notes

A few libs:

- [prattle](https://github.com/askeladdk/prattle). **PROBABLY NOT**: light and understandable, but geared toward expressions with operator precedence. creates a fairly traditional PRAT scanner and parser, with [docs](https://pkg.go.dev/github.com/askeladdk/prattle#) in the godoc site. Google AI says, "Primarily designed for parsing expressions with operator precedence and associativity rules (like arithmetic expressions or function calls)." This aligns with my investigation.
- [re2c](https://github.com/skvadrik/re2c) code generator
- [ragel](https://github.com/adrian-thurston/ragel) a code generator
- [participle](https://github.com/alecthomas/participle) uses tagged structs
- [peg](https://github.com/pointlander/peg/tree/main) a PEG parser
- [goldmark](https://github.com/yuin/goldmark) Markdown parser
- [goparsec](https://github.com/prataprc/goparsec) a parser combinator. Google AI says this can be less performant for large inputs and it can struggle with operator precedence (see prattle for that).
- [Nearly](https://github.com/kach/nearley) is a JavaScript BNF compiler with some great docs and theory. It can compile to TypeScript too.

With all of these, the hardest thing is figuring out the docs. Good examples are a bonus.

- testing without testify
  - https://antonz.org/do-not-testify/
  - https://www.alexedwards.net/blog/the-9-go-test-assertions-i-use
