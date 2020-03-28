package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
	"strconv"
	"fmt"
)

func evalFunc(n *ast.Function, consts *values.SymTab) values.Value {
	fnC := values.NewSymTab()
	fnC.AddValsFrom(consts)

	return &values.Function{n.Args, n.DoExpr, fnC, n.GetTok()}
}

func evalCall(n *ast.Call, consts *values.SymTab) values.Value {

	calledVal := EvalExpr(n.CalledVal, consts)

	switch calledVal := calledVal.(type) {
	case *values.Function:
		return evalFnCall(calledVal, n, consts)
	case *values.Builtin:
		return evalBuiltinCall(calledVal, n, consts)
	case *values.Struct:
		return evalStructCall(calledVal, n, consts)
	default:
		Type := fmt.Sprintf("%T", calledVal)
		message.RaiseError("Type", "Cannot call type '" + Type + "'", calledVal.GetTok())
	}
	return NIL
}

func evalFnCall(fn *values.Function, call *ast.Call, consts *values.SymTab) values.Value {
	newFnC := newSymTabCopy(fn.Consts)

	//Insert 'recurse' function for tail recursion
	newC := newSymTabCopy(newFnC)
	newFn := &values.Function{fn.Args, fn.Expr, newC, fn.GetTok()}
	newFnC.SetVal("recurse", newFn)
	//

	if len(fn.Args) != len(call.Args) {
		message.RaiseError("Argument", "Expected "+strconv.Itoa(len(fn.Args))+" argument(s)", call.Tok)
	}

	for c, i := range fn.Args {
		newFnC.SetVal(i.Value, EvalExpr(call.Args[c],  consts))
	}

	return EvalExpr(fn.Expr, newFnC)
}

func evalBuiltinCall(builtin *values.Builtin, call *ast.Call, consts *values.SymTab) values.Value {
	vals := []values.Value{}
	for _, i := range call.Args {
		vals = append(vals, EvalExpr(i, consts))
	}

	return builtin.Fn(vals)
}

func evalStructCall(strct *values.Struct, call *ast.Call, consts *values.SymTab) values.Value {

	initFn := strct.Consts.GetVal("__init__")
	if initFn != nil {
		initCall := call
		initCall.CalledVal = &ast.Identifier {"__init__", call.GetTok()} 
		evalCall(initCall, strct.Consts)
	}


	return &values.Object {strct.Consts, strct.GetTok()}
}
