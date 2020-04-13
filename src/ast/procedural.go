package ast

import (
	"ludwig/src/tokens"
	"fmt"
	"strconv"
)

type Block struct {
	Body []Node
	IsScoped bool
	Tok  tokens.Token
}

func (b *Block) PrintAll(tab string) {
	fmt.Print(b.Stringify(tab))
}
func (b *Block) Stringify(tab string) string {
	rtrnVal := ""
	rtrnVal += tab + "<Block>\n"

	rtrnVal += tab + "<IsScoped?=" + strconv.FormatBool(b.IsScoped) + ">\n"

	for _, i := range b.Body {
		rtrnVal += i.Stringify(tab + "\t")
	}

	rtrnVal += tab + "</Block>\n"
	return rtrnVal
}

func (b *Block) GetTok() tokens.Token {
	return b.Tok
}

func (b *Block) Type() string {
	return BLOCK
}

/////////////////////////////////////////////////

type IfEl struct {
	Cond     Node
	Do       Node
	ElseExpr Node
	Tok      tokens.Token
}

func (i *IfEl) PrintAll(tab string) {
	fmt.Print(i.Stringify(tab))
}

func (i *IfEl) Stringify(tab string) string {
	rtrnVal := ""
	rtrnVal += tab + "<IfEl>\n"

	rtrnVal += tab + "<Cond>\n"
	rtrnVal += i.Cond.Stringify(tab + "\t")
	rtrnVal += tab + "<\\Cond>\n"

	rtrnVal += tab + "<Do>\n"
	rtrnVal += i.Do.Stringify(tab + "\t")
	rtrnVal += tab + "<\\Do>\n"

	rtrnVal += tab + "<Else>\n"
	rtrnVal += i.ElseExpr.Stringify(tab + "\t")
	rtrnVal += tab + "<\\Else>\n"

	return rtrnVal + tab + "<\\IfEl>\n"
}

func (i *IfEl) GetTok() tokens.Token {
	return i.Tok
}

func (i *IfEl) Type() string {
	return IFEL
}

//////////////////////////////////////////////

type Import struct {
	Filename Node
	Tok tokens.Token
}

func (i *Import) PrintAll(tab string) {
	fmt.Print(i.Stringify(tab))
}

func (i *Import) Stringify(tab string) string {
 	rtrnVal := ""
	rtrnVal += tab + "<import>\n"
	rtrnVal += i.Filename.Stringify("\t"+tab)
	rtrnVal += "<\\import>\n"
	return rtrnVal
}

func (i *Import) GetTok() tokens.Token {
	return i.Tok 
}

func (i *Import) Type() string {
	return IMPRT
}