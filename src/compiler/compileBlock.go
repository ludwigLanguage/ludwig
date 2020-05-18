package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

//TODO: Enforce scoping rules
func (c *Compiler) compileBlock(node ast.Node) {
	block := node.(ast.Block)

	numOfDefsBeforeBlock := c.symbols.NumberOfDefinitions
	length := len(block.Body)
	for iter, expr := range block.Body {
		c.Compile(expr)

		if iter < length-1 {
			c.emit(bytecode.POP)
		}
	}

	if block.IsScoped {
		c.symbols.ClearDefsBackTo(numOfDefsBeforeBlock)
	}

	if length == 0 {
		c.compileNil(nil)
	}
}
