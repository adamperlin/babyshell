package scanner

import (
	"bufio"
	"bytes"
	"io"
)

const (
	ILLEGAL Token = iota
	EOF
	WS
	AMPERSAND
	GREAT
	GREATGREAT
	LESS
	AMPERSANDGREAT
	GREATAMPERSAND
	PIPE
	IDENT
	NEWLINE

)
 var ReservedChars []rune = []rune{ //characters that are researved for some commands/command arguments
	'+',
	'"',
	'\'',
	'\'',
	'$',
	'/',
	'.',
	'*',
	'-',
	':',
}

var eof = rune(0) //represents standard EOF

type Token int

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func isWhitespaceNotNewline(ch rune) bool {
	return ch == ' ' || ch == '\t'
}
func isNewline(ch rune) bool {
	return ch == '\n'
}
func isContiguousOperator(ch rune) bool {
	switch ch {
	case '&':
		return true
	case '>':
		return true
	}
	return false
}
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
func isDigit(ch rune) bool {
	return (ch >= '1' && ch <= '9')
}
func isOtherSpecialChar(ch rune) bool{
	for _, elem := range ReservedChars {
		if ch == elem {
			return true
		}
	}
	return false
}
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()
	if isWhitespaceNotNewline(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) || isOtherSpecialChar(ch) {
		s.unread()
		return s.scanIdent()
	}
	switch ch {
	//scan for EOF, and non-contiguous operators: i.e., |, <. Others, like '&', '>' could be contiguous, i.e.: '&>', '>&', '>>', etc.
	case eof:
		return EOF, ""
	case '\n':
		return NEWLINE,string(ch)
	case '|':
		return PIPE, string(ch)
	case '<':
		return LESS, string(ch)
	case '>':
		return s.switchGreat()
	case '&':
		return s.switchAmpersand()
	}
	return ILLEGAL, string(ch)
}

func (s *Scanner) switchGreat()(Token,string){
	nextch := s.read()
	switch nextch {
	case '>':
		return GREATGREAT, ">>"
	case '&':
		return GREATAMPERSAND, ">&"
	}
	s.unread()
	return GREAT, ">"
}

func (s *Scanner) switchAmpersand()(Token, string){
	nextch := s.read()
	switch nextch {
		case '>':
		return AMPERSANDGREAT, "&>"
	}
	s.unread()
	return AMPERSAND, "&"
}
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buff bytes.Buffer
	buff.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespaceNotNewline(ch) {
			s.unread()
			break
		} else {
			buff.WriteRune(ch)
		}
	}
	return WS, buff.String()
}
func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buff bytes.Buffer
	buff.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && !isOtherSpecialChar(ch) {
			s.unread()
			break
		} else {
			_, _ = buff.WriteRune(ch)
		}
	}
	str := buff.String()
	return IDENT, str
}
