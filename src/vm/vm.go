package vm

import (
	"ludwig/src/bytecode"
	"ludwig/src/compiler"
	"ludwig/src/message"
	"ludwig/src/tokens"
	"ludwig/src/values"
	"strconv"
)

const STACK_SIZE = 2048

/*These functions will take in an int
 * that points to the instructions, and
 * once done executing the instructions,
 * it will return a pointer pointing to the
 * next instruction
 */
type executeFn func(int) int

type VM struct {
	pool         []values.Value
	instructions bytecode.Instructions

	executeFnMap map[bytecode.OpCode]executeFn

	stack []values.Value

	/* A note on conventions:
	 * stackPointer will always point to the next
	 * empty slot in the stack, if the stack is empty,
	 * it will point to index 0, if the stack has one
	 * item it will point to index 1, and so on
	 */
	stackPointer int
}

func New(program *compiler.CompiledProg) *VM {
	vm := &VM{
		instructions: program.Instructions,
		pool:         program.Pool,
		stack:        make([]values.Value, STACK_SIZE),
		stackPointer: 0,
	}

	vm.executeFnMap = map[bytecode.OpCode]executeFn{
		bytecode.LOADCONST: vm.evalOpConst,
		bytecode.ADD:       vm.add,
	}

	return vm
}

//FIXME
func (v *VM) raiseError(errtype, errmsg string) {
	tok := tokens.Token{"", 0, 0, "", 0}
	if len(v.stack) != 0 {
		tok = v.stack[v.stackPointer-1].GetTok()
	}

	message.RaiseError(errtype, errmsg, tok)
}

func (v *VM) StackTop() values.Value {
	if v.stackPointer == 0 {
		return nil
	}

	return v.stack[v.stackPointer-1]
}

func (v *VM) push(val values.Value) {
	if v.stackPointer >= STACK_SIZE {
		v.raiseError("Stack", "Stack overflow occured")
	}

	v.stack[v.stackPointer] = val
	v.stackPointer++
}

func (v *VM) pop() values.Value {
	val := v.stack[v.stackPointer-1]
	v.stackPointer--
	return val
}

func (v *VM) Run() {
	for insPos := 0; insPos < len(v.instructions); insPos++ {
		opcode := v.instructions[insPos]
		executeFn := v.executeFnMap[bytecode.OpCode(opcode)]

		if executeFn == nil {
			strOfOpcode := strconv.Itoa(int(opcode))
			v.raiseError("Implement", "Implementation not added for instruction '"+strOfOpcode+"'")
		}

		insPos = executeFn(insPos)
	}
}
