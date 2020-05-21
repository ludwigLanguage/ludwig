package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compileCall(node ast.Node) {
	call := node.(ast.Call)

	/* Because the stack is first in; last out,
	 * we must compile the arguments from the end
	 * to the beginning
	 */
	maxArgsIndex := len(call.Args) - 1
	for i := 0; i <= maxArgsIndex; i++ {
		arg := call.Args[maxArgsIndex-i]

		c.Compile(arg)
	}

	c.Compile(call.CalledVal)
	c.emit(bytecode.CALL, len(call.Args))
}
