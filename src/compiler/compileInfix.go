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
	case "==":
		c.emit(bytecode.EQUALTO)
	case "!=":
		c.emit(bytecode.NOTEQUAL)
	case "<":
		c.emit(bytecode.LESSTHAN)
	case ">":
		c.emit(bytecode.GREATERTHAN)
	case "<=":
		c.emit(bytecode.LESSEREQUALS)
	case ">=":
		c.emit(bytecode.GREATEREQUALS)
	case "||":
		c.emit(bytecode.OR)
	case "&&":
		c.emit(bytecode.AND)
	default:
		c.raiseError("Syntax", "Unknown operator", node.GetTok())
	}
}
