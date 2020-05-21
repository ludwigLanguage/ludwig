package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compileCall(node ast.Node) {
	call := node.(ast.Call)

	c.Compile(call.CalledVal)
	c.emit(bytecode.CALL)
}
