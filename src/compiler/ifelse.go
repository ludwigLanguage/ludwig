package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compileIfElse(node ast.Node) {
	ifel := node.(ast.IfEl)

	c.Compile(ifel.Cond)
	locationJumpNt := c.emit(bytecode.JUMPNT, 0)

	c.Compile(ifel.Do)
	locationJump := c.emit(bytecode.JUMP, 0)

	afterDoPos := len(c.instructions)
	c.changeArg(afterDoPos, locationJumpNt)

	c.Compile(ifel.ElseExpr)

	afterElsePos := len(c.instructions)
	c.changeArg(afterElsePos, locationJump)

}
