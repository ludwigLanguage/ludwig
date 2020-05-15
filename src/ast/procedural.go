package ast

import (
	"fmt"
	"ludwig/src/tokens"
	"strconv"
)

type Block struct {
	Body     []Node
	IsScoped bool
	Tok      tokens.Token
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

func (b *Block) Type() byte {
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

func (i *IfEl) Type() byte {
	return IFEL
}

///////////////////////////////////////////

type For struct {
	IndexNumIdent *Identifier
	IndexIdent    *Identifier
	List          Node
	DoExpr        Node
	IsScoped      bool //This is necessary, trust me
	Tok           tokens.Token
}

func (f *For) Stringify(tab string) string {
	rtrnVal := tab + "<for " + f.IndexNumIdent.Stringify("")
	rtrnVal += ", "

	rtrnVal += f.IndexIdent.Stringify("")
	rtrnVal += ">\n"

	rtrnVal += tab + "<list>\n"
	rtrnVal += f.List.Stringify(tab + "\t")
	rtrnVal += tab + "<\\list>\n"

	rtrnVal += tab + "<Do>\n"
	rtrnVal += f.DoExpr.Stringify(tab + "\t")
	rtrnVal += tab + "<\\Do>\n"
	return rtrnVal + tab + "<\\for>\n"
}

func (f *For) PrintAll(tab string) {
	fmt.Print(f.Stringify(tab))
}

func (f *For) GetTok() tokens.Token {
	return f.Tok
}

func (f *For) Type() byte {
	return FOR
}

// While ///////////////////////////////////
type While struct {
	Cond     Node
	Body     Node
	IsScoped bool
	Tok      tokens.Token
}

func (w *While) Stringify(tab string) string {
	rtrnVal := tab + "<while>\n"
	rtrnVal += tab + "<conditional>\n"
	rtrnVal += w.Cond.Stringify(tab + "\t")
	rtrnVal += tab + "<\\conditional>\n"
	rtrnVal += tab + "<body>\n"
	rtrnVal += w.Body.Stringify(tab + "\t")
	rtrnVal += tab + "<\\body>\n"
	return rtrnVal
}

func (w *While) PrintAll(tab string) {
	fmt.Print(w.Stringify(tab))
}

func (w *While) GetTok() tokens.Token {
	return w.Tok
}

func (w *While) Type() byte {
	return WHILE
}
