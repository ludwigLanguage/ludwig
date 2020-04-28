package values

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
	"strconv"
)

type Function struct {
	Args       []*ast.Identifier
	Expr       ast.Node
	Consts     *SymTab
	IsVariadic bool
	Tok        tokens.Token
}

func (f *Function) Stringify() string {
	return "fn(" + strconv.Itoa(len(f.Args)) + ")"
}

func (f *Function) Type() string {
	return FUNC
}

func (f *Function) GetTok() tokens.Token {
	return f.Tok
}

//////////////////////////////////////////////////

type Struct struct {
	Consts *SymTab
	Body   ast.Node
	Tok    tokens.Token
}

func (s *Struct) Stringify() string {
	return "struct()"
}

func (s *Struct) Type() string {
	return STRUCT
}

func (s *Struct) GetTok() tokens.Token {
	return s.Tok
}

///////////////////////////////////////////

type Object struct {
	Consts *SymTab
	Tok    tokens.Token
}

func (o *Object) Stringify() string {
	return "object()"
}

func (o *Object) Type() string {
	return OBJ
}

func (o *Object) GetTok() tokens.Token {
	return o.Tok
}
