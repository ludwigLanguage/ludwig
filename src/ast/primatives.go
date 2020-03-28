package ast

import (
	"ludwig/src/tokens"
	"fmt"
)

type Number struct {
	Value float64
	Tok   tokens.Token
}

func (n *Number) PrintAll(tab string) {
	fmt.Printf("%s<Number value=%f>\n", tab, n.Value)
}

func (n *Number) GetTok() tokens.Token {
	return n.Tok
}

/////////////////////////////////////////////////

type String struct {
	Value string
	Tok   tokens.Token
}

func (s *String) PrintAll(tab string) {
	fmt.Printf("%s<String value='%s'>\n", tab, s.Value)
}

func (s *String) GetTok() tokens.Token {
	return s.Tok
}

/////////////////////////////////////////////////

type Boolean struct {
	Value bool
	Tok   tokens.Token
}

func (b *Boolean) PrintAll(tab string) {
	fmt.Printf("%s<Boolean value=%v>\n", tab, b.Value)
}

func (b *Boolean) GetTok() tokens.Token {
	return b.Tok
}

/////////////////////////////////////////////////

type Identifier struct {
	Value string
	Tok   tokens.Token
}

func (i *Identifier) PrintAll(tab string) {
	fmt.Printf("%s<Identifier value=%s>\n", tab, i.Value)
}

func (i *Identifier) GetTok() tokens.Token {
	return i.Tok
}

////////////////////////////////////////////////

type Nil struct {
	Tok tokens.Token
}

func (n *Nil) PrintAll(tab string) {
	fmt.Println(tab + "<nil>")
}

func (n *Nil) GetTok() tokens.Token {
	return n.Tok
}