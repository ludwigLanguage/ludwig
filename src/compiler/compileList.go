package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

func (c *Compiler) compileList(node ast.Node) {
	list := node.(ast.List)

	for _, i := range list.Entries {
		c.Compile(i)
	}
	length := len(list.Entries)

	c.emit(bytecode.BUILDLIST, length)
}

func (c *Compiler) compileSlice(node ast.Node) {
	slice := node.(ast.Slice)
	c.Compile(slice.Src)
	c.Compile(slice.Start)
	if slice.End == nil {
		c.compileNil(slice.End)
	} else {
		c.Compile(slice.End)
	}
	c.emit(bytecode.SLICE)
}

func (c *Compiler) compileIndex(node ast.Node) {
	index := node.(ast.Index)
	c.Compile(index.Src)
	c.Compile(index.Index)
	c.emit(bytecode.INDEX)
}
