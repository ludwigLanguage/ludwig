package evaluator

import (
	"ludwig/src/values"
	"ludwig/src/ast"
)

func evalStruct(n *ast.Struct, consts *values.SymTab) values.Value {
	newSymTab := values.NewSymTab()
	newSymTab.AddValsFrom(consts)
	self := &values.Object {newSymTab,  n.GetTok()}
	newSymTab.SetVal("self", self)
	EvalExpr(n.Body, newSymTab)

	return &values.Struct {newSymTab, n.Body, n.GetTok()}
}
