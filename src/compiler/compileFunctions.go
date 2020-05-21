package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
	"ludwig/src/values"
)

func (c *Compiler) compileFunctions(node ast.Node) {
	fn := node.(ast.Function)

	c.addScope()
	c.symbols.SaveState()

	c.Compile(fn.DoExpr)

	c.symbols.Revert()
	instructions := c.popScope()
	poolIndex := c.addToPool(values.Function{instructions})
	c.emit(bytecode.LOADCONST, poolIndex)

}
