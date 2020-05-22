package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
)

const (
	PRINT int = iota
	PRINTLN
)

var builtinMap = map[string]int{
	"print":   PRINT,
	"println": PRINTLN,
}

func (c *Compiler) compileBuiltinCall(node ast.Node) {
	call := node.(ast.Builtin)

	for _, i := range call.Args {
		c.Compile(i)
	}

	id := builtinMap[call.BuiltinName]
	c.emit(bytecode.CALLBUILTIN, id)
}
