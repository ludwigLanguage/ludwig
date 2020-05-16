package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compileInfix(node ast.Node) {
	infix := node.(ast.InfixExpr)
	c.Compile(infix.Left)
	c.Compile(infix.Right)

	switch infix.Op {
	case "+":
		c.emit(bytecode.ADD)
	case "-":
		c.emit(bytecode.SUB)
	case "*":
		c.emit(bytecode.MULT)
	case "/":
		c.emit(bytecode.DIV)
	case "^":
		c.emit(bytecode.POW)
	default:
		c.raiseError("Syntax", "Unknown operator", node.GetTok())
	}
}
