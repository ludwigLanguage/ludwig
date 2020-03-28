/* This file contains the 'lexer' object which turns characters from the
 * 'source.Source' object into tokens that represent the symbols and values
 * found in the file. These tokens will next be used by the parser to creat an ast
 */
package lexer

import (
	"ludwig/src/message"
	"ludwig/src/source"
	"ludwig/src/tokens"
)

type Lexer struct {
	src    *source.Source
	CurTok tokens.Token
}

func New(src *source.Source) *Lexer {
	l := &Lexer{}

	l.src = src
	l.MoveUp() //Fill in 'l.CurTok' value to be used by parser

	return l
}

func (l *Lexer) MoveUp() {
	l.skipWhitespace()

	focusString := l.curString() + l.nextString()
	alias, ok := doubleChar[focusString]

	if ok {
		l.setTok(focusString, alias)
		l.src.MoveUp()
		l.src.MoveUp()
		return
	} else if focusString == "|#" {
		l.skipComments()
		l.MoveUp()
		return
	}

	focusString = l.curString()
	alias, ok = singleChar[focusString]

	if ok {
		l.setTok(focusString, alias)
		l.src.MoveUp()
		return
	} else if focusString == "#" {
		l.skipComments()
		l.MoveUp()
		return
	}

	if isLetter(l.src.CurChar) {
		id := l.buildIdent()

		alias, ok = keywords[id]
		if ok {
			l.setTok(id, alias)
		} else {
			l.setTok(id, tokens.IDENT)
		}
		return

	} else if isDigit(l.src.CurChar) {
		num := l.buildNum()
		l.setTok(num, tokens.NUM)
		return

	} else if isQuote(l.src.CurChar) {
		str := l.buildStr()
		l.setTok(str, tokens.STR)
		return

	} else {
		message.Error(l.src.Filename, "Lexer",
			"Could not lex character this value '"+string(l.src.CurChar)+"'", l.src.LineNo, l.src.ColumnNo)
	}

}

func (l *Lexer) Src() *source.Source {
	return l.src
}
