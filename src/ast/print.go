package ast

import (
	"fmt"
	"ludwig/src/tokens"
)

type Print struct {
	Tok  tokens.Token
	Args []Node
}

func (p Print) GetTok() tokens.Token {
	return p.Tok
}

func (p Print) Stringify(tab string) string {
	rtrnVal := tab + "<print>\n"
	for _, i := range p.Args {
		rtrnVal += i.Stringify(tab + "\t")
	}
	return rtrnVal + tab + "<\\print>\n"
}

func (p Print) PrintAll(tab string) {
	fmt.Print(p.Stringify(tab))
}

func (p Print) Type() byte {
	return PRINT
}
