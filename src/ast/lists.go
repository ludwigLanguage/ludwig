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
	fmt.Println(tab, "<List>")

	for _, i := range l.Entries {
		i.PrintAll(tab + "\t")
	}

	fmt.Println(tab, "<\\List>")
}

func (l *List) GetTok() tokens.Token {
	return l.Tok
}

/////////////////////////////////////////////////

type Index struct {
	Src   Node
	Index Node
	Tok   tokens.Token
}

func (i *Index) PrintAll(tab string) {
	fmt.Println(tab, "<Index>\n")
	fmt.Println(tab, "<Source>")
	i.Src.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Source>")
	fmt.Println(tab, "<Value>")
	i.Index.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Value>")
}

func (i *Index) GetTok() tokens.Token {
	return i.Tok
}
