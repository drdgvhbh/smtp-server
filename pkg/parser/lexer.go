package parser

import "github.com/alecthomas/participle/lexer"

func Lexer() (lexer.Definition, error) {
	return lexer.Regexp(`(?P<FQDN>[^\s]*\.[^\s]*)|(?P<HELO>HELO)|(?P<ws>\s+)`)
}
