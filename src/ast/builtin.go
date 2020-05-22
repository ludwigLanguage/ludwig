package ast

import (
	"fmt"
	"ludwig/src/tokens"
)

type Builtin struct {
	Tok         tokens.Token
	BuiltinName string
	Args        []Node
}

func (b Builtin) Type() byte {
	return BUILTIN
}

func (b Builtin) GetTok() tokens.Token {
	return b.Tok
}

func (b Builtin) Stringify(tab string) string {
	rtrn := tab + "<builtin=" + b.BuiltinName + ">\n"
	rtrn += tab + "<args>\n"
	for _, i := range b.Args {
		rtrn += i.Stringify(tab + "\t")
	}
	rtrn += tab + "<\\args>\n"
	rtrn += tab + "<\\builtin>"
	return rtrn
}

func (b Builtin) PrintAll(tab string) {
	fmt.Print(b.Stringify(tab))
}
