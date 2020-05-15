package compiler

import (
	"ludwig/src/ast"
	"ludwig/src/bytecode"
	"ludwig/src/message"
	"ludwig/src/tokens"
	"ludwig/src/values"
	"strconv"
)

type compileFn func(ast.Node)

type Compiler struct {
	instructions bytecode.Instructions
	pool         []values.Value

	mapNodesToCompileFn map[byte]compileFn
}

func New() *Compiler {
	c := &Compiler{
		instructions: bytecode.Instructions{},
		pool:         []values.Value{},
	}

	c.mapNodesToCompileFn = map[byte]compileFn{
		ast.PROG:  c.compileProgram,
		ast.INFIX: c.compileInfix,
		ast.NUM:   c.compileNumber,
	}

	return c
}

func (c *Compiler) raiseError(errtype, errmsg string, tok tokens.Token) {
	message.RaiseError(errtype, errmsg, tok)
}

func (c *Compiler) Compile(node ast.Node) {
	fn := c.mapNodesToCompileFn[node.Type()]

	if fn == nil {
		t := strconv.Itoa(int(node.Type()))
		c.raiseError("Implementation", "Impossible to compile this node. ID: "+t, node.GetTok())
	}

	fn(node)
}

func (c *Compiler) GetCompiled() *CompiledProg {
	return &CompiledProg{
		Instructions: c.instructions,
		Pool:         c.pool,
	}
}

func (c *Compiler) addToPool(val values.Value) int {
	c.pool = append(c.pool, val)
	return len(c.pool) - 1 //Return the location of the constant
}

func (c *Compiler) emit(op bytecode.OpCode, args ...int) int {
	instruction := bytecode.MakeInstruction(op, args...)
	location := c.addInstruction(instruction)

	return location
}

func (c *Compiler) addInstruction(instruction []byte) int {
	location := len(c.instructions)
	c.instructions = append(c.instructions, instruction...)
	return location
}
