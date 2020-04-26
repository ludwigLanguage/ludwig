package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/values"
)

func evalStruct(n *ast.Struct, consts *values.SymTab) values.Value {

	newSymTab := values.NewSymTab()
	newSymTab.AddValsFrom(consts)

	return &values.Struct{newSymTab, n.Body, n.GetTok()}
}

func evalStructCall(strct *values.Struct, call *ast.Call, consts *values.SymTab) values.Value {

	objSt := values.NewSymTab()
	objSt.AddValsFrom(strct.Consts)

	self := &values.Object{objSt, strct.GetTok()}
	objSt.SetVal("self", self)

	EvalExpr(strct.Body, objSt)

	initFn := objSt.GetVal("__init__")
	if initFn != nil {
		initCall := call
		initCall.CalledVal = &ast.Identifier{"__init__", call.GetTok()}

		/*
			We evaluate the init in the outer symtab so that when
			we evaluate the function arguments we can access that
			outer symtab.
			However, to access __init__, we must add it. Futhermore,
			because functions have a symtab attached to it, we have
			access to the outer symtab.
		*/
		consts.SetVal("__init__", initFn)

		initConsts := values.NewSymTab()
		initConsts.AddValsFrom(consts)

		evalCall(initCall, initConsts)
		consts.RmVal("__init__")
	}

	return self
}
