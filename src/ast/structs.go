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
	fmt.Print(c.Stringify(tab))
}
func (s *Struct) Stringify(tab string) string {
	rtrnVal := tab + "<Struct>\n"

	rtrnVal += tab + "<Body>\n"
	rtrnVal += s.Body.Stringify(tab + "\t")
	rtrnVal += tab + "<\\Body>\n"

	return rtrnVal + tab+ "<\\Struct>\n"
}

func (c *Struct) GetTok() tokens.Token {
	return c.Tok
}
