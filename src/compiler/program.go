package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
	"ludwig/src/values"
	"os"
)

func (c *Compiler) compileProgram(node ast.Node) {
	program := node.(ast.Program)

	for _, i := range program.Body {
		c.Compile(i)
		c.emit(bytecode.POP)
	}

	id, ok := c.symbols.Resolve("__main__")
	if !ok {
		c.raiseError("Procedural", "All programs must have a '__main__' function", program.GetTok())
	}

	argsList := []values.Value{}
	for _, i := range os.Args[2:] {
		argsList = append(argsList, values.String{i})
	}

	argsListVal := values.List{argsList}
	c.emit(bytecode.LOADCONST, c.addToPool(argsListVal))
	c.emit(bytecode.GETV, id)
	c.emit(bytecode.CALL, 1)
}
