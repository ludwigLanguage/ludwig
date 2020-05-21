package ast

import (
	"fmt"
	"ludwig/src/tokens"
)

type Program struct {
	Body []InfixExpr //TODO: Convert to []*ast.InfixExpr
	Tok  tokens.Token
}

func (p Program) Type() byte {
	return PROG
}

func (p Program) Stringify(tab string) string {
	rtrnVal := tab + "<program>\n"

	for _, i := range p.Body {
		rtrnVal += i.Stringify(tab + "\t")
	}
	return rtrnVal + tab + "<\\program>\n"
}

func (p Program) PrintAll(tab string) {
	fmt.Print(p.Stringify(tab))
}

func (p Program) GetTok() tokens.Token {
	return p.Tok
}
