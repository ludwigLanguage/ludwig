package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compileAssignment(expr ast.InfixExpr) {
	c.Compile(expr.Right)
	id := expr.Left.(ast.Identifier).Value

	symbol := c.symbols.Define(id)
	c.emit(bytecode.SETG, symbol.Index)
}

func (c *Compiler) compileIdent(expr ast.Node) {
	id := expr.(ast.Identifier).Value

	symbol, ok := c.symbols.Resolve(id)
	if !ok {
		c.raiseError("Identifier", "Could not resolve this identifier", expr.GetTok())
	}

	c.emit(bytecode.GETG, symbol.Index)
}
