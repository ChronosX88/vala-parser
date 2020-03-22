package scanner

import (
	"bytes"
	"fmt"
	"io"

	"github.com/ChronosX88/vala-parser/utils"
)

const (
	eof = rune(0) // end of file
)

type Scanner struct {
	buf *bytes.Reader
}

func NewScanner(reader io.Reader) *Scanner {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reader)
	r := bytes.NewReader(buffer.Bytes())
	return &Scanner{
		buf: r,
	}
}

func (s *Scanner) Scan() Token {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if isSpecialSymbol(ch) {
		s.unread()
		return s.scanSpecSymbol()
	} else if isDigit(ch) {
		s.unread()
		return s.scanNumber()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return Token{EOF, ""}
	}

	return Token{Illegal, string(ch)}
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() Token {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Whitespace, buf.String()}
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() Token {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	//buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	parsedToken := Token{
		Kind:    -1,
		Literal: buf.String(),
	}

	// If the string matches a keyword then return that keyword.
	switch buf.String() {
	case Using.String(): // using
		parsedToken.Kind = Using
	case Namespace.String(): // namespace
		parsedToken.Kind = Namespace
	case PublicModifier.String(): // public
		parsedToken.Kind = PublicModifier
	case PrivateModifier.String(): // private
		parsedToken.Kind = PrivateModifier
	case Class.String(): // class
		parsedToken.Kind = Class
	case Var.String(): // var
		parsedToken.Kind = Var
	case Return.String(): // return
		parsedToken.Kind = Return
	case Null.String(): // null
		parsedToken.Kind = Null
	case If.String(): // if
		parsedToken.Kind = If
	case ProtectedModifier.String(): // protected
		parsedToken.Kind = ProtectedModifier
	case False.String(): // false
		parsedToken.Kind = False
	case True.String(): // true
		parsedToken.Kind = True
	case New.String(): // new
		parsedToken.Kind = New
	default:
		parsedToken.Kind = Identifier
	}

	return parsedToken
}

func (s *Scanner) scanSpecSymbol() Token {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isSpecialSymbol(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	parsedToken := Token{
		Kind:    -1,
		Literal: buf.String(),
	}

	// If the string matches a keyword then return that keyword.
	matchSpecSymbol(&parsedToken)

	if parsedToken.Kind == Illegal && len(parsedToken.Literal) > 1 { // then two or more special characters in a row detected
		for i := 0; i < len(parsedToken.Literal)-1; i++ {
			s.buf.Seek(-1, io.SeekCurrent)
		}
		parsedToken.Literal = string(utils.RuneAt(parsedToken.Literal, 0))
		matchSpecSymbol(&parsedToken)
	}

	return parsedToken
}

func (s *Scanner) scanNumber() Token {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) && !isXDigit(ch) && (ch != 'x') && (ch != '.') {
			fmt.Println(string(ch))
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	parsedToken := Token{
		Kind:    IntegerLiteral,
		Literal: buf.String(),
	}

	for _, v := range []rune(parsedToken.Literal) {
		if v == '.' {
			parsedToken.Kind = RealLiteral
		}
	}

	return parsedToken
}

func matchSpecSymbol(tok *Token) {
	switch tok.Literal {
	case Add.String(): // +
		tok.Kind = Add
	case Minus.String(): // -
		tok.Kind = Minus
	case Mult.String(): // *
		tok.Kind = Mult
	case Divide.String(): // /
		tok.Kind = Divide
	case Colon.String(): // :
		tok.Kind = Colon
	case Comma.String(): // ,
		tok.Kind = Comma
	case Semicolon.String(): // ;
		tok.Kind = Semicolon
	case OpenBrace.String(): // {
		tok.Kind = OpenBrace
	case CloseBrace.String(): // }
		tok.Kind = CloseBrace
	case StringLiteral.String(): // "
		tok.Kind = StringLiteral
	case Percent.String(): // %
		tok.Kind = Percent
	case OpenParens.String(): // (
		tok.Kind = OpenParens
	case CloseParens.String(): // )
		tok.Kind = CloseParens
	case Dot.String(): // .
		tok.Kind = Dot
	case OpenMultilineComments.String(): // /*
		tok.Kind = OpenMultilineComments
	case CloseMultilineComments.String(): // */
		tok.Kind = CloseMultilineComments
	case LambdaArrow.String(): // =>
		tok.Kind = LambdaArrow
	case Equal.String(): // ==
		tok.Kind = Equal
	case OpenSingleComments.String():
		tok.Kind = OpenSingleComments
	case Assign.String():
		tok.Kind = Assign
	case OpenBracket.String():
		tok.Kind = OpenBracket
	case CloseBracket.String():
		tok.Kind = CloseBracket
	}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.buf.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	err := s.buf.UnreadRune()
	if err != nil {
		fmt.Println("Error when unread: " + err.Error())
	}
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

func isSpecialSymbol(ch rune) bool {
	return (ch >= '!' && ch <= '/') || (ch >= ':' && ch <= '?') || (ch >= '[' && ch <= '`') || (ch >= '{' && ch <= '~') && (ch != '_')
}

func isXDigit(ch rune) bool { return (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'F') }
