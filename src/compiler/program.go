package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compileProgram(node ast.Node) {
	program := node.(ast.Program)

	for _, i := range program.Body {
		c.Compile(i)
		c.emit(bytecode.POP)
	}
}
