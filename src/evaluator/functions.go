package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
	"strconv"
)

func evalFunc(n *ast.Function, consts *values.SymTab) values.Value {
	fnC := values.NewSymTab()
	fnC.AddValsFrom(consts)

	return &values.Function{n.Args, n.DoExpr, fnC, n.IsVariadic, n.GetTok()}
}

func evalCall(n *ast.Call, consts *values.SymTab) values.Value {
	//consts.PrintAll()
	calledVal := EvalExpr(n.CalledVal, consts)

	switch calledVal := calledVal.(type) {
	case *values.Function:
		return evalFnCall(calledVal, n, consts)
	case *values.Builtin:
		return evalBuiltinCall(calledVal, n, consts)
	case *values.Struct:
		return evalStructCall(calledVal, n, consts)
	default:
		Type := calledVal.Type()
		message.RaiseError("Type", "Cannot call type '"+Type+"'", calledVal.GetTok())
	}
	return NIL
}

func evalFnCall(fn *values.Function, call *ast.Call, consts *values.SymTab) values.Value {
	newFnC := newSymTabCopy(fn.Consts)

	//Insert 'recurse' function for tail recursion
	newC := newSymTabCopy(newFnC)
	newFn := &values.Function{fn.Args, fn.Expr, newC, fn.IsVariadic, fn.GetTok()}
	newFnC.SetVal("recurse", newFn)
	//

	if (len(fn.Args) != len(call.Args)) && !fn.IsVariadic {
		message.RaiseError("Argument", "Expected "+strconv.Itoa(len(fn.Args))+" argument(s)", call.Tok)
	} else if !(len(call.Args) >= len(fn.Args)) {
		message.RaiseError("Argument", "Expected at least "+strconv.Itoa(len(fn.Args))+" argument(s)", call.Tok)
	}

	if !fn.IsVariadic {
		for c, i := range fn.Args {
			newFnC.SetVal(i.Value, EvalExpr(call.Args[c], consts))
		}
	} else {
		for c, i := range fn.Args[:len(fn.Args)-1] {
			newFnC.SetVal(i.Value, EvalExpr(call.Args[c], consts))
		}

		//Make a list containing the remaining args, and insert
		//it into the function's symbol table with the proper identifier
		id := fn.Args[len(fn.Args)-1]
		lst := &values.List{[]values.Value{}, call.GetTok()}
		for _, i := range call.Args[len(fn.Args)-1:] {
			val := EvalExpr(i, consts)
			lst.Values = append(lst.Values, val)
		}

		newFnC.SetVal(id.Value, lst)
	}

	rtrnVal := EvalExpr(fn.Expr, newFnC)
	newFnC.RmVal("recurse")
	consts.AddValsFromExcept(newFnC, fn.Args) //Remove all values that are not explicitly created in function

	return rtrnVal
}

func evalBuiltinCall(builtin *values.Builtin, call *ast.Call, consts *values.SymTab) values.Value {
	vals := []values.Value{}
	for _, i := range call.Args {
		vals = append(vals, EvalExpr(i, consts))
	}

	return builtin.Fn(vals, call.GetTok())
}
