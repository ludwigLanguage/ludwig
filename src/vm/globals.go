package vm

import (
	"ludwig/src/bytecode"
)

func (v *VM) evalSetg(location int) int {
	valIndex := bytecode.ReadUint16(v.instructions[location+1:])
	location += 2

	val := v.pop()
	v.globals[valIndex] = val
	v.push(val)

	return location
}

func (v *VM) evalGetg(location int) int {
	valIndex := bytecode.ReadUint16(v.instructions[location+1:])
	location += 2
	v.push(v.globals[valIndex])
	return location
}
