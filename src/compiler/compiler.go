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
	pool                []values.Value
	symbols             *SymTab
	mapNodesToCompileFn map[byte]compileFn

	scopeStack   []Scope
	scopePointer int
}

func New() *Compiler {
	mainScope := Scope{bytecode.Instructions{}}

	c := &Compiler{
		pool:    []values.Value{},
		symbols: NewST(),

		scopeStack:   []Scope{mainScope},
		scopePointer: 0,
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
		ast.FN:     c.compileFunctions,
		ast.PRINT:  c.compilePrint,
		ast.CALL:   c.compileCall,
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

func (c *Compiler) getCurScopeInstructions() bytecode.Instructions {
	return c.scopeStack[c.scopePointer].instructions
}

func (c *Compiler) addScope() {
	sub := Scope{bytecode.Instructions{}}
	c.scopeStack = append(c.scopeStack, sub)
	c.scopePointer++
}

func (c *Compiler) popScope() bytecode.Instructions {
	ins := c.getCurScopeInstructions()
	c.scopeStack = c.scopeStack[:len(c.scopeStack)-1]
	c.scopePointer--
	return ins
}

func (c *Compiler) GetCompiled() *CompiledProg {
	return &CompiledProg{
		Instructions: c.getCurScopeInstructions(),
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
	location := len(c.getCurScopeInstructions())

	c.scopeStack[c.scopePointer].instructions =
		append(c.getCurScopeInstructions(), instruction...)

	return location
}

func (c *Compiler) changeArg(arg int, opPos int) {
	op := bytecode.OpCode(c.getCurScopeInstructions()[opPos])
	newInstruction := bytecode.MakeInstruction(op, arg)

	c.backpatch(newInstruction, opPos)
}

func (c *Compiler) backpatch(instruction []byte, pos int) {
	for i := 0; i < len(instruction); i++ {
		c.scopeStack[c.scopePointer].instructions[pos+i] = instruction[i]
	}
}
