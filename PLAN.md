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
- This [Go ChordPro Parser](https://github.com/mmbros/chordpro) uses Rob Pike's parsing methodology. [Muckbucket's Processor](https://github.com/muckbucket/chordpro) may also be instructive

- [LALR Theory](https://web.archive.org/web/20210507215636/http://web.cs.dal.ca:80/~sjackson/lalr1.html)
- Other work (search GitHub for *guitar notation*)
    - [alphaTab](https://github.com/CoderLine/alphaTab) another take on horizontally describing guitar tab, and [the docs page](https://www.alphatab.net/).
    - [jTab](https://jtab.tardate.com/) mostly a means of describing chords
    - [note2tab](https://github.com/mrclksr/note2tab) older c program with kind of janky notation
    - [SawNote](https://github.com/xsawyerx/sawnotr) a fairly straightforward text to tab format with a Perl implementation.
    - [Tulip](https://github.com/rafael-santiago/tulip) another take on textual notation based on a 1966a book by Fernando Azevedo (Brazillian). One unique aspect is using *places* to denote string and fret (e.g. `12` is first string second fret).
    - [NoodleTab](https://github.com/jrd730/NoodleTab) is a textual fret notation that's pretty compact and converts to tablature in JavaScript. Includes a PEG grammar.

With all of these, the hardest thing is figuring out the docs. Good examples are a bonus.

- testing without testify
  - https://antonz.org/do-not-testify/
  - https://www.alexedwards.net/blog/the-9-go-test-assertions-i-use
