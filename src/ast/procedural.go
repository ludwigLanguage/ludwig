package ast

import (
	"ludwig/src/tokens"
	"fmt"
)

type Block struct {
	Body []Node
	IsScoped bool
	Tok  tokens.Token
}

func (b *Block) PrintAll(tab string) {
	fmt.Println(tab, "<Block>")

	for _, i := range b.Body {
		i.PrintAll(tab + "\t")
	}

	fmt.Println(tab, "</Block>")
}

func (b *Block) GetTok() tokens.Token {
	return b.Tok
}

/////////////////////////////////////////////////

type IfEl struct {
	Cond     Node
	Do       Node
	ElseExpr Node
	Tok      tokens.Token
}

func (i *IfEl) PrintAll(tab string) {
	fmt.Println(tab, "<IfEl>")

	fmt.Println(tab, "<Cond>")
	i.Cond.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Cond>")

	fmt.Println(tab, "<Do>")
	i.Do.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Do>")

	fmt.Println(tab, "<Else>")
	i.ElseExpr.PrintAll(tab + "\t")
	fmt.Println(tab, "<\\Else>")

	fmt.Println(tab, "<\\IfEl>")
}

func (i *IfEl) GetTok() tokens.Token {
	return i.Tok
}

//////////////////////////////////////////////

type Import struct {
	Filename Node
	Tok tokens.Token
}

func (i *Import) PrintAll(tab string) {
	fmt.Println(tab + "<import>")
	i.Filename.PrintAll("\t"+tab)
	fmt.Println("<\\import>")
}

func (i *Import) GetTok() tokens.Token {
	return i.Tok 
}
