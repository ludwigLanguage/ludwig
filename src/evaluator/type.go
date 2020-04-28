package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/values"
)

func evalTypeIdent(n *ast.TypeIdent, consts *values.SymTab) values.Value {
	return &values.TypeIdent{n.Assoc_Type, n.Tok}
}
