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
	boolean := node.(ast.Boolean).Value
	val := values.Boolean{boolean}
	c.emit(bytecode.LOADCONST, c.addToPool(val))
}

func (c *Compiler) compileStr(node ast.Node) {
	str := node.(ast.String).Value
	val := values.String{str}
	c.emit(bytecode.LOADCONST, c.addToPool(val))
}

func (c *Compiler) compileNil(node ast.Node) {
	c.emit(bytecode.LOADCONST, c.addToPool(values.Nil{}))
}
