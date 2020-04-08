package values

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"

	"strconv"
)

type Function struct {
	Args   		[]*ast.Identifier
	Expr   		ast.Node
	Consts		*SymTab
	IsVariadic 	bool
	Tok    		tokens.Token
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
	Tok tokens.Token
}

func (s *Struct) Stringify() string {
	rtrnVal := "struct ( " 
	for k, v := range s.Consts.values {
		if k != "self" {
			rtrnVal += k+":"+v.Stringify()+" "
		}
	}

	return rtrnVal + " )"
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
	Tok tokens.Token
}

func (o *Object) Stringify() string {
	rtrnVal := "object ( " 
	for k, v := range o.Consts.values {
		if k != "self" {
			rtrnVal += k+":"+v.Stringify()+" "
		}
	}

	return rtrnVal + " )"
}

func (o *Object) Type() string {
	return OBJ
}

func (o *Object) GetTok() tokens.Token {
	return o.Tok
}

//////////////////////////////////////////////
type Exec struct {
	Location string
	ShouldPrint bool
	Tok tokens.Token
}

func (e *Exec) Stringify() string {
	shouldPrint := "Non-Printing Executable"
	if e.ShouldPrint {
		shouldPrint = "Printing Executable"
	}
	return "Executable ( " + e.Location + ", " + shouldPrint + " )"
}

func (e *Exec) Type() string {
	return EXEC 
}

func (e *Exec) GetTok() tokens.Token {
	return e.Tok
}