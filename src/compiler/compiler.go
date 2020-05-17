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
	instructions        bytecode.Instructions
	pool                []values.Value
	symbols             *SymTab
	mapNodesToCompileFn map[byte]compileFn
}

func New() *Compiler {
	c := &Compiler{
		instructions: bytecode.Instructions{},
		pool:         []values.Value{},
		symbols:      NewST(),
	}

	c.mapNodesToCompileFn = map[byte]compileFn{
		ast.PROG:   c.compileProgram,
		ast.INFIX:  c.compileInfix,
		ast.PREFIX: c.compilePrefix,
		ast.NUM:    c.compileNumber,
		ast.BOOL:   c.compileBool,
		ast.BLOCK:  c.compileBlock,
		ast.IFEL:   c.compileIfElse,
		ast.NIL:    c.compileNil,
		ast.IDENT:  c.compileIdent,
		ast.STR:    c.compileStr,
		ast.LIST:   c.compileList,
		ast.SLICE:  c.compileSlice,
		ast.INDEX:  c.compileIndex,
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

func (c *Compiler) changeArg(arg int, opPos int) {
	op := bytecode.OpCode(c.instructions[opPos])
	newInstruction := bytecode.MakeInstruction(op, arg)

	c.backpatch(newInstruction, opPos)
}

func (c *Compiler) backpatch(instruction []byte, pos int) {
	for i := 0; i < len(instruction); i++ {
		c.instructions[pos+i] = instruction[i]
	}
}
