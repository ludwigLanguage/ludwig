package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
	"ludwig/src/values"
)

func (c *Compiler) compileNumber(node ast.Node) {
	number := node.(ast.Number)

	val := values.Number{number.Value}
	c.emit(bytecode.LOADCONST, c.addToPool(val))
}

func (c *Compiler) compileBool(node ast.Node) {
	bool := node.(ast.Boolean)
	val := values.Boolean{bool.Value}
	c.emit(bytecode.LOADCONST, c.addToPool(val))
}

func (c *Compiler) compileNil(node ast.Node) {
	c.emit(bytecode.LOADCONST, c.addToPool(values.Nil{}))
}
