package values

import (
	"ludwig/src/ast"
	"strconv"
)

type Function struct {
	Args []*ast.Identifier
	Expr ast.Node
	//Consts     *SymTab
	IsVariadic bool
}

func (f Function) Stringify() string {
	return "fn(" + strconv.Itoa(len(f.Args)) + ")"
}

func (f Function) Type() byte {
	return FUNC
}

//////////////////////////////////////////////////

type Struct struct {
	//Consts *SymTab
	Body ast.Node
}

func (s Struct) Stringify() string {
	return "struct()"
}

func (s Struct) Type() byte {
	return STRUCT
}

///////////////////////////////////////////

type Object struct {
	//Consts *SymTab
}

func (o Object) Stringify() string {
	return "object()"
}

func (o Object) Type() byte {
	return OBJ
}
