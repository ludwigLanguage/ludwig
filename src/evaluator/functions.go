package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
	"strconv"
)

func evalFunc(n *ast.Function, consts *values.SymTab, log *message.Log) values.Value {
	fnC := values.NewSymTab()
	fnC.AddValsFrom(consts)

	return &values.Function{n.Args, n.DoExpr, fnC, n.IsVariadic, n.GetTok()}
}

func evalCall(n *ast.Call, consts *values.SymTab, log *message.Log) values.Value {
	calledVal := EvalExpr(n.CalledVal, consts, log)

	log.Add(getIdFrom(n), n.GetTok())

	var rtrnVal values.Value
	switch calledVal := calledVal.(type) {
	case *values.Function:
		rtrnVal = evalFnCall(calledVal, n, consts, log)
	case *values.Builtin:
		rtrnVal = evalBuiltinCall(calledVal, n, consts, log)
	case *values.Struct:
		rtrnVal = evalStructCall(calledVal, n, consts, log)
	default:
		Type := calledVal.Type()
		message.RuntimeErr("Type", "Cannot call type '"+Type+"'", calledVal.GetTok(), log)
	}

	log.Rm(getIdFrom(n))

	return rtrnVal
}

func evalFnCall(fn *values.Function, call *ast.Call, consts *values.SymTab, log *message.Log) values.Value {
	newFnC := newSymTabCopy(fn.Consts)

	//Insert 'recurse' function for tail recursion
	newC := newSymTabCopy(newFnC)
	newFn := &values.Function{fn.Args, fn.Expr, newC, fn.IsVariadic, fn.GetTok()}
	newFnC.SetVal("recurse", newFn)
	//

	if (len(fn.Args) != len(call.Args)) && !fn.IsVariadic {
		message.RuntimeErr("Argument", "Expected "+strconv.Itoa(len(fn.Args))+" argument(s)", call.Tok, log)
	} else if !(len(call.Args) >= len(fn.Args)) {
		message.RuntimeErr("Argument", "Expected at least "+strconv.Itoa(len(fn.Args))+" argument(s)", call.Tok, log)
	}

	if !fn.IsVariadic {
		for c, i := range fn.Args {
			newFnC.SetVal(i.Value, EvalExpr(call.Args[c], consts, log))
		}
	} else {
		for c, i := range fn.Args[:len(fn.Args)-1] {
			newFnC.SetVal(i.Value, EvalExpr(call.Args[c], consts, log))
		}

		//Make a list containing the remaining args, and insert
		//it into the function's symbol table with the proper identifier
		id := fn.Args[len(fn.Args)-1]
		lst := &values.List{[]values.Value{}, call.GetTok()}
		for _, i := range call.Args[len(fn.Args)-1:] {
			val := EvalExpr(i, consts, log)
			lst.Values = append(lst.Values, val)
		}

		newFnC.SetVal(id.Value, lst)
	}

	rtrnVal := EvalExpr(fn.Expr, newFnC, log)
	newFnC.RmVal("recurse")

	return rtrnVal
}

func evalBuiltinCall(builtin *values.Builtin, call *ast.Call, consts *values.SymTab, log *message.Log) values.Value {
	vals := []values.Value{}
	for _, i := range call.Args {
		vals = append(vals, EvalExpr(i, consts, log))
	}

	return builtin.Fn(vals, call.GetTok(), log)
}

func getIdFrom(n *ast.Call) string {
	if n.CalledVal.Type() == ast.INFIX {
		cv := n.CalledVal.(*ast.InfixExpr)
		if cv.Op == "." {
			var left string
			if cv.Left.Type() != ast.IDENT {
				left = "lambda_obj"
			} else {
				left = cv.Left.(*ast.Identifier).Value
			}

			var right string
			if cv.Right.Type() != ast.IDENT {
				right = "lambda()"
			} else {
				right = cv.Right.(*ast.Identifier).Value + "()"
			}

			return left + "." + right
		} else {
			return "lambda()"
		}
	} else if n.CalledVal.Type() == ast.IDENT {
		return n.CalledVal.(*ast.Identifier).Value + "()"
	} else {
		return "lambda()"
	}
}
