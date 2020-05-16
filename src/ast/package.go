package ast

import (
	"fmt"
	"ludwig/src/tokens"
)

type Package struct {
	Id          Identifier
	PublicBody  []InfixExpr
	PrivateBody []InfixExpr
	Tok         tokens.Token
}

func (p Package) Type() byte {
	return PACK
}

func (p Package) GetTok() tokens.Token {
	return p.Tok
}

func (p Package) Stringify(tab string) string {
	rtrnVal := tab + "<package>\n"
	rtrnVal += p.Id.Stringify(tab + "\t")

	rtrnVal += tab + "<private>\n"
	for _, i := range p.PrivateBody {
		rtrnVal += i.Stringify(tab + "\t")
	}
	rtrnVal += tab + "<\\private>\n"

	rtrnVal += tab + "<public>\n"
	for _, i := range p.PublicBody {
		rtrnVal += i.Stringify(tab + "\t")
	}
	rtrnVal += tab + "<\\private>\n"

	return rtrnVal + tab + "<\\package>\n"
}

func (p Package) PrintAll(tab string) {
	fmt.Print(p.Stringify(tab))
}
