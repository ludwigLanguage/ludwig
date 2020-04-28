package ast

import (
	"fmt"
	"ludwig/src/tokens"
)

type TypeIdent struct {
	Assoc_Type string
	Tok        tokens.Token
}

func (t *TypeIdent) Stringify(tab string) string {
	return tab + "<type_ident=" + t.Assoc_Type + ">\n"
}

func (t *TypeIdent) PrintAll(tab string) {
	fmt.Print(t.Stringify(tab))
}

func (t *TypeIdent) Type() string {
	return T_IDENT
}

func (t *TypeIdent) GetTok() tokens.Token {
	return t.Tok
}
