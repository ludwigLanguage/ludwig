package vm

import (
	"ludwig/src/bytecode"
	"ludwig/src/values"
)

func (v *VM) evalOpConst(location int) int {
	constIndex := bytecode.ReadUint16(v.currentFrame().Instructions()[location+1:])
	location += 2

	v.push(v.pool[constIndex])
	v.pool[constIndex] = nil //Clear out object

	return location
}

func (v *VM) evalPop(location int) int {
	v.pop()
	return location
}

func (v *VM) StackTop() values.Value {
	if v.stackPointer == 0 {
		return nil
	}

	return v.stack[v.stackPointer-1]
}

func (v *VM) LastPopped() values.Value {
	return v.stack[v.stackPointer]
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
