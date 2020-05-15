package compiler

import (
	"ludwig/src/ast"
)

func (c *Compiler) compileProgram(node ast.Node) {
	program := node.(*ast.Program)

	for _, i := range program.Body {
		c.Compile(i)
	}
}
