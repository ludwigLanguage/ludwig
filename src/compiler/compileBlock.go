package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

//TODO: Enforce scoping rules
func (c *Compiler) compileBlock(node ast.Node) {
	block := node.(ast.Block)

	if len(block.Body) == 0 {
		c.compileNil(nil)
	}

	if block.IsScoped {
		c.compileScopedBlock(block)
	} else {
		c.compileUnScopedBlock(block)
	}
}

func (c *Compiler) compileScopedBlock(block ast.Block) {
	c.symbols.SaveState()
	c.compileUnScopedBlock(block)
	c.symbols.Revert()
}

func (c *Compiler) compileUnScopedBlock(block ast.Block) {
	length := len(block.Body)
	for iter, expr := range block.Body {
		c.Compile(expr)

		if iter < length-1 {
			c.emit(bytecode.POP)
		}
	}
}
