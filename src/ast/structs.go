package ast

import (
	"ludwig/src/tokens"

	"fmt"
)

type Struct struct {
	Tok  tokens.Token
	Body Node
}

func (c *Struct) PrintAll(tab string) {
	fmt.Println(tab, "<Struct>")

	fmt.Println(tab, "<Body>")
	c.Body.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Body>")

	fmt.Println(tab, "<\\Struct>")
}

func (c *Struct) GetTok() tokens.Token {
	return c.Tok
}
