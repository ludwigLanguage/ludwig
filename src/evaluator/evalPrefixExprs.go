package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
	"ludwig/src/source"
	"ludwig/src/lexer"
	"ludwig/src/parser"
)

func evalPrefix(n *ast.PrefixExpr, consts *values.SymTab) values.Value {
	switch n.Op {
	case "!":
		return evalNot(EvalExpr(n.Expr, consts))
	case "-":
		return evalNegative(EvalExpr(n.Expr, consts))
	case "$":
		return evalDollar(EvalExpr(n.Expr, consts), consts)
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

func evalDollar(val values.Value, consts *values.SymTab) values.Value {
	if val.Type() != values.STR {
		message.RaiseError("Type", "Must have a string for '$'", val.GetTok())
	}

	src := source.NewWithStr(val.(*values.String).Value, val.GetTok().Filename)
	lxr := lexer.New(src)
	prs := parser.New(lxr)
	prs.ParseProgram()

	return EvalExpr(prs.Tree, consts)
}
