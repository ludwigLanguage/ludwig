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

	for _, i := range fn.Args {
		symbol := c.symbols.Define(i.Value)
		c.emit(bytecode.SAVEV, symbol)
		c.emit(bytecode.POP)
	}
	c.Compile(fn.DoExpr)

	c.symbols.Revert()
	instructions := c.popScope()
	poolIndex := c.addToPool(values.Function{instructions, len(fn.Args)})
	c.emit(bytecode.LOADCONST, poolIndex)

}
