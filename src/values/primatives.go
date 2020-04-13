package values

import (
	"ludwig/src/tokens"
	"ludwig/src/ast"
	"strconv"
)

type String struct {
	Value string
	Tok   tokens.Token
}

func (s *String) Stringify() string {
	return s.Value
}

func (s *String) Type() string {
	return STR
}

func (s *String) GetTok() tokens.Token {
	return s.Tok
}

func (s *String) ConvertToAst() ast.Node {
	return &ast.String {s.Value, s.Tok}
}

/////////////////////////////////////////////////

type Number struct {
	Value float64
	Tok   tokens.Token
}

func (n *Number) Stringify() string {
	return strconv.FormatFloat(n.Value, 'f', -1, 64)
}

func (n *Number) Type() string {
	return NUM
}

func (n *Number) GetTok() tokens.Token {
	return n.Tok
}

func (n *Number) ConvertToAst() ast.Node {
	return &ast.Number {n.Value, n.Tok}
}
/////////////////////////////////////////////////

type Boolean struct {
	Value bool
	Tok   tokens.Token
}

func (b *Boolean) Stringify() string {
	return strconv.FormatBool(b.Value)
}

func (b *Boolean) Type() string {
	return BOOL
}

func (b *Boolean) GetTok() tokens.Token {
	return b.Tok
}

func (b *Boolean) ConvertToAst() ast.Node {
	return &ast.Boolean {b.Value, b.Tok}
}

/////////////////////////////////////////

type Nil struct {
	Tok tokens.Token
}

func (n *Nil) Stringify() string {
	return "nil"
}

func (n *Nil) Type() string {
	return NIL
}

func (n *Nil) GetTok() tokens.Token {
	return n.Tok
}

func (n *Nil) ConvertToAst() ast.Node {
	return &ast.Nil {n.Tok}
}

