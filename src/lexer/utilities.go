/* This file contains various utility functions for the lexer
 * to use while tokenizeing the text
 */
package lexer

import (
	"ludwig/src/message"
	"ludwig/src/tokens"
)

const (
	EOF byte = 0
)

func (l *Lexer) curString() string {
	return string(l.src.CurChar)
}

func (l *Lexer) nextString() string {
	return string(l.src.NextChar)
}

func (l *Lexer) IsDone() bool {
	return l.CurTok.Alias == tokens.EOF
}

func (l *Lexer) skipWhitespace() {
	isWhitespace := func(ch byte) bool {
		return ch == ' ' || ch == '\t' || ch == '\r'
	}

	for isWhitespace(l.src.CurChar) {
		l.src.MoveUp()
	}
}

func (l *Lexer) setTok(v string, a byte) {
	tok := tokens.Token{l.src.Filename, l.src.LineNo,
		l.src.ColumnNo, v, a}

	l.CurTok = tok
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isQuote(ch byte) bool {
	return (ch == '"') || (string(ch) == "'") || (ch == '`')
}

func (l *Lexer) buildIdent() string {
	str := ""

	for isLetter(l.src.CurChar) || isDigit(l.src.CurChar) {
		str += string(l.src.CurChar)
		l.src.MoveUp()
	}

	return str
}

func (l *Lexer) buildNum() string {
	str := ""

	for isDigit(l.src.CurChar) || l.src.CurChar == '.' {
		str += string(l.src.CurChar)
		l.src.MoveUp()
	}

	return str
}

func (l *Lexer) buildStr() string {
	quote := l.src.CurChar
	l.src.MoveUp()

	str := ""

	for l.src.CurChar != quote  && l.src.CurChar != EOF {

		if l.src.CurChar == '\\' {
			str += l.processEscapeChars()
		} else {
			str += string(l.src.CurChar)
		}

		l.src.MoveUp()
	}

	if l.src.CurChar == EOF {
		message.Error(l.src.Filename, "Syntax", "Expected closing quote", l.src.LineNo, l.src.ColumnNo)
	}

	l.src.MoveUp() //Skip over closing quote

	return str
}

func (l *Lexer) processEscapeChars() string {
	l.src.MoveUp() //Move over '\'
	if l.src.CurChar == 'n' {
		return "\n"
	} else if l.src.CurChar == 't' {
		return "\t"
	} else if l.src.CurChar == '\\' {
		return "\\"
	}
	message.Error(l.src.Filename, "Syntax", "Invalid escape char", l.src.LineNo, l.src.ColumnNo)
	return ""
}

func (l *Lexer) skipComments() {
	if l.src.CurChar == '#' {
		for l.src.CurChar != '\n' && l.src.CurChar != EOF {
			l.src.MoveUp()
		}
		l.src.MoveUp()

	} else if (string(l.src.CurChar) + string(l.src.NextChar)) == "|#" {

		for (string(l.src.CurChar)+string(l.src.NextChar)) != "#|" && l.src.CurChar != EOF {
			l.src.MoveUp()
		}
		l.src.MoveUp()
		l.src.MoveUp()
	}

	if l.src.CurChar == EOF {
		l.src.CurChar = ')'
	}
}
