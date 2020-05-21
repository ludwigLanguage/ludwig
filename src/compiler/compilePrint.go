package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compilePrint(node ast.Node) {
	print := node.(ast.Print)

	for _, i := range print.Args {
		c.Compile(i)
	}

	c.emit(bytecode.PRINT)
}
