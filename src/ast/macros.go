package ast

import(
	"fmt"
	"ludwig/src/tokens"
)

type Quote struct {
	Expr Node
	Tok tokens.Token
}

func (q *Quote) PrintAll(tab string) {
	fmt.Print(q.Stringify(tab))
}

func (q *Quote) Stringify(tab string) string {
	return q.Expr.Stringify(tab)
}

func (q *Quote) GetTok() tokens.Token {
	return q.Tok
}

//////////////////////////////////////////

type UnQuote struct {
	Expr Node
	Tok tokens.Token
}

func (q *UnQuote) PrintAll(tab string) {
	fmt.Print(q.Stringify(tab))
}

func (q *UnQuote) Stringify(tab string) string {
	return q.Expr.Stringify(tab)
}

func (q *UnQuote) GetTok() tokens.Token {
	return q.Tok
}