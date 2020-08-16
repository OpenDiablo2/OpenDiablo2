// Package d2lexer contains the code for tokenizing calculation strings.
package d2lexer

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type tokenType int

const (
	// Name represents a name token, such as skill, par1 etc.
	Name tokenType = iota

	// String represents a quoted string token, such as "Sacrifice".
	String

	// Symbol represents a symbol token, such as '+', '-', '?, '.' etc.
	Symbol

	// Number represents an integer token.
	Number

	// EOF is the end-of-file token, generated when the end of data is reached.
	EOF
)

func (t tokenType) String() string {
	return []string{
		"Name",
		"String",
		"Symbol",
		"Number",
		"EOF",
	}[t]
}

// Token is a lexical token of a calculation string.
type Token struct {
	Type  tokenType
	Value string
}

func (t *Token) String() string {
	return "(" + t.Type.String() + ", " + t.Value + ")\n"
}

// Lexer is the tokenizer for calculation strings.
type Lexer struct {
	data         []byte
	CurrentToken Token
	index        int
	peeked       bool
	nextToken    Token
}

// New creates a new Lexer for tokenizing the given data.
func New(input []byte) *Lexer {
	return &Lexer{
		data: input,
	}
}

func (l *Lexer) peekNext() (byte, error) {
	if l.index+1 >= len(l.data) {
		return 0, errors.New("cannot peek")
	}

	return l.data[l.index+1], nil
}

func (l *Lexer) extractOpToken() Token {
	c := l.data[l.index]
	if c == '=' || c == '!' {
		next, ok := l.peekNext()
		if ok != nil || next != '=' {
			panic("Invalid operator at index!" + strconv.Itoa(l.index))
		} else {
			l.index += 2
			return Token{Symbol, string(c) + "="}
		}
	}

	if c == '<' || c == '>' {
		next, ok := l.peekNext()
		if ok == nil && next == '=' {
			l.index += 2
			return Token{Symbol, string(c) + "="}
		}
		l.index++

		return Token{Symbol, string(c)}
	}
	l.index++

	return Token{Symbol, string(c)}
}

func (l *Lexer) extractNumber() Token {
	var sb strings.Builder

	for l.index < len(l.data) && unicode.IsDigit(rune(l.data[l.index])) {
		sb.WriteByte(l.data[l.index])
		l.index++
	}

	return Token{Number, sb.String()}
}

func (l *Lexer) extractString() Token {
	var sb strings.Builder
	l.index++

	for l.index < len(l.data) && l.data[l.index] != '\'' {
		sb.WriteByte(l.data[l.index])
		l.index++
	}
	l.index++

	return Token{String, sb.String()}
}

func (l *Lexer) extractName() Token {
	var sb strings.Builder

	for l.index < len(l.data) &&
		(unicode.IsLetter(rune(l.data[l.index])) ||
			unicode.IsDigit(rune(l.data[l.index]))) {
		sb.WriteByte(l.data[l.index])
		l.index++
	}

	return Token{Name, sb.String()}
}

// Peek returns the next token, but does not advance the tokenizer.
// The peeked token is cached until the tokenizer advances.
func (l *Lexer) Peek() Token {
	if l.peeked {
		return l.nextToken
	}

	if l.index == len(l.data) {
		l.nextToken = Token{EOF, ""}
		return l.nextToken
	}

	for l.index < len(l.data) && unicode.IsSpace(rune(l.data[l.index])) {
		l.index++
	}

	if l.index == len(l.data) {
		l.nextToken = Token{EOF, ""}
		return l.nextToken
	}

	switch {
	case strings.IndexByte("^=!><+-/*.,:?()", l.data[l.index]) != -1:
		l.nextToken = l.extractOpToken()
	case unicode.IsDigit(rune(l.data[l.index])):
		l.nextToken = l.extractNumber()
	case l.data[l.index] == '\'':
		l.nextToken = l.extractString()
	case unicode.IsLetter(rune(l.data[l.index])):
		l.nextToken = l.extractName()
	default:
		panic("Invalid token at index: " + strconv.Itoa(l.index))
	}

	l.peeked = true

	return l.nextToken
}

// NextToken returns the next token and advances the tokenizer.
func (l *Lexer) NextToken() Token {
	if l.peeked {
		l.CurrentToken = l.nextToken
	} else {
		l.CurrentToken = l.Peek()
	}

	l.peeked = false

	return l.CurrentToken
}
