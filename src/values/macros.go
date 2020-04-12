package values

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

type QuotedVal struct {
	Node ast.Node
	Tok tokens.Token
}

func (q *QuotedVal) Stringify() string {
	return q.Node.Stringify("")
}

func (q *QuotedVal) Type() string {
	return QUOTE
}

func (q *QuotedVal) GetTok() tokens.Token {
	return q.Tok
}
