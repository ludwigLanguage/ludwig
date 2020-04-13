package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/values"
	"ludwig/src/message"
)

func evalQuote(n *ast.Quote, consts *values.SymTab) values.Value {
	return &values.QuotedVal {n.Expr, n.GetTok()}
}

func evalUnQuote(n *ast.UnQuote, consts *values.SymTab) values.Value {
	quotedTree := EvalExpr(n.Expr, consts)
	
	if quotedTree.Type() != values.QUOTE {
		message.RaiseError("Type", "Expected quoted expression got '" + quotedTree.Type() + "'", n.GetTok())
	}

	return EvalExpr(quotedTree.(*values.QuotedVal).Node, consts)
} 