package ast

import (
	"fmt"

	"ludwig/src/tokens"
)

type List struct {
	Entries []Node
	Tok     tokens.Token
}

func (l *List) PrintAll(tab string) {
	fmt.Print(l.Stringify(tab))
}

func (l *List) Stringify(tab string) string {
	rtrnStr := ""
	rtrnStr += tab + "<List>\n"

	for _, i := range l.Entries {
		rtrnStr += i.Stringify(tab + "\t")
	}

	rtrnStr += tab + "<\\List>"
	return rtrnStr
}

func (l *List) GetTok() tokens.Token {
	return l.Tok
}

func (l *List) Type() string {
	return LIST
}

/////////////////////////////////////////////////

type Index struct {
	Src   Node
	Index Node
	Tok   tokens.Token
}

func (i *Index) PrintAll(tab string) {
	fmt.Print(i.Stringify(tab))
}
func (i *Index) Stringify(tab string) string {
	rtrnStr := ""
	rtrnStr += tab + "<Index>\n"
	rtrnStr += tab + "<Source>\n"
	rtrnStr += i.Src.Stringify(tab + "\t")
	rtrnStr += tab + "<\\Source>\n"
	rtrnStr += tab + "<Value>\n"
	rtrnStr += i.Index.Stringify(tab + "\t")
	rtrnStr += tab + "<\\Value>\n"
	rtrnStr += tab + "<\\Index>"

	return rtrnStr
}

func (i *Index) GetTok() tokens.Token {
	return i.Tok
}

func (i *Index) Type() string {
	return INDEX
}

/////////////////////////////////////////////////

type Slice struct {
	Src   Node
	Start Node
	End   Node
	Tok   tokens.Token
}

func (s *Slice) PrintAll(tab string) {
	fmt.Print(s.Stringify(tab))
}

func (s *Slice) Stringify(tab string) string {
	rtrnStr := ""
	rtrnStr += tab + "<Slice>\n"
	rtrnStr += tab + "<Source>\n"
	rtrnStr += s.Src.Stringify(tab + "\t")
	rtrnStr += tab + "<\\Source>\n"
	rtrnStr += tab + "<Start>\n"
	rtrnStr += s.Start.Stringify(tab + "\t")
	rtrnStr += tab + "<\\Start>\n"

	if s.End == nil {
		rtrnStr += "<NoEnd>"
	} else {
		rtrnStr += tab + "<End>"
		rtrnStr += s.End.Stringify(tab + "\t")
		rtrnStr += tab + "<\\End>"
	}
	rtrnStr += tab + "<\\Slice>"
	return rtrnStr
}

func (s *Slice) GetTok() tokens.Token {
	return s.Tok
}

func (s *Slice) Type() string {
	return SLICE
}
