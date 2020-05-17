package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compilePrefix(node ast.Node) {
	prefixExpr := node.(ast.PrefixExpr)
	c.Compile(prefixExpr.Expr)

	switch prefixExpr.Op {
	case "!":
		c.emit(bytecode.NOT)
	case "-":
		c.emit(bytecode.NEGATIVE)
	default:
		c.raiseError("Operator", "This operator is not definde", node.GetTok())
	}

}
