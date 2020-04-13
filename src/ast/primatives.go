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
	fmt.Print(n.Stringify(tab))
}
func (n *Number) Stringify(tab string) string {
	return fmt.Sprintf("%s<Number value=%f>\n", tab, n.Value)
}

func (n *Number) GetTok() tokens.Token {
	return n.Tok
}

func (n *Number) Type() string {
	return NUM
}

/////////////////////////////////////////////////

type String struct {
	Value string
	Tok   tokens.Token
}

func (s *String) PrintAll(tab string) {
	fmt.Print(s.Stringify(tab))
}
func (s *String) Stringify(tab string) string {
	return fmt.Sprintf("%s<String value='%s'>\n", tab, s.Value)
}

func (s *String) GetTok() tokens.Token {
	return s.Tok
}

func (s *String) Type() string {
	return STR
}

/////////////////////////////////////////////////

type Boolean struct {
	Value bool
	Tok   tokens.Token
}

func (b *Boolean) PrintAll(tab string) {
	fmt.Print(b.Stringify(tab))
}
func (b *Boolean) Stringify(tab string) string {
	return fmt.Sprintf("%s<Boolean value=%v>\n", tab, b.Value)
}

func (b *Boolean) GetTok() tokens.Token {
	return b.Tok
}

func (b *Boolean) Type() string {
	return BOOL
}
/////////////////////////////////////////////////

type Identifier struct {
	Value string
	Tok   tokens.Token
}

func (i *Identifier) PrintAll(tab string) {
	fmt.Print(i.Stringify(tab))
}

func (i *Identifier) Stringify(tab string) string {
	return fmt.Sprintf("%s<Identifier value=%s>\n", tab, i.Value)
}

func (i *Identifier) GetTok() tokens.Token {
	return i.Tok
}

func (i *Identifier) Type() string {
	return IDENT
}

////////////////////////////////////////////////

type Nil struct {
	Tok tokens.Token
}

func (n *Nil) PrintAll(tab string) {
	fmt.Print(n.Stringify(tab))
}

func (n *Nil) Stringify(tab string) string {
	return tab + "<nil>"
}

func (n *Nil) GetTok() tokens.Token {
	return n.Tok
}

func (n *Nil) Type() string {
	return NIL
}