# Plan Notes

A few libs:

- [prattle](https://github.com/askeladdk/prattle) creates a fairly traditional PRAT scanner and parser, with [docs](https://pkg.go.dev/github.com/askeladdk/prattle#) in the godoc site. Looks like it's based on parsing operators with precedence
- [re2c](https://github.com/skvadrik/re2c) code generator
- [ragel](https://github.com/adrian-thurston/ragel) a code generator
- [participle](https://github.com/alecthomas/participle) uses tagged structs
- [peg](https://github.com/pointlander/peg/tree/main) a PEG parser
- [goldmark](https://github.com/yuin/goldmark) Markdown parser
- [goparsec](https://github.com/prataprc/goparsec) a parser combinator

With all of these, the hardest thing is figuring out the docs. Good examples are a bonus.

- testing without testify
  - https://antonz.org/do-not-testify/
  - https://www.alexedwards.net/blog/the-9-go-test-assertions-i-use
