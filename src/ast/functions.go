package ast

import (
	"ludwig/src/tokens"

	"fmt"
)

type Function struct {
	Args       []*Identifier
	DoExpr     Node
	IsVariadic bool
	Tok        tokens.Token
}

func (f *Function) Stringify(tab string) string {
	rtrnStr := ""

	rtrnStr += tab + "<Function>\n"

	if f.IsVariadic {
		rtrnStr += tab + "<Is Variadic=true>\n"
	} else {
		rtrnStr += tab + "<Is Variadic=false>\n"
	}

	rtrnStr += tab + "<Args>\n"

	for c, i := range f.Args {
		rtrnStr += fmt.Sprintf("%v<Arg%v>\n", tab+"\t", c)
		rtrnStr += i.Stringify(tab + "\t")
		rtrnStr += fmt.Sprintf("%v<\\Arg%v>\n", tab+"\t", c)
	}
	rtrnStr += tab + "<\\Args>\n"

	rtrnStr += tab + "<Do>\n"
	rtrnStr += f.DoExpr.Stringify(tab + "\t")
	rtrnStr += tab + "<\\Do>\n"
	rtrnStr += tab + "<\\Function>\n"

	return rtrnStr
}

func (f *Function) PrintAll(tab string) {
	fmt.Print(f.Stringify(tab))
}

func (f *Function) GetTok() tokens.Token {
	return f.Tok
}

func (f *Function) Type() byte {
	return FN
}

/////////////////////////////////////////////////

type Call struct {
	CalledVal Node
	Args      []Node
	Tok       tokens.Token
}

func (c *Call) PrintAll(tab string) {
	fmt.Print(c.Stringify(tab))
}
func (c *Call) Stringify(tab string) string {
	rtrnStr := ""

	rtrnStr += tab + "<Call>\n"

	rtrnStr += tab + "<CalledValue>\n"
	c.CalledVal.Stringify(tab + "\t")
	rtrnStr += tab + "<\\CalledValue>\n"

	rtrnStr += tab + "<Args>\n"
	for c, i := range c.Args {
		rtrnStr += fmt.Sprintf("%v<Arg%v>\n", tab, c)
		rtrnStr += i.Stringify(tab + "\t")
		rtrnStr += fmt.Sprintf("%v<\\Arg%v>\n", tab, c)
	}
	rtrnStr += tab + "<\\Args>\n"
	return rtrnStr
}

func (c *Call) GetTok() tokens.Token {
	return c.Tok
}

func (c *Call) Type() byte {
	return CALL
}
