package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
)

func evalNum(n *ast.Number) values.Value {
	return &values.Number{n.Value, n.Tok}
}

func evalStr(n *ast.String) values.Value {
	return &values.String{n.Value, n.Tok}
}

func evalBool(n *ast.Boolean) values.Value {
	return &values.Boolean{n.Value, n.Tok}
}

func evalNil(n *ast.Nil) values.Value {
	return &values.Nil {n.Tok}
}

func evalIdent(n *ast.Identifier, consts *values.SymTab) values.Value {
	v := consts.GetVal(n.Value)
	if v != nil {
		return v
	}

	v = values.BuiltinsMap[n.Value]
	if v != nil {
		return v
	}

	message.RaiseError("Ident", "No such identifier '"+n.Value+"'", n.Tok)
	return NIL
}
