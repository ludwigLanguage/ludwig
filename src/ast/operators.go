package ast

import (
	"ludwig/src/tokens"
	"fmt"
)

type PrefixExpr struct {
	Expr Node
	Op   string
	Tok  tokens.Token
}

func (q *PrefixExpr) PrintAll(tab string) {
	fmt.Printf("%s<Prefix Expression Operator='%s'>\n", tab, q.Op)
	q.Expr.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Prefix Expression>")
}

func (q *PrefixExpr) GetTok() tokens.Token {
	return q.Tok
}

/////////////////////////////////////////////////

type InfixExpr struct {
	Left  Node
	Right Node
	Op    string
	Tok   tokens.Token
}

func (i *InfixExpr) PrintAll(tab string) {
	fmt.Printf("%s<Infix Expression Operator='%s'>\n", tab, i.Op)

	fmt.Println(tab, "<Left>")
	i.Left.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Left>")

	fmt.Println(tab, "<Right>")
	i.Right.PrintAll(tab + "\t ")
	fmt.Println(tab, "<\\Right>")

	fmt.Println(tab, "<\\Infix Expression>")
}

func (i *InfixExpr) GetTok() tokens.Token {
	return i.Tok
}
