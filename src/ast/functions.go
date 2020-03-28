package ast

import (
	"ludwig/src/tokens"

	"fmt"
)

type Function struct {
	Args   []*Identifier
	DoExpr Node
	Tok    tokens.Token
}

func (f *Function) PrintAll(tab string) {
	fmt.Println(tab, "<Function>")
	fmt.Println(tab, "<Args>")

	for c, i := range f.Args {
		fmt.Printf("%v <Arg%v>\n", tab, c)
		i.PrintAll(tab + "\t")
		fmt.Printf("%v <\\Arg%v>\n", tab, c)
	}
	fmt.Println(tab, "<\\Args>")

	fmt.Println(tab, "<Do>")
	f.DoExpr.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Do>")
	fmt.Println(tab, "<\\Function>")
}

func (f *Function) GetTok() tokens.Token {
	return f.Tok
}

/////////////////////////////////////////////////

type Call struct {
	CalledVal Node
	Args      []Node
	Tok       tokens.Token
}

func (c *Call) PrintAll(tab string) {
	fmt.Println(tab, "<Call>")

	fmt.Println(tab, "<CalledValue>")
	c.CalledVal.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\CalledValue>")

	fmt.Println(tab, "<Args>")
	for c, i := range c.Args {
		fmt.Printf("%v <Arg%v>\n", tab, c)
		i.PrintAll(tab + "\t")
		fmt.Printf("%v <\\Arg%v>\n", tab, c)
	}
	fmt.Println(tab, "<\\Args>")

}

func (c *Call) GetTok() tokens.Token {
	return c.Tok
}
