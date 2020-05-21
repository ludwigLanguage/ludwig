package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compileAssignment(expr ast.InfixExpr) {
	c.Compile(expr.Right) //The value of the right will become the topmost value on the stack
	id := expr.Left.(ast.Identifier).Value

	symbol := c.symbols.Define(id)
	c.emit(bytecode.SAVEV, symbol) //This bytecode saves the topmost value on the stack
}

func (c *Compiler) compileIdent(expr ast.Node) {
	id := expr.(ast.Identifier).Value

	symbol, ok := c.symbols.Resolve(id)
	if !ok {
		c.raiseError("Identifier", "Could not resolve this identifier '"+id+"'", expr.GetTok())
	}

	c.emit(bytecode.GETV, symbol)
}
