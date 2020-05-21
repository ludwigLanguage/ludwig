package vm

import (
	"ludwig/src/bytecode"
	"ludwig/src/compiler"
	"ludwig/src/message"
	"ludwig/src/values"
	"strconv"
)

const STACK_SIZE = 2048
const GLOBAL_HEAP_SIZE = 65536
const FRAME_STACK_SIZE = 1024

/*These functions will take in an int
 * that points to the instructions, and
 * once done executing the instructions,
 * it will return a pointer pointing to the
 * next instruction
 */
type executeFn func(int) int

type VM struct {
	pool []values.Value

	executeFnMap map[bytecode.OpCode]executeFn

	curLineNo int
	curFile   string

	stack []values.Value
	/* A note on conventions:
	 * stackPointer will always point to the next
	 * empty slot in the stack, if the stack is empty,
	 * it will point to index 0, if the stack has one
	 * item it will point to index 1, and so on
	 */
	stackPointer int

	frames       []*Frame
	framePointer int

	variables []values.Value
}

func New(program *compiler.CompiledProg) *VM {

	vm := &VM{
		pool: program.Pool,

		stack:        make([]values.Value, STACK_SIZE),
		stackPointer: 0,

		frames:       make([]*Frame, FRAME_STACK_SIZE),
		framePointer: 0,

		curFile:   "unknown.ldg",
		curLineNo: 0,

		variables: make([]values.Value, GLOBAL_HEAP_SIZE),
	}

	globalFrameFn := values.Function{program.Instructions, 0}
	globalFrame := NewFrame(globalFrameFn, nil)
	vm.pushFrame(globalFrame)

	vm.executeFnMap = map[bytecode.OpCode]executeFn{
		bytecode.LOADCONST: vm.evalOpConst,
		bytecode.POP:       vm.evalPop,

		bytecode.ADD:  vm.evalAdd,
		bytecode.SUB:  vm.evalSubtract,
		bytecode.MULT: vm.evalMultiply,
		bytecode.DIV:  vm.evalDivide,
		bytecode.POW:  vm.evalPower,

		bytecode.EQUALTO:       vm.evalEqualTo,
		bytecode.NOTEQUAL:      vm.evalNotEqual,
		bytecode.LESSTHAN:      vm.evalLessThan,
		bytecode.GREATERTHAN:   vm.evalGreaterThan,
		bytecode.LESSEREQUALS:  vm.evalLessEquals,
		bytecode.GREATEREQUALS: vm.evalGreaterLessEquals,

		bytecode.OR:  vm.evalOr,
		bytecode.AND: vm.evalAnd,

		bytecode.NOT:      vm.evalNot,
		bytecode.NEGATIVE: vm.evalNegative,

		bytecode.JUMP:   vm.evalJump,
		bytecode.JUMPNT: vm.evalJumpIfNotTrue,

		bytecode.SAVEV: vm.evalSaveVal,
		bytecode.GETV:  vm.evalGetVal,

		bytecode.BUILDLIST: vm.evalBuildList,
		bytecode.SLICE:     vm.evalSlice,
		bytecode.INDEX:     vm.evalIndex,
		bytecode.PRINT:     vm.evalPrint,
		bytecode.CALL:      vm.evalCall,
	}

	return vm
}

//FIXME
func (v *VM) raiseError(errtype, errmsg string) {
	message.VmError(errtype, errmsg, v.curFile, v.curLineNo)
}

func (v *VM) Run() {
	for v.currentFrame().localInsLocation <= len(v.currentFrame().Instructions())-1 {

		opcode := v.currentFrame().Instructions()[v.currentFrame().localInsLocation]
		executeFn := v.executeFnMap[bytecode.OpCode(opcode)]

		if executeFn == nil {
			strOfOpcode := strconv.Itoa(int(opcode))
			v.raiseError("Implement", "Implementation not added for instruction '"+strOfOpcode+"'")
		}

		nextLocation := executeFn(v.currentFrame().localInsLocation)
		v.currentFrame().localInsLocation = nextLocation + 1 //The one accounts for the length of the opcode
	}
}
