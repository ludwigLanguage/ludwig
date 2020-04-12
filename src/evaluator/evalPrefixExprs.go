package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
)

func evalPrefix(n *ast.PrefixExpr, consts *values.SymTab) values.Value {
	switch n.Op {
	case "!":
		return evalNot(EvalExpr(n.Expr, consts))
	case "-":
		return evalNegative(EvalExpr(n.Expr, consts))
	default:
		message.RaiseError("Operator", "Cannot use this operator as a prefix", n.GetTok())
	}
	return NIL
}

func evalNot(v values.Value) values.Value {
	if v.Type() != values.BOOL {
		message.RaiseError("Type", "Must have a boolean for '!' got '"+v.Stringify()+"'", v.GetTok())
	}

	return &values.Boolean{!v.(*values.Boolean).Value, v.GetTok()}
}

func evalNegative(v values.Value) values.Value {
	if v.Type() != values.NUM {
		message.RaiseError("Type", "Must have number for '-'", v.GetTok())
	}

	return &values.Number{-v.(*values.Number).Value, v.GetTok()}
}
