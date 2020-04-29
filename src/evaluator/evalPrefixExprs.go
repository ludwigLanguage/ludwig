package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
)

func evalPrefix(n *ast.PrefixExpr, consts *values.SymTab, log *message.Log) values.Value {
	switch n.Op {
	case "!":
		return evalNot(EvalExpr(n.Expr, consts, log), log)
	case "-":
		return evalNegative(EvalExpr(n.Expr, consts, log), log)
	default:
		message.RuntimeErr("Operator", "Cannot use this operator as a prefix", n.GetTok(), log)
	}
	return NIL
}

func evalNot(v values.Value, log *message.Log) values.Value {
	if v.Type() != values.BOOL {
		message.RuntimeErr("Type", "Must have a boolean for '!' got '"+v.Stringify()+"'", v.GetTok(), log)
	}

	return &values.Boolean{!v.(*values.Boolean).Value, v.GetTok()}
}

func evalNegative(v values.Value, log *message.Log) values.Value {
	if v.Type() != values.NUM {
		message.RuntimeErr("Type", "Must have number for '-'", v.GetTok(), log)
	}

	return &values.Number{-v.(*values.Number).Value, v.GetTok()}
}
